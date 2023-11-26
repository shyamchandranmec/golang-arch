package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Email string `json: email`
	jwt.RegisteredClaims
}

var myKey = []byte("This is my new key")

func getJWT(email string) string {
	expTime := time.Now().Add(5 * time.Minute)
	claims := &MyClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(myKey)
	fmt.Println("Generation jwt")
	fmt.Printf("%v ", ss)
	return ss
}

func getCode(msg string) []byte {
	h := hmac.New(sha256.New, []byte("mykey"))
	h.Write([]byte(msg))
	code := h.Sum(nil)
	fmt.Printf("GetCodeFn ---->   Hex value for email %s  is %x \n", msg, code)
	return code
}

var isLoggedIn = false
var formEmail = ""

func main() {
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		email := r.FormValue("email")
		if email == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		formEmail = email
		cookie, err := r.Cookie("session")
		if err != nil {
			jwt := getJWT(email)
			c := http.Cookie{
				Name:  "session",
				Value: jwt,
			}
			fmt.Printf("\nSetting cookie ---> %s\n", c.Value)

			http.SetCookie(w, &c)
			isLoggedIn = true
		} else {
			fmt.Println("Cookie already present ", cookie.Value)
			cjwt := cookie.Value
			token, err := jwt.ParseWithClaims(cjwt, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(myKey), nil
			})
			if err != nil {
				fmt.Println("Token verification failed for cookie jwt ", err)
				isLoggedIn = false
			} else {
				if token.Valid {
					claims := token.Claims.(*MyClaims)
					fmt.Println("Email in claim is ", claims.Email)
					if claims.Email == email {
						isLoggedIn = true
					} else {
						isLoggedIn = false
					}
				} else {
					fmt.Println("Token is invalid")
				}

			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session")
		msg := "Not logged in"
		if err != nil {
			fmt.Println("No cookie found")
			c = &http.Cookie{}
		}
		if isLoggedIn {
			fmt.Println("isLoggedIn ", isLoggedIn)
			msg = "Logged in with same email"
		} else {
			fmt.Println("isLoggedIn", isLoggedIn)
			if formEmail == "" {
				msg = "Please add email"
			}
		}
		html := `<!DOCTYPE html>
				<html>
				<body>

				<h1>` + msg + `</h1>
				<p>` + c.Value + `</p>
				<form method  = 'POST' action  = "/submit">
					<input type = "email" name = "email" />
					<input type = "submit"/>
				</form>

				</body>
				</html>`
		io.WriteString(w, html)

	})
	fmt.Println(http.ListenAndServe(":8000", nil))
	fmt.Println("after")
}
