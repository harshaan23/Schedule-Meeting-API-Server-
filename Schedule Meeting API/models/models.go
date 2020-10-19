package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import "time"


type Meeting struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	StartTime  time.Time             `json:"starttime" bson:"starttime,omitempty"`
	EndTime  time.Time             `json:"endtime" bson:"endtime,omitempty"`
	Timestamp  time.Time             `json:"timestamp" bson:"timestamp,omitempty"`
	Participant *Participant            `json:"participant" bson:"participant,omitempty"`
}

type Paticipant struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Email  string `json:"email,omitempty" bson:"email,omitempty"`
	RSVP  string `json:"rsvp,omitempty" bson:"rsvp,omitempty"`
}
