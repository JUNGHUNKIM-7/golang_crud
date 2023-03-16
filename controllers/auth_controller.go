package controllers

import (
	"net/http"
	"sync"
	"time"

	"runner/src/db"
	"runner/src/middleware"
	"runner/src/models"

	"github.com/gin-gonic/gin"
)



func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var wg sync.WaitGroup
		var atToken string
		var err error

		wg.Add(2)
		if err = c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusNotAcceptable, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}

		if user := db.GetUser(user.Email); user == nil {
			c.JSON(http.StatusNotFound, models.ErrResponse{
				Message: "user not exist",
			})
			return
		}

		go func() {
			atToken, err = middleware.GenerateToken(user.Email, time.Hour*1)
			if err != nil {
				c.JSON(http.StatusBadRequest, models.ErrResponse{
					Message: err.Error(),
				})
				return
			}
			wg.Done()
		}()

		go func() {
			rtToken, err := middleware.GenerateToken(user.Email, time.Hour*24*30)
			if err != nil {
				c.JSON(http.StatusBadRequest, models.ErrResponse{
					Message: err.Error(),
				})
				return
			}
			err = db.UpdateOne(user.Email, map[string]any{"rt_token": rtToken})
			if err != nil {
				c.JSON(http.StatusBadRequest, models.ErrResponse{
					Message: err.Error(),
				})
				return
			}
			wg.Done()
		}()
		wg.Wait()

		c.SetCookie("at", atToken, 3600, "/", "localhost", false, true)

		c.JSON(http.StatusOK, models.SignedResponse{
			Message: "ok",
		})
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail, err := middleware.GetEmailFromCookie(c, "at")
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}

		err = db.UpdateOne(userEmail, map[string]any{"rt_token": ""})
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, models.OkResponse{
			Message: "ok",
		})
	}
}
