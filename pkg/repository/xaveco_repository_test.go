package repository

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	log "github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"github.com/vhbfernandes/xaveco/pkg/models"
	"os"
	"testing"
)

const SeedModel = `{
    "content": "xaveco teste",
    "tags": ["pedreiro", "outra tag"]
  }`

var cfg *models.Xaveco
var ctx context.Context
var xvc *XavecoMongoRepository

func TestItShouldSaveAModel(t *testing.T) {
	err := xvc.Create(ctx, cfg)
	if err != nil {
		t.Errorf("Error saving model %v", err)
	}
}

func TestItShouldReturnAllSavedItems(t *testing.T) {
	items, err := xvc.FindAll(ctx)
	if len(items) < 1 {
		t.Errorf("Collection size mismatch, should be at least one")
	}
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestItShouldReturnARandomItem(t *testing.T) {
	item, err := xvc.FindRandom(ctx, "any")
	if item == nil {
		t.Errorf("Item mismatch, should be at least one")
	}
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestItShouldReturnAllSavedItemsOnSameTag(t *testing.T) {
	items, err := xvc.FindByTag(ctx,"pedreiro")
	if len(items) < 1 {
		t.Errorf("Collection size mismatch, should be at least one")
	}
	if err != nil {
		t.Errorf(err.Error())
	}
	if !funk.Contains(items[0].Tags, "pedreiro") {
		t.Errorf("Something really weird happened")
	}
}

func TestItShouldNotFindNonExistingItems(t *testing.T) {
	items, err := xvc.FindByTag(ctx,"salada")
	if len(items) > 0 {
		t.Errorf("Collection size mismatch, there should be no salada xavecos")
	}
	if err != nil {
		t.Errorf(err.Error())
	}
}


func setup() {
	err := json.Unmarshal([]byte(SeedModel), &cfg)
	ctx = &gin.Context{}
	xvc = NewXavecoMongoRepository()
	if err != nil {
		panic("Error creating model " + err.Error())
	}
}

func shutdown() {
	_,client,_, err := mgm.DefaultConfigs()
	err = client.Database("test").Drop(context.TODO())
	if err != nil {
		log.Errorf("Error shutting down tests %v", err)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
