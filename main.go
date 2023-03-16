package main

import (
	"net/http"
	controller "runner/src/controllers"
	"runner/src/db"
	"runner/src/middleware"
	"runner/src/models"

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
		auth.POST("login", controller.Login())
		auth.GET("token", controller.GenTokenHandler()) // debug
		auth.GET("users", controller.GetUsers())        // debug
		auth.DELETE("debug", controller.DeleteAll())    // debug
	}

	authorized := r.Group("/").Use(middleware.CheckJwtFromCookie)
	{
		authorized.GET("", func(c *gin.Context) {
			token, err := middleware.GetEmailFromCookie(c, "at")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, models.OkResponse{
					Message: "failed to get token",
				})
				return
			}

			c.JSON(http.StatusOK, models.OkResponse{
				Message: token,
			})
		})
		authorized.POST("logout", controller.Logout())
	}

	r.Run(":9090")
}
