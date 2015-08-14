package profile

import (
	"gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Created  int32         `json:"created"`
	Status   string        `json:"status"`
	Username string        `json:"username" binding:"required"`
	Email    string        `json:"email" binding:"required"`
	Password string        `json:"password"`
}
