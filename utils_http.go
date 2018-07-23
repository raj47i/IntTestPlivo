package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func http405(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"message": "",
		"error":   "Method Not Allowed",
	})
}

func http500(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "",
		"error":   err.Error(),
	})
}
func http422(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": "",
		"error":   err.Error(),
	})
}
func http403(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, gin.H{
		"message": "",
		"error":   err.Error(),
	})
}
func http200(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"error":   "",
	})
}
