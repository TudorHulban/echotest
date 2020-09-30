package logic

import (
	"time"
)

func DecisionAmount(amount uint64) bool {
	if amount < uint64(5000) {
		time.Sleep(time.Duration(2) * time.Second)
		return true
	}

	if amount >= 5000 && amount <= 10000 {
		return false
	}

	time.Sleep(time.Duration(3) * time.Second)
	return false
}
