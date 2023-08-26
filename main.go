package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"restapi.com/my-apis/controller"
	"restapi.com/my-apis/services"
)

var (
	server         *gin.Engine
	userservices   services.UserServices
	usercontroller controller.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongopclient   *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo client connected ")

	usercollection = mongoclient.Database("userdb").Collection("users")
	userservices = services.NewUserService(usercollection, ctx)
	usercontroller = controller.New(userservices)
	server = gin.Default()

}
func main() {

	defer mongopclient.Disconnect(ctx)
	basepath := server.Group("/v1")

	usercontroller.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))
}
