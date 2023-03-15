package controllers

import (
	"net/http"
	"runner/src/db"
	model "runner/src/models"

	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Param("email")
		if user := db.GetUser(params); user != nil {
			c.JSON(200, gin.H{
				"ok": user,
			})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"users": "not found",
		})
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users := db.GetAllUser()
		c.JSON(200, gin.H{
			"ok": users,
		})
	}
}

func Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.User

		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{
				"err": err.Error(),
			})
			return
		}

		err = db.Insert(&body)
		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
		})
	}
}

func BulkAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := db.BulkAdd()
		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"status": "ok",
		})
	}
}

func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := db.Delete(c.Param("email"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func UpdateOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := db.UpdateOne(c.Param("email"), map[string]any{"email": "update@email.com"})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func DeleteAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := db.DeleteAll()
		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}
