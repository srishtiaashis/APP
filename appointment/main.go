package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/srishtiaashis/App/appointment/database"

	"github.com/srishtiaashis/App/appointment/apntmtup"

	_ "github.com/go-sql-driver/mysql"
)

func ChangeAppointment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var apoint apntmtup.Apntmtupdate
		bodybytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodybytes, &apoint)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = apntmtup.Papntmtupdate(apoint)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
func main() {
	http.HandleFunc("/changeAppoint", ChangeAppointment)
	database.SetupConnection()
	http.ListenAndServe(":9000", nil)
}
