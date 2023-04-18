package router

import (
	"github.com/anudeep-mp/portfolio-backend/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.ServeHomeHandler).Methods("GET")
	router.HandleFunc("/api/sendmessage", controller.SendMessageHandler).Methods("POST")
	router.HandleFunc("/api/messages", controller.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/messages", controller.DeleteAllMessagesHandler).Methods("DELETE")
	router.HandleFunc("/api/track", controller.WatchStampHandler).Methods("POST")

	return router
}
