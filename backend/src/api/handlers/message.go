package handlers

import (
	"context"
	"log"
	"wrapup/database"
	"wrapup/models"

	"net/http"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
)
var (
	upgrader = websocket.Upgrader{}
	// create a map to store websocket connections keyed by user ID
	connections = make(map[string]*websocket.Conn)
)

type MessageHanlder struct {
	Db *database.Client
}

func (mh *MessageHanlder) SendMessage(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	
	// create a new MongoDB collection to store messages
	coll := mh.Db.Database("wrapup").Collection("messages")

	for {
		// read message from the websocket connection
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		// parse the message payload to extract the recipient ID
		var message models.Message
		err = bson.Unmarshal(msg, &message)
		if err != nil {
			log.Println(err)
			break
		}

		if !CheckUserExists(mh.Db, message.SenderID) {
			conn.WriteMessage(websocket.TextMessage, []byte("Sender does not exist"))
			continue
		}
		if !CheckUserExists(mh.Db, message.RecipientID) {
			conn.WriteMessage(websocket.TextMessage, []byte("Recipient does not exist"))
			continue
		}
		// look up the recipient's websocket connection using the recipient ID
		recipientConn, ok := connections[message.RecipientID]
		if !ok {
			// recipient is not connected, send an error message back to the sender
			conn.WriteMessage(websocket.TextMessage, []byte("Recipient is not connected"))
			continue
		}
		// send message to the recipient
		err = recipientConn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
		// insert the message into the MongoDB collection
		_, err = coll.InsertOne(context.TODO(), message)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
