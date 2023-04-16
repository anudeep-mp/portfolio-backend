package controller

import (
	"encoding/json"
	"net/http"

	"github.com/anudeep-mp/portfolio-backend/helper"
	"github.com/anudeep-mp/portfolio-backend/model"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Anudeep's portfolio</h1>"))
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var message model.Message

	_ = json.NewDecoder(r.Body).Decode(&message)

	helper.InsertMessage(message)
	helper.SendMail(message)

	json.NewEncoder(w).Encode("Message sent successfully")
}
