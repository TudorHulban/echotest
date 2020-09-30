package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TudorHulban/echotest/pkg/models"
	"github.com/TudorHulban/echotest/pkg/repository"
)

func main() {

	config := &repository.DBConfig{DatabaseName: "decisions", DBUrl: "mongodb://localhost:27017"}
	helper, err := repository.NewClient(config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = helper.Connect()
	if err != nil {
		log.Fatal(" Cound not connect to mongo {} ", err.Error())
	}
	dbHelper := repository.NewDatabase(config, helper)
	ctx := context.Background()
	dbHelper.Client().StartSession()
	decision1 := &models.Decision{Name: "David", Amount: 1500}
	decision2 := &models.Decision{Name: "Gibran", Amount: 2500}
	decision3 := &models.Decision{Name: "Tudor", Amount: 3500}
	dbHelper.Collection("decisions").InsertOne(ctx, decision1)
	dbHelper.Collection("decisions").InsertOne(ctx, decision2)
	dbHelper.Collection("decisions").InsertOne(ctx, decision3)
	res, err := dbHelper.Collection("decisions").FindAll(ctx)
	fmt.Println(res)
}
