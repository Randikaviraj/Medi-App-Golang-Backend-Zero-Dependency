package routes

import (
	"net/http"

	"github.com/randika/Medi-App-Golang-Backend/controllers"
)

func Setup()  {
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("Server working.."))
	})

	http.HandleFunc("/date",controllers.FindByDateController)
	http.HandleFunc("/doctor",controllers.FindByNameController)
	http.HandleFunc("/hospital",controllers.FindByHospitalController)
	http.HandleFunc("/specility",controllers.FindBySpecilityController)
}