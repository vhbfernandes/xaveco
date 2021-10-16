package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vhbfernandes/xaveco/pkg/models"
	"github.com/vhbfernandes/xaveco/pkg/repository"
	"net/http"
	"os"
)

func xavecoGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		tag := c.Query("tag")
		var data map[string]interface{}
		var err error
		if len(tag) < 1 {
			data, err = repository.FindRandom(c, "any")
		} else {
			data, err = repository.FindRandom(c, tag)
		}
		if err != nil {
			log.Errorf("Error finding data %v", err)
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

func xavecoCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("AUTH_TOKEN") {
			c.JSON(http.StatusUnauthorized, "")
			return
		}
		var xaveco models.Xaveco
		if err := c.ShouldBindJSON(&xaveco); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repository.Create(c, &xaveco); err != nil {
			log.Errorf("Error writing to database %s", err)
			c.JSON(http.StatusInternalServerError, "")
			return
		}
		c.JSON(http.StatusCreated, &xaveco)
	}
}

