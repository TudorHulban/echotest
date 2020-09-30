package logic

import (
	"time"
)

func DecisionAmount(amo int) bool {
	if amo < 5000 {
		time.Sleep(time.Duration(2) * time.Second)
		return true
	}

	if amo >= 5000 && amo <= 10000 {
		return false
	}

	time.Sleep(time.Duration(3) * time.Second)
	return false
}
