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
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertMessage(message model.Message) (primitive.ObjectID, error) {
	inserted, err := database.MessagesCollection.InsertOne(context.Background(), message)

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
	cursor, err := database.MessagesCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, fmt.Errorf("error finding messages: %w", err)
	}

	defer cursor.Close(context.Background())

	var messages = make([]model.Message, 0)

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
	_, err := database.MessagesCollection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		return fmt.Errorf("failed to delete messages: %w", err)
	}
	return nil
}

func PostWatchStamp(watchStamp model.WatchStamp) (model.UserStamp, error) {

	filter := bson.M{"userId": watchStamp.UserID}

	var user model.UserStamp

	databaseErr := database.TrackingCollection.FindOne(context.Background(), filter).Decode(&user)

	if databaseErr == mongo.ErrNoDocuments {
		//if documnet if not available, create a new one
		user = model.UserStamp{
			UserID:   watchStamp.UserID,
			Sessions: make([]model.SessionStamp, 0),
		}

		user = updateSession(watchStamp, user)

		insertedRes, insertionErr := database.TrackingCollection.InsertOne(context.Background(), user)

		if insertionErr != nil {
			return user, fmt.Errorf("failed to insert new user %w", insertionErr)
		}

		user.ID = insertedRes.InsertedID.(primitive.ObjectID)

	} else {

		user = updateSession(watchStamp, user)

		update := bson.M{"$set": bson.M{"sessions": user.Sessions}}

		_, updatedErr := database.TrackingCollection.UpdateOne(context.Background(), filter, update)

		if updatedErr != nil {
			return user, fmt.Errorf("failed to update user %w", updatedErr)
		}
	}

	return user, nil
}

func updateSession(watchStamp model.WatchStamp, user model.UserStamp) model.UserStamp {

	var isSessionAvailable = false

	for i := range user.Sessions {
		tempSession := user.Sessions[i]

		if tempSession.SessionID == watchStamp.SessionID {
			tempSession.TimeStamps = append(tempSession.TimeStamps, watchStamp.TimeStamp)
			user.Sessions[i] = tempSession
			isSessionAvailable = true
			break
		}
	}

	if !isSessionAvailable {
		var newSession = model.SessionStamp{
			UserID:    watchStamp.UserID,
			SessionID: watchStamp.SessionID,
			TimeStamps: []string{
				watchStamp.TimeStamp,
			},
		}
		user.Sessions = append(user.Sessions, newSession)
	}

	return user
}

func GetWatchStamps() ([]model.UserStamp, error) {
	cursor, err := database.TrackingCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	defer cursor.Close(context.Background())

	var users = make([]model.UserStamp, 0)

	for cursor.Next(context.Background()) {
		var user model.UserStamp

		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("error decoding user details: %w", err)
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error during cursor iteration: %w", err)
	}

	return users, nil
}

func DeleteAllWatchStamps() error {
	_, err := database.TrackingCollection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		return fmt.Errorf("failed to delete userstamps: %w", err)
	}
	return nil
}
