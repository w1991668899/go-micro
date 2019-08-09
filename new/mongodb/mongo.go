package mongodb

import (
	"go-micro/golib/lib/lib_config"
	"go-micro/golib/lib/lib_mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client


func InitMongo(configMongoDb lib_config.ConfMongoDb)  {
	Client = lib_mongodb.CreateMongoDb(configMongoDb)
}
