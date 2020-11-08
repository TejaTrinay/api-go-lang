package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Articles struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"_title" bson:"_title,omitempty"`
	SubTitle  string          `json:"_subtitle" bson:"_subtitle,omitempty"`
	Content  string           `json:"_content" bson:"_content,omitempty"`
	CreationStamp  string     `json:"_creationstamp" bson:"_creationstamp,omitempty"`
}
