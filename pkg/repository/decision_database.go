package repository

import (
	"context"
	"log"

	"github.com/TudorHulban/echotest/pkg/models"
)

const collectionName = "decisions"

// DecisionDatabase represents the collection
type DecisionDatabase interface {
	FindAll(context.Context) (*[]models.Decision, error)
	FindOne(context.Context, interface{}) (*models.Decision, error)
	Create(context.Context, *models.Decision) error
	DeleteByRequestID(context.Context, string) error
}

type decisionDatabase struct {
	db DatabaseHelper
}

var currentInstance DecisionDatabase

// GetInstance provides a working instance
func GetInstance() DecisionDatabase {
	if currentInstance == nil {
		config := &DBConfig{DatabaseName: "decisions", DBUrl: "mongodb://localhost:27017"}
		helper, err := NewClient(config)
		if err != nil {
			log.Fatalf(err.Error())
		}

		err = helper.Connect()
		if err != nil {
			log.Fatal(" Cound not connect to mongo {} ", err.Error())
		}
		dbHelper := NewDatabase(config, helper)
		currentInstance = NewDecisionDatabase(dbHelper)
	}
	return currentInstance
}

// NewDecisionDatabase creates a new instance
func NewDecisionDatabase(db DatabaseHelper) DecisionDatabase {
	return &decisionDatabase{
		db: db,
	}
}

func (u *decisionDatabase) FindAll(ctx context.Context) (*[]models.Decision, error) {
	var result *[]models.Decision
	result, err := u.db.Collection(collectionName).FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *decisionDatabase) FindOne(ctx context.Context, filter interface{}) (*models.Decision, error) {
	decision := &models.Decision{}
	err := u.db.Collection(collectionName).FindOne(ctx, filter).Decode(decision)
	if err != nil {
		return nil, err
	}
	return decision, nil
}

func (u *decisionDatabase) Create(ctx context.Context, decision *models.Decision) error {
	_, err := u.db.Collection(collectionName).InsertOne(ctx, decision)
	return err
}

func (u *decisionDatabase) DeleteByRequestID(ctx context.Context, requestID string) error {
	decision := &models.Decision{
		RequestID: requestID,
	}
	_, err := u.db.Collection(collectionName).DeleteOne(ctx, decision)
	return err
}
