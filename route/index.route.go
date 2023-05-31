package route

import (
	"rest-go/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/", handler.UserHandlerRead)
}
