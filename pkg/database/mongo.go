package database

import (
	"github.com/kamva/mgm/v3"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

// Connect creates a mongodb connection and returns a Mongo struct
func Connect() {
	err := mgm.SetDefaultConfig(nil, os.Getenv("MONGODB_DATABASE"), options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))
	if err != nil {
		log.Errorf("Database connection error: %v", err)
	}
}
