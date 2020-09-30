package repository

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/TudorHulban/echotest/pkg/models"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/suite"
	containers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson"
)

type TestSuite struct {
	suite.Suite
	mongoImage containers.Container
	ctx        context.Context
	database   DecisionDatabase
}

func (s *TestSuite) SetupSuite() {
	var err error
	s.ctx = context.Background()
	req := containers.ContainerRequest{
		Image:        "mongo",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForHTTP("/"),
	}
	s.mongoImage, err = containers.GenericContainer(s.ctx, containers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	port, err := nat.NewPort("tcp", "27017")
	firstMappedPort, err := s.mongoImage.MappedPort(s.ctx, port)
	if err != nil {
		s.Error(err)
	}
	config := &DBConfig{DatabaseName: "decisions_test", DBUrl: fmt.Sprintf("mongodb://localhost:%s", firstMappedPort.Port())}
	helper, err := NewClient(config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = helper.Connect()
	if err != nil {
		log.Fatal(" Cound not connect to mongo {} ", err.Error())
	}
	dbHelper := NewDatabase(config, helper)
	s.database = NewDecisionDatabase(dbHelper)
}

func (s *TestSuite) AfterTest(suiteName, testName string) {
	if testName == "TestDeleteByRequestID" {
		s.mongoImage.Terminate(s.ctx)
	}
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestCreateAndFindOne() {
	decision1 := &models.Decision{RequestID: "1", Name: "David", Amount: 1500, Answer: true}
	err := s.database.Create(s.ctx, decision1)
	s.Assert().Nil(err, "Error should be nil")
	res, err := s.database.FindOne(s.ctx, bson.D{{"requestid", "1"}})
	isNil := s.Assert().Nil(err, "Error should be nil")
	if !isNil {
		log.Fatalf("not nil!")
	}
	s.Assert().Equal(res.Amount, 1500, "Recorded amount should match")
	s.Assert().Nil(s.database.DeleteByRequestID(s.ctx, "1"), "Error should be nil")
}

func (s *TestSuite) TestFindAll() {
	/*
		decision1 := &models.Decision{RequestID: "1", Name: "David", Amount: 1500, Answer: true}
		decision2 := &models.Decision{RequestID: "2", Name: "Gibran", Amount: 2500, Answer: false}
		decision3 := &models.Decision{RequestID: "3", Name: "Tudor", Amount: 3500, Answer: true}
		s.database.Create(s.ctx, decision1)
		s.database.Create(s.ctx, decision2)
		s.database.Create(s.ctx, decision3)

		res, err := s.database.FindAll(s.ctx)
		s.Assert().Nil(err, "Error should be nil")
		s.Assert().True(len(*res) == 3, "There should be 3 elements") */
}

func (s *TestSuite) TestDeleteByRequestID() {

}
