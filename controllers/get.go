package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OKAY",
	})
}

func GetData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Here is your data",
	})
}
