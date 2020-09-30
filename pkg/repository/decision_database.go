package repository

import (
	"context"

	"github.com/TudorHulban/echotest/pkg/models"
)

const collectionName = "decisions"

// DecisionDatabase represents the collection
type DecisionDatabase interface {
	FindAll(context.Context) (*[]models.Decision, error)
	FindOne(context.Context, interface{}) (*models.Decision, error)
	Create(context.Context, *models.Decision) error
	DeleteByName(context.Context, string) error
}

type decisionDatabase struct {
	db DatabaseHelper
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

func (u *decisionDatabase) DeleteByName(ctx context.Context, name string) error {
	// In this case it is possible to use bson.M{"username":username} but I tend
	// to avoid another dependency in this layer and for demonstration purposes
	// used omitempty in the model
	decision := &models.Decision{
		Name: name,
	}
	_, err := u.db.Collection(collectionName).DeleteOne(ctx, decision)
	return err
}
