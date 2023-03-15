package main

import (
	controller "runner/src/controllers"
	db "runner/src/db"

	"github.com/gin-gonic/gin"
)

func init() {
	db.InitializeDb()
}

func main() {
	defer db.Client.Disconnect(db.Ctx)

	r := gin.Default()
	auth := r.Group("auth")
	{
		auth.POST("register", controller.Insert())
		auth.GET("users", controller.GetUsers())
		auth.GET("user/:email", controller.GetUser())
		auth.PUT("update/:email", controller.UpdateOne())
		auth.DELETE("delete/:email", controller.Delete())
		auth.DELETE("debug", controller.DeleteAll())
	}
	r.Run(":9090")
}
