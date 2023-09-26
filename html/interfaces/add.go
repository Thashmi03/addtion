package interfaces

import (
	"add/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Add interface{
	Add(m *models.Add)(*mongo.InsertOneResult)
}