package services

import (
	"context"
	"add/interfaces"
	"add/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Add struct{
	Collection *mongo.Collection
	ctx context.Context
}

func InitService(collection *mongo.Collection,ctx context.Context)interfaces.Add{
	return &Add{collection,ctx}
}

func (a * Add)Add(m *models.Add)(*mongo.InsertOneResult){
	m.Sum = m.FirstNumber+m.SecondNumber
	res,err:=a.Collection.InsertOne(a.ctx,m)
	if err != nil {
		return nil
	}
	return res
}