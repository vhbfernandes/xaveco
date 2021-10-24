package controllers

import "github.com/gin-gonic/gin"

func (s *server) healthsHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "OK")
	}
}
