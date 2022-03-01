package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gmarcha/ent-goswagger-app/internal/utils"
)

func main() {

	names := []string{"azir", "cassiopeia", "lissandra", "corki"}
	i := -1

	r := gin.Default()
	r.GET("/oauth/authorize", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "http://localhost:5000/v2/auth/callback?code=42&state="+c.Query("state"))
	})
	r.POST("/oauth/token", func(c *gin.Context) {
		i++
		c.Header("Cache-Control", "no-store")
		c.JSON(200, gin.H{
			"access_token":  names[i],
			"token_type":    "Bearer",
			"expires_in":    3600,
			"refresh_token": utils.RandomString(64),
			"scope":         "public",
		})
	})
	r.GET("/v2/me", func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		token := strings.Split(bearerToken, " ")[1]
		c.JSON(200, gin.H{
			"login":      token,
			"first_name": "melissa",
			"last_name":  "melissa",
			"image_url":  "none",
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
