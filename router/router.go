package router

import (
	"log"

	"github.com/anudeep-mp/portfolio-backend/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/api/sendmessage", controller.SendMessage).Methods("POST")

	return router
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
