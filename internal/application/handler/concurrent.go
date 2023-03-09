package handler

import (
	"concurrencybenchmark/internal/application/service"
	"concurrencybenchmark/internal/domain"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func Concurrent(c *fiber.Ctx) error {
	u := domain.GetUsersToNotify()

	service.NotifyConcurrently(u)

	j, _ := json.Marshal(u)

	return c.JSON(map[string]string{
		"notified_users": string(j),
	})
}
