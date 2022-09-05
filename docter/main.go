package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/srishtiaashis/App/docter/Plist"
	"github.com/srishtiaashis/App/docter/database"

	_ "github.com/go-sql-driver/mysql"
)

func getPatient(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var DocName Plist.Docname
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &DocName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		PatientDetails, err := Plist.Getmypatient(DocName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		PatientDetailsJSON, err := json.Marshal(PatientDetails)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application.json")
		w.Write(PatientDetailsJSON)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}
}
func main() {
	http.HandleFunc("/myPatient", getPatient)
	database.SetupConnection()
	http.ListenAndServe(":5000", nil)
}
