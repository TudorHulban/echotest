package logic

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	url0 = "https://httpstat.us/200?sleep=50"
	url1 = "https://httpstat.us/200?sleep=5000"
	url2 = "https://httpstat.us/400?sleep=5000"
)

func DecisionAmount(amo int) (bool, error) {
	if amo < 5000 {
		time.Sleep(time.Duration(2) * time.Second)
		return true, nil
	}

	if amo >= 5000 && amo <= 10000 {
		random := generateRandomNo(0, 500)

		parseCode := func(code int) bool {
			if code < 400 {
				return true
			}
			return false
		}

		if random <= 250 {
			statusCode, errGetURL1 := makeGETRequest(url1)
			if errGetURL1 != nil {
				return false, errors.WithMessagef(errGetURL1, "error in decision amount accesing: %v", url1)
			}

			return parseCode(statusCode), nil
		}

		statusCode, errGetURL2 := makeGETRequest(url2)
		if errGetURL2 != nil {
			return false, errors.WithMessagef(errGetURL2, "error in decision amount accesing: %v", url2)
		}

		return parseCode(statusCode), nil
	}

	time.Sleep(time.Duration(3) * time.Second)
	return false, nil
}

func generateRandomNo(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// makeGETRequest Helper makes GET request and returns HTTP code.
func makeGETRequest(url string) (int, error) {
	log.Println("making GET request for: ", url)

	resp, errGet := http.Get(url)
	if errGet != nil {
		return 0, errGet
	}

	return resp.StatusCode, nil
}
