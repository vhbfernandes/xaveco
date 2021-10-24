package repository

import (
	"context"
	"github.com/kamva/mgm/v3"
	"github.com/vhbfernandes/xaveco/pkg/database"
	"github.com/vhbfernandes/xaveco/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type XavecoMongoRepository struct {
	coll *mgm.Collection
}

// NewXavecoMongoRepository Initializes repository;
// mongodb connection is required
func NewXavecoMongoRepository() *XavecoMongoRepository {
	database.Connect()
	xvc := XavecoMongoRepository{
		coll: mgm.Coll(&models.Xaveco{}),
	}
	xvc.ensureIndexes()
	return &xvc
}
// FindRandom returns a random xaveco, tagged one if a tag is provided
func (x *XavecoMongoRepository) FindRandom(ctx context.Context, tag string) (xvc map[string]interface{}, err error) {
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

	cursor, err := x.coll.Aggregate(ctx,pipe)
	err = cursor.All(ctx,&result)
	//todo arrumar essa gambi
	xaveco["data"] = result[0].Content
	xaveco["tags"] = result[0].Tags
	xvc = xaveco
	return
}

// FindAll returns all xavecos on database
func (x *XavecoMongoRepository) FindAll(ctx context.Context) (xavecos []*models.Xaveco, err error) {
	cursor, err := x.coll.Find(ctx, bson.M{})
	err = cursor.All(ctx, &xavecos)
	return
}

// FindByTag returns all xavecos declared with the same tag
func (x *XavecoMongoRepository) FindByTag(ctx context.Context, tag string) (xavecos []*models.Xaveco, err error) {
	cursor, err := x.coll.Find(ctx, bson.M{"tags": tag})
	err = cursor.All(ctx, &xavecos)
	return
}

// Create saves a model xaveco to the database
func (x *XavecoMongoRepository) Create(ctx context.Context, xaveco *models.Xaveco) error {
	return x.coll.CreateWithCtx(ctx, xaveco)
}

func (x *XavecoMongoRepository) ensureIndexes() {
	x.coll.Indexes().CreateOne(context.TODO(), mongo.IndexModel{Keys: bson.D{{"tags", 1}}}, &options.CreateIndexesOptions{})
}