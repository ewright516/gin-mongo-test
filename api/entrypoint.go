package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ewright516/gin-mongo-test/handler"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var app *gin.Engine
var client *mongo.Client
var err error

func init() {
	conn := os.Getenv("MONGO_URL")
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(conn))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	app = gin.New()
	router := app.Group("/api")
	auth := router.Group("/auth")
	data := router.Group("/data")

	auth.POST("/signin", handler.SignIn(client))
	auth.POST("/signup", handler.SignUp(client))

	data.GET("/post/:id", handler.GetPostById(client))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
