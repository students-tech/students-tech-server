package assignment

import (
	"github.com/gofiber/fiber/v2"
	"students-tech-server/interfaces"
	"students-tech-server/shared"
)

type Controller struct {
	Shared     shared.Holder
	Interfaces interfaces.Holder
}

func (c *Controller) Routes(app *fiber.App) {

}

func NewController(shared shared.Holder, interfaces interfaces.Holder) Controller {
	return Controller{
		Shared:     shared,
		Interfaces: interfaces,
	}
}
