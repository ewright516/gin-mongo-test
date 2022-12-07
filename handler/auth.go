package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SignIn(c *mongo.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {

	}

	return gin.HandlerFunc(fn)
}
