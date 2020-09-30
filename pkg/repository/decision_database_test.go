package repository

/*
import (
	"context"
	"errors"
	"testing"

	"github.com/TudorHulban/echotest/pkg/models"
	"github.com/TudorHulban/echotest/pkg/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFindOne(t *testing.T) {

	// Define variables for interfaces
	var dbHelper DatabaseHelper
	var collectionHelper CollectionHelper
	var srHelperErr SingleResultHelper
	var srHelperCorrect SingleResultHelper

	// Set interfaces implementation to mocked structures
	dbHelper = &mock.DatabaseHelper{}
	collectionHelper = &mocks.CollectionHelper{}
	srHelperErr = &mocks.SingleResultHelper{}
	srHelperCorrect = &mocks.SingleResultHelper{}

	// Because interfaces does not implement mock.Mock functions we need to use
	// type assertion to mock implemented methods
	srHelperErr.(*mocks.SingleResultHelper).
		On("Decode", mock.AnythingOfType("*models.Decision")).
		Return(errors.New("mocked-error"))

	srHelperCorrect.(*mocks.SingleResultHelper).
		On("Decode", mock.AnythingOfType("*models.Decision")).
		Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Decision)
		arg.Username = "mocked-user"
	})

	collectionHelper.(*mocks.CollectionHelper).
		On("FindOne", context.Background(), bson.M{"error": true}).
		Return(srHelperErr)

	collectionHelper.(*mocks.CollectionHelper).
		On("FindOne", context.Background(), bson.M{"error": false}).
		Return(srHelperCorrect)

	dbHelper.(*mocks.DatabaseHelper).
		On("Collection", "users").Return(collectionHelper)

	// Create new database with mocked Database interface
	userDba := NewDecisionDatabase(dbHelper)

	// Call method with defined filter, that in our mocked function returns
	// mocked-error
	user, err := userDba.FindOne(context.Background(), bson.M{"error": true})

	assert.Empty(t, user)
	assert.EqualError(t, err, "mocked-error")

	// Now call the same function with different different filter for correct
	// result
	user, err = userDba.FindOne(context.Background(), bson.M{"error": false})

	assert.Equal(t, &models.Decision{Name: "mocked-user"}, user)
	assert.NoError(t, err)
} */
