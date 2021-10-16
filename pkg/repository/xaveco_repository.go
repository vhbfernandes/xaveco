package repository

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/vhbfernandes/xaveco/pkg/database"
	"github.com/vhbfernandes/xaveco/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var coll *mgm.Collection
var ctx context.Context
// Init Initializes repository;
// mongodb connection is required
func Init() {
	database.Connect()
	ctx = mgm.Ctx()
	coll = mgm.Coll(&models.Xaveco{})
	ensureIndexes()
}
// FindRandom returns a random xaveco, tagged one if a tag is provided
func FindRandom(ctx *gin.Context, tag string) (xvc map[string]interface{}, err error) {
	var pipe mongo.Pipeline
	var result []*models.Xaveco
	var xaveco = make(map[string]interface{})
	if tag == "any" {
		pipe = mongo.Pipeline{
			{{"$sample", bson.M{"size": 1}}},
		}
	} else {
		pipe = mongo.Pipeline{
			{{"$match", bson.M{"tags": tag}}},
			{{"$sample", bson.M{"size": 1}}},
		}
	}

	cursor, err := coll.Aggregate(ctx,pipe)
	err = cursor.All(ctx,&result)
	//todo arrumar essa gambi
	xaveco["data"] = result[0].Content
	xaveco["tags"] = result[0].Tags
	xvc = xaveco
	return
}

// FindAll returns all xavecos on database
func FindAll(ctx *gin.Context) (xavecos []*models.Xaveco, err error) {
	cursor, err := coll.Find(ctx, bson.M{})
	err = cursor.All(ctx, &xavecos)
	return
}

// FindByTag returns all xavecos declared with the same tag
func FindByTag(ctx *gin.Context, tag string) (xavecos []*models.Xaveco, err error) {
	cursor, err := coll.Find(ctx, bson.M{"tags": tag})
	err = cursor.All(ctx, &xavecos)
	return
}

// Create saves a model xaveco to the database
func Create(ctx *gin.Context, xaveco *models.Xaveco) error {
	return coll.CreateWithCtx(ctx, xaveco)
}

func ensureIndexes() {
	coll.Indexes().CreateOne(ctx, mongo.IndexModel{Keys: bson.D{{"tags", 1}}}, &options.CreateIndexesOptions{})
}