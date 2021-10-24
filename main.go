package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vhbfernandes/xaveco/pkg/controllers"
	"os"
)

func main() {
	setupLogger()
	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Panicf("PORT environment variable is not set")
	}
	router := controllers.RouterSetup()
	log.Fatalf("Error starting webserver %v\n", router.Run(":"+port))
}

func setupLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.TraceLevel)
}
