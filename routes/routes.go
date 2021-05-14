package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/randika/Medi-App-Golang-Backend/controllers"
)

func Setup() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.RequestURI
		path=strings.TrimSuffix(path,"/")
		fmt.Printf(path)

		switch path {

		case "/date":
			controllers.FindByDateController(w, r)

		case "/doctor":
			controllers.FindByNameController(w, r)

		case "/hospital":
			controllers.FindByHospitalController(w, r)

		case "/specility":
			controllers.FindBySpecilityController(w, r)

		default:
			w.Write([]byte("Server working.." + path))

		}

	})
}
