package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vhbfernandes/xaveco/pkg/controllers"
	"github.com/vhbfernandes/xaveco/pkg/repository"
	"os"
)

func main() {
	setupLogger()
	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Panicf("PORT environment variable is not set")
	}
	setupRepository()
	router := controllers.RouterSetup()
	log.Fatalf("Error starting webserver %v\n", router.Run(":"+port))
}

func setupRepository() {
	repository.Init()
}

func setupLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.TraceLevel)
}
