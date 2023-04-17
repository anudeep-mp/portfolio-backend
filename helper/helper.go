package helper

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/anudeep-mp/portfolio-backend/database"
	"github.com/anudeep-mp/portfolio-backend/model"
)

func InsertMessage(message model.Message) {
	inserted, err := database.Collection.InsertOne(context.Background(), message)
	CheckError(err)

	fmt.Println("Inserted one result : ", inserted.InsertedID)

}

func SendMail(message model.Message) {
	auth := smtp.PlainAuth(
		"",
		"anudeep.mp7@gmail.com",
		os.Getenv("EMAIL_SMTP_PASSWORD"),
		"smtp.gmail.com",
	)

	msg := []byte("To: m.anudeep2000@gmail.com\r\n" +
		"Subject: " + message.Subject + " - " + message.Name + "\r\n" +
		"\r\n" +
		message.Message + "\r\n")

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"anudeep.mp7@gmail.com",
		[]string{"m.anudeep2000@gmail.com"},
		[]byte(msg),
	)

	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
