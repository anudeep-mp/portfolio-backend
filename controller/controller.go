package controller

import (
	"encoding/json"
	"net/http"

	"github.com/anudeep-mp/portfolio-backend/helper"
	"github.com/anudeep-mp/portfolio-backend/model"
	"github.com/anudeep-mp/portfolio-backend/utilities"
)

func ServeHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Anudeep's portfolio</h1>"))
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var message model.Message
	_ = json.NewDecoder(r.Body).Decode(&message)

	insertedId, errorInsertingMessage := helper.InsertMessage(message)

	message.ID = insertedId

	if err := helper.SendMail(message); err != nil || errorInsertingMessage != nil {
		utilities.ResponseWrapper(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	utilities.ResponseWrapper(w, http.StatusCreated, true, "Message sent successfully", message)
}

func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {

	messages, err := helper.GetMessages()

	if err != nil {
		utilities.ResponseWrapper(w, http.StatusInternalServerError, false, err.Error(), nil)
	}

	utilities.ResponseWrapper(w, http.StatusOK, true, "Messages fetched successfully", messages)
}

func DeleteAllMessagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	err := helper.DeleteAllMessages()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	utilities.ResponseWrapper(w, http.StatusOK, true, "All messages deleted successfully", nil)
}

func WatchStampHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var watchStamp model.WatchStamp

	_ = json.NewDecoder(r.Body).Decode(&watchStamp)

	user, err := helper.PostWatchStamp(watchStamp)

	if err != nil {
		utilities.ResponseWrapper(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	utilities.ResponseWrapper(w, http.StatusOK, true, "Added user document succesfully", user)
}

func GetWatchStampsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	var users []model.UserStamp

	users, err := helper.GetWatchStamps()

	if err != nil {
		utilities.ResponseWrapper(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	utilities.ResponseWrapper(w, http.StatusOK, true, "Users fetched successfully", users)
}

func DeleteAllWatchStampsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	err := helper.DeleteAllWatchStamps()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	utilities.ResponseWrapper(w, http.StatusOK, true, "All users deleted successfully", nil)
}
