package models

//DBResult is a generic type to hold databse results
type DBResult struct{}

//Decision is the structure that holds user values
type Decision struct {
	DBResult
	Name   string
	Amount int
}
