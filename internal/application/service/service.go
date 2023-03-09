package service

import (
	"concurrencybenchmark/internal/domain"
	"concurrencybenchmark/internal/infra"
)

const maxWorkers = 5

func NotifyConcurrently(u []domain.User) {
	toNotify := make(chan domain.User, len(u))

	for i := 0; i < maxWorkers; i++ {
		go sendUserToExternalAPI(toNotify)
	}

	for _, usr := range u {
		toNotify <- usr
	}

	close(toNotify)
}

func NotifyDefault(u []domain.User) {
	for _, usr := range u {
		infra.DoRequestToExternalAPI(usr)
	}
}

func sendUserToExternalAPI(users <-chan domain.User) {
	for u := range users {
		infra.DoRequestToExternalAPI(u)
	}
}
