package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/faygun/go-rest-api/helper"
	"github.com/faygun/go-rest-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = helper.ConnectDB()

func getMeetings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var meetings []models.Meeting
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var meeting models.Meeting
		err := cur.Decode(&meeting) 
		if err != nil {
			log.Fatal(err)
		}
		books = append(meetings, meeting)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(meetings) }

func getMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var meeting models.Meeting
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&meeting)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(meeting)
}

func createMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var meeting models.Meeting

	_ = json.NewDecoder(r.Body).Decode(&meeting)

	result, err := collection.InsertOne(context.TODO(), meeting)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	meeting.Timestamp = time.Now()
	json.NewEncoder(w).Encode(result)
}

func updateMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var meeting models.Meeting

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&meeting)

	update := bson.D{
		{"$set", bson.D{
			{"title", meeting.Title},
			{"starttime", meeting.StartTime},
			{"endtime", meeting.EndTime},
			{time.Now(), meeting.Timestamp},
			{"Partcipant", bson.D{
				{"name", meeting.Participant.Name},
				{"email", meeting.Participant.Email},
				{"rsvp", meeting.Participant.RSVP}
			}},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&meeting)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	meeting.ID = id

	json.NewEncoder(w).Encode(book)
}

func deleteMeeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/meetings", getMeetings).Methods("GET")
	r.HandleFunc("/api/meetings/{id}", getMeeting).Methods("GET")
	r.HandleFunc("/api/meetings", createMeeting).Methods("POST")
	r.HandleFunc("/api/meetings/{id}", updateMeeting).Methods("PUT")
	r.HandleFunc("/api/meetings/{id}", deleteMeeting).Methods("DELETE")

	config := helper.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))

}
