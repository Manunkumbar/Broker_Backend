package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func InitCircuitBreaker() {
	settings := gobreaker.Settings{
		Name:        "HTTP Circuit Breaker",
		MaxRequests: 3,
		Interval:    10 * time.Second,
		Timeout:     5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 5
		},
	}
	cb = gobreaker.NewCircuitBreaker(settings)
}

func GetWithCircuitBreaker(url string) ([]byte, error) {
	body, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("upstream error: %s", resp.Status)
		}

		return ioutil.ReadAll(resp.Body)
	})

	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}
