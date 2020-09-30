package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TudorHulban/echotest/pkg/models"
	"github.com/TudorHulban/echotest/pkg/repository"
)

func main() {

	dbHelper := repository.GetInstance()
	ctx := context.Background()
	decision1 := &models.Decision{Name: "David", Amount: 1500}
	decision2 := &models.Decision{Name: "Gibran", Amount: 2500}
	decision3 := &models.Decision{Name: "Tudor", Amount: 3500}
	dbHelper.Create(ctx, decision1)
	dbHelper.Create(ctx, decision2)
	dbHelper.Create(ctx, decision3)
	res, err := dbHelper.FindAll(ctx)
	if err != nil {
		log.Fatal("Error retrieving values {}", err.Error())
	}
	fmt.Println(res)
}
