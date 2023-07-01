package allquestions

import "go.mongodb.org/mongo-driver/bson/primitive"

type AllQuestion struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Videourl string             `json:"videourl"`
	Category string             `json:"Category"`
	Name     string             `json:"Name"`
	Link     string             `json:"Link"`
	ID0      int                `json:"Id"`
	Level    string             `json:"Level"`
}
