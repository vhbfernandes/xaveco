package controllers

import (
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/mcuadros/go-gin-prometheus"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
	"github.com/vhbfernandes/xaveco/pkg/repository"
)

type server struct {
	repo repository.Repository
}

// RouterSetup creates a setup for the gin controllers
func RouterSetup() *gin.Engine {
	r := gin.New()
	r.Use(ginlogrus.Logger(log.StandardLogger()), gin.Recovery())

	//setup prometheus before adding routes
	setupPrometheus(r)
	s := setupServer()

	r.GET("/healthz", s.healthsHealth())
	r.GET("/xavecos", s.xavecoGet())
	r.POST("/xavecos", s.xavecoCreate())

	return r
}

func setupPrometheus(router *gin.Engine) {
	prometheus := ginprometheus.NewPrometheus("gin")
	prometheus.Use(router)
}

func setupServer() *server {
	return &server{
		repo: repository.NewXavecoMongoRepository(),
	}
}