package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	originsString := "http://localhost:3000" //Definindo o front
	allowedOrigins := strings.Split(originsString, ",")

	// Return the middleware handler function
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Check if the origin is allowed
		if isOriginAllowed(origin, allowedOrigins) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		}

		// Handle preflight OPTIONS requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Function to check if a given origin is allowed
func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}
