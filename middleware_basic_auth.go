package main

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/raj47i/IntTestPlivo/models"
)

func basicAuthMiddleware(c *gin.Context) {
	realm := "Basic realm=" + strconv.Quote("Authorization Required")
	// Search user in the slice of allowed credentials
	auth := strings.SplitN(c.GetHeader("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		c.Header("WWW-Authenticate", realm)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "",
			"error":   "Unauthorized",
		})
		c.Abort()
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 {
		c.Header("WWW-Authenticate", realm)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "",
			"error":   "Unauthorized",
		})
		c.Abort()
		return
	}
	var acc models.Account
	if !acc.LoadByUserName(pair[0]) || !acc.Authenticate(pair[1]) {
		c.Header("WWW-Authenticate", realm)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "",
			"error":   "Unauthorized",
		})
		c.Abort()
		return
	}
	c.Set("AccountID", acc.ID)
	c.Next()
}
