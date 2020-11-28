package main

import (
	"admin-rt/config"
	"admin-rt/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// db := config.GetDB()
	//when server.go start, it will be run function InitDB (connecting to database)
	// config.InitDB()
	config.GetDB()

	router := gin.Default()

	v1 := router.Group("api/v1") // /api/v1
	{
		v1.POST("/login", routes.Login) // /api/v1/login
		adminrt := v1.Group("/admin")   // /api/v1/sewaAset
		{
			adminrt.GET("/account", routes.GetAccount)   // /api/v1/admin/account
			adminrt.POST("/account", routes.PostAccount) // /api/v1/admin/account
		}
	}

	router.Run()
}
