package handlers

//import (
//	"encoding/json"
//	"net/http"
//
//	"wrapup/models"
//
//	"github.com/go-chi/chi/v5"
//)
//
//// CreateConversation creates a new conversation in the database
//func CreateConversation(w http.ResponseWriter, r *http.Request) {
//	var conversation models.Conversation
//	if err := json.NewDecoder(r.Body).Decode(&conversation); err != nil {
//		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
//		return
//	}
//	defer r.Body.Close()
//
//	if err := conversation.Create(); err != nil {
//		RespondWithError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	RespondWithJSON(w, http.StatusCreated, conversation)
//}
//
//// GetConversations retrieves a list of all conversations from the database
//func GetConversations(w http.ResponseWriter, r *http.Request) {
//	conversations, err := models.GetConversations()
//	if err != nil {
//		RespondWithError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	RespondWithJSON(w, http.StatusOK, conversations)
//}
//
//// GetConversation retrieves a single conversation by ID from the database
//func GetConversation(w http.ResponseWriter, r *http.Request) {
//	id := chi.URLParam(r, "conversationID")
//
//	conversation, err := models.GetConversationByID(id)
//	if err != nil {
//		RespondWithError(w, http.StatusBadRequest, "Invalid conversation ID")
//		return
//	}
//
//	RespondWithJSON(w, http.StatusOK, conversation)
//}
//
//// UpdateConversation updates an existing conversation in the database
//func UpdateConversation(w http.ResponseWriter, r *http.Request) {
//	id := chi.URLParam(r, "conversationID")
//
//	var conversation models.Conversation
//	if err := json.NewDecoder(r.Body).Decode(&conversation); err != nil {
//		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
//		return
//	}
//	defer r.Body.Close()
//	conversation.ID = id
//
//	if err := conversation.Update(); err != nil {
//		RespondWithError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//}
//
