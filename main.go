package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getCode(msg string) []byte {
	h := hmac.New(sha256.New, []byte("mykey"))
	h.Write([]byte(msg))
	code := h.Sum(nil)
	fmt.Printf("GetCodeFn ---->   Hex value for email %s  is %x \n", msg, code)
	return code
}

var isEqual = false
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
		code := getCode(email)
		hexCode := fmt.Sprintf("%x", code)
		cookie, err := r.Cookie("session")
		if err != nil {
			c := http.Cookie{
				Name:  "session",
				Value: hexCode + "|" + email,
			}
			fmt.Printf("Setting cookie ---> %s\n", c.Value)
			http.SetCookie(w, &c)
		} else {
			fmt.Println("Cookie already present ", cookie.Value)
			xs := strings.SplitN(cookie.Value, "|", 2)
			if len(xs) == 2 {
				cCode := xs[0]
				data, _ := hex.DecodeString(cCode)
				code := getCode(email)
				fmt.Println("cCode is ", cCode)
				fmt.Println("Code is ", code)
				isEqual = hmac.Equal(data, code)
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
		} else {
			fmt.Println("Foundn cookie %s", c.Value)
			xs := strings.SplitN(c.Value, "|", 2)
			if len(xs) == 2 {
				cCode := xs[0]
				code := getCode(formEmail)
				data, _ := hex.DecodeString(cCode)
				fmt.Printf("cCode is %x", data)
				fmt.Printf("Code is   %x\n", code)

				if hmac.Equal(data, code) {
					isEqual = true
					fmt.Println("Match found")
					msg = "logged in"
				} else {
					fmt.Println("Match not found")
				}
			}
		}
		if isEqual {
			fmt.Println("Is Equal ", isEqual)
		} else {
			fmt.Println("isEqual ", isEqual)
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
