package models

//Decision is the structure that holds user values
type Decision struct {
	RequestID string
	Name      string
	Amount    int
	Answer    bool
}
