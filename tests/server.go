package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmarcha/ent-goswagger-app/internal/utils"
)

func main() {

	names := []string{"test", "student", "tutor", "calendar", "admin"}
	i := -1

	r := gin.Default()
	r.GET("/oauth/authorize", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "http://localhost:5000/v2/auth/callback?code=42&state="+c.Query("state"))
	})
	r.POST("/oauth/token", func(c *gin.Context) {
		i++
		c.Header("Cache-Control", "no-store")
		c.JSON(200, gin.H{
			"access_token":  utils.RandomString(64),
			"token_type":    "Bearer",
			"expires_in":    3600,
			"refresh_token": utils.RandomString(64),
			"scope":         "public",
		})
	})
	r.GET("/v2/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"login":      names[i%len(names)],
			"first_name": "test",
			"last_name":  "test",
			"image_url":  "test",
		})
	})
	_ = r.Run("0.0.0.0:8080")
}
