package controllers

import "github.com/gin-gonic/gin"

//simple healthcheck, not really suitable for production
func healthsHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "OK")
	}
}
