package login

import (
	//"database/sql"
	"fmt"
	//"html/template"
	"log"
	"net/http"
	//"unicode"
	//_ "github.com/go-sql-driver/mysql"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "----------WELCOME TO APPOINTMENT BOOKING PORTAL---------")
}
func handleRequest() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
