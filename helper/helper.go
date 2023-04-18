package helper

import (
	"context"
	"fmt"
	"net/smtp"
	"os"

	"github.com/anudeep-mp/portfolio-backend/database"
	"github.com/anudeep-mp/portfolio-backend/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertMessage(message model.Message) (primitive.ObjectID, error) {
	inserted, err := database.Collection.InsertOne(context.Background(), message)

	insertedId := inserted.InsertedID.(primitive.ObjectID)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to insert message : %s", err)
	}

	fmt.Println("Inserted one result : ", insertedId)
	return insertedId, nil
}

func SendMail(message model.Message) error {
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

	if err != nil {
		return fmt.Errorf("failed to send mail: %s", err)
	}

	return nil
}

func GetMessages() ([]model.Message, error) {
	cursor, err := database.Collection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, fmt.Errorf("error finding messages: %w", err)
	}

	defer cursor.Close(context.Background())

	var messages []model.Message

	for cursor.Next(context.Background()) {
		var message model.Message

		if err := cursor.Decode(&message); err != nil {
			return nil, fmt.Errorf("error decoding message: %w", err)
		}

		messages = append(messages, message)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error during cursor iteration: %w", err)
	}

	return messages, nil
}

func DeleteAllMessages() error {
	_, err := database.Collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		return fmt.Errorf("failed to delete messages: %w", err)
	}
	return nil
}
