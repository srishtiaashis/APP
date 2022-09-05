package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/srishtiaashis/App/Patient/Dlist"
	"github.com/srishtiaashis/App/Patient/database"

	_ "github.com/go-sql-driver/mysql"
)

func DoctorList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		doctorName, err := Dlist.GetDoctor()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		doctorJSON, err := json.Marshal(doctorName)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application.json")
		w.Write(doctorJSON)
	case http.MethodPost:
		var DoctorName Dlist.DoctorDB
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &DoctorName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		doctorData, err := Dlist.GetDoctorByTimeSlot(DoctorName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		doctorDataJSON, err := json.Marshal(doctorData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application.json")
		w.Write(doctorDataJSON)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}
func SetAppointment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var app Dlist.Appoint
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &app)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = Dlist.InsertAppointment(app)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}
}
func main() {
	http.HandleFunc("/getDoctor", DoctorList)
	http.HandleFunc("/setAppointment", SetAppointment)
	database.SetupConnection()
	http.ListenAndServe(":8000", nil)
}
