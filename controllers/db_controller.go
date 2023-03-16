package controllers

import (
	"net/http"
	"runner/src/db"
	"runner/src/middleware"
	"runner/src/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Param("email")
		if user := db.GetUser(params); user != nil {
			return
		}
		c.JSON(http.StatusNotFound, models.ErrResponse{
			Message: "user not found",
		})
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users := db.GetAllUser()
		c.JSON(http.StatusOK, models.OkResponse{
			Users:   users,
			Message: "ok",
		})
	}
}

func Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body models.User
		var err error

		if err = c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusNotImplemented, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}

		err = db.Insert(&body)
		if err != nil {
			c.JSON(http.StatusNotImplemented, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, models.OkResponse{
			Message: "ok",
		})
	}
}

func BulkAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := db.BulkAdd()
		if err != nil {
			c.JSON(http.StatusNotImplemented, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, models.OkResponse{
			Message: "ok",
		})
	}
}

func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := db.Delete(c.Param("email"))
		if err != nil {
			c.JSON(http.StatusNotFound, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, models.OkResponse{
			Message: "ok",
		})
	}
}

func UpdateOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := db.UpdateOne(c.Param("email"), map[string]any{"email": "update@email.com"})
		if err != nil {
			c.JSON(http.StatusNotFound, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, models.OkResponse{
			Message: "ok",
		})
	}
}

func DeleteAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := db.DeleteAll()
		if err != nil {
			c.JSON(http.StatusNotImplemented, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, models.OkResponse{
			Message: "ok",
		})
	}
}

func GenTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := middleware.GenerateToken("test@gmail.com", time.Minute*1)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.UnsignedResponse{
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, models.SignedResponse{
			Token:   token,
			Message: "ok",
		})
	}
}
