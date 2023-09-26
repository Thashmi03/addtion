package models

type Add struct {
	FirstNumber  int64 `json:"first_number,omitempty" bson:"first_number"`
	SecondNumber int64 `json:"second_number,omitempty" bson:"second_number"`
	Sum          int64 `json:"sum,omitempty" bson:"sum"`
}
