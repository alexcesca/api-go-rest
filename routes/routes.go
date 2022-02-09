package routes

import (
	"go-rest-api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades/", controllers.TodasAsPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	http.ListenAndServe(":8081", r)
}
