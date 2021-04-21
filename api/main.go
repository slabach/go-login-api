package main

import (
	"github.com/gin-gonic/gin"
	"schnelllegal.com/api/models"
	"schnelllegal.com/api/routes"
)

func main() {
	// init list of users
	user := models.NewUserList()

	// init gin server
	r := gin.Default()

	// routes
	api := r.Group("/api")
	{
		api.POST("/login", routes.LoginReq(user))
		api.POST("/register", routes.RegisterReq(user))
	}

	// listen and serve on localhost:5000
	r.Run("0.0.0.0:5000")
}
