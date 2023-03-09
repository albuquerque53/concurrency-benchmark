package infra

import (
	"concurrencybenchmark/internal/domain"
	"encoding/json"
	"log"
	"time"
)

type request struct {
	Target string `json:"target"`
	User   domain.User
}

func DoRequestToExternalAPI(u domain.User) error {
	req := request{
		Target: "https://example-api.com/",
		User:   u,
	}
	logRequest(req)

	time.Sleep(time.Second * 3)
	return nil
}

func logRequest(req request) {
	j, _ := json.Marshal(req)

	log.Printf("Executing request: %s\n", j)
}
