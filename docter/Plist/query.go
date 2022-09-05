package Plist

import (
	"github.com/srishtiaashis/App/docter/database"
)

func Getmypatient(doc Docname) ([]PDetails, error) {

	results, err := database.Dbconn.Query(`SELECT patient_name,patient_phone,patient_email FROM appointment WHERE doctor_id = (SELECT DOC_id FROM doctorlist WHERE Doc_name=?)`, doc.Name)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	patientslist := make([]PDetails, 0)
	for results.Next() {
		var patientlist PDetails
		results.Scan(&patientlist.P_name, &patientlist.P_phone, &patientlist.P_email)
		patientslist = append(patientslist, patientlist)
	}
	return patientslist, nil
}
