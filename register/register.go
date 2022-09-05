package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template
var db *sql.DB

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("register handler running...")
	tpl.ExecuteTemplate(w, "register.html", nil)
	/*if err := r.ParseForm(); err != nil {
	fmt.Fprintf(w, "ParseForm() err: %v", err)
	return
	}*/
	fmt.Fprintf(w, "Post request successeful")

}

func authregHandler(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	firstName := r.FormValue("firstName")
	secondName := r.FormValue("secondName")
	//var nia bool
	//n = false
	if 5 <= len(firstName) && len(firstName) <= 15 && 5 <= len(secondName) && len(secondName) <= 15 {
		//nia := true

	}

	phoneNo := r.FormValue("PhoneNo")
	//numeric := true
	for _, char := range phoneNo {
		if unicode.IsNumber(char) == false {
			//numeric = false

		}
	}
	password := r.FormValue("password")
	fmt.Fprintf(w, "password:", password, "\npswdlen:", len(password))
	var pswdlc, pswduc, pswdno, pswdspec, pswdlen, pswdnospace bool
	pswdnospace = true
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			pswdlc = true
		case unicode.IsUpper(char):
			pswduc = true
		case unicode.IsNumber(char):
			pswdno = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdspec = true
		case unicode.IsSpace(int32(char)):
			pswdnospace = false
		}
	}
	if 11 < len(password) && len(password) < 20 {
		pswdlen = true
	}
	fmt.Println("pswdlc:", pswdlc, "\npswduc:", pswduc, "\npswdnumber:", pswdno, "\npswdnospce:", pswdnospace, "\npaswdspec:", pswdspec)
	if !pswdlc || !pswduc || !pswdlen || !pswdno || !pswdnospace || !pswdspec {
		tpl.ExecuteTemplate(w, "register.html", "Invalid username or password criteria")
		return
	}
	//stmt := "SELECT UserID FROM bcrypt WHERE username = ?"
	fmt.Fprintln(w, "account created successfully")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/authreg", authregHandler)
	http.ListenAndServe("localhost:8080", nil)

	fmt.Printf("Started server at port : 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
