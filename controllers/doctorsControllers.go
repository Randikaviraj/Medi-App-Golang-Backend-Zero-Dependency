package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/randika/Medi-App-Golang-Backend/database"
)

type doctorName struct{
	name string
}

type speciality struct{
	specializedField string
}

// var doctors []Doctor
// 	QueryDoctorByDate("Navaloka hospitals","Mon-Fri",&doctors)
// 	fmt.Println(doctors)

func FindByNameController(w http.ResponseWriter,r *http.Request){

	var doc doctorName

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&doc)
	
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	var data []database.HospitalAndDoctor
	if err:=database.QueryByDoctorName(doc.name,&data);err!=nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}
	jsonData,err:=json.Marshal(data)

	if err!=nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	w.Header().Add("Access-Control-Allow-Origin","*")
	w.Write(jsonData)
}