package main

import (
	"fmt"
	"io"
	"net/http"
)

func getCode(msg string) string {

}

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
		c := http.Cookie{
			Name:  "session",
			Value: "",
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `<!DOCTYPE html>
				<html>
				<body>

				<h1>My First Heading</h1>
				<p>My first paragraph.</p>
				<form method  = 'POST' action  = "/submit">
					<input type = "email" name = "email" />
					<input type = "submit"/>
				</form>

				</body>
				</html>`
		io.WriteString(w, html)

	})
	fmt.Print("before")
	fmt.Println(http.ListenAndServe(":8000", nil))
	fmt.Println("after")
}
