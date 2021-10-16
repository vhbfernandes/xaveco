package controllers

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

// RouterSetup creates a setup for the gin controllers
func RouterSetup() *gin.Engine {
	r := gin.New()
	r.Use(ginlogrus.Logger(log.StandardLogger()), gin.Recovery())

	//setup prometheus before adding routes
	setupPrometheus(r)

	r.GET("/healthz", healthsHealth())
	r.GET("/xavecos", xavecoGet())
	r.POST("/xavecos", xavecoCreate())

	return r
}

func setupPrometheus(router *gin.Engine) {
	prometheus := ginprometheus.NewPrometheus("gin")
	prometheus.Use(router)
}
