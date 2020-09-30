package repository

import (
	"context"

	"github.com/TudorHulban/echotest/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseHelper interface {
	Collection(name string) CollectionHelper
	Client() ClientHelper
}

type CollectionHelper interface {
	FindAll(context.Context) (*[]models.Decision, error)
	FindOne(context.Context, interface{}) SingleResultHelper
	InsertOne(context.Context, interface{}) (interface{}, error)
	DeleteOne(ctx context.Context, filter interface{}) (int64, error)
}

type SingleResultHelper interface {
	Decode(v interface{}) error
}

type DBResultHelper interface {
	Decode(v interface{}) error
}

type ClientHelper interface {
	Database(string) DatabaseHelper
	Connect() error
	StartSession() (mongo.Session, error)
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoSession struct {
	mongo.Session
}

//NewClient creates a client
func NewClient(cnf *DBConfig) (ClientHelper, error) {
	c, err := mongo.NewClient(options.Client().ApplyURI(cnf.DBUrl))

	return &mongoClient{cl: c}, err

}

// NewDatabase creates the database
func NewDatabase(cnf *DBConfig, client ClientHelper) DatabaseHelper {
	return client.Database(cnf.DatabaseName)
}

func (mc *mongoClient) Database(dbName string) DatabaseHelper {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

func (mc *mongoClient) Connect() error {
	// mongo client does not use context on connect method. There is a ticket
	// with a request to deprecate this functionality and another one with
	// explanation why it could be useful in synchronous requests.
	// https://jira.mongodb.org/browse/GODRIVER-1031
	// https://jira.mongodb.org/browse/GODRIVER-979
	return mc.cl.Connect(nil)
}

func (md *mongoDatabase) Collection(colName string) CollectionHelper {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() ClientHelper {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) FindAll(ctx context.Context) (*[]models.Decision, error) {
	var result []models.Decision
	cursor, err := mc.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResultHelper {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}
	return id.InsertedID, err
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.coll.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count.DeletedCount, err
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}
