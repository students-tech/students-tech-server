package students

import (
	"students-tech-server/interfaces"
	"students-tech-server/shared"
	"students-tech-server/shared/common"
	"students-tech-server/shared/dto"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
  Shared shared.Holder
  Interfaces interfaces.Holder
}

func (c *Controller) Routes(app *fiber.App) {
  students := app.Group("/students")
  students.Post("/", c.register)
}

// All godoc
// @Tags Students
// @Summary Register new students
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.CreateStudentsRequest
// @Failure 200 {array} dto.CreateStudentsRequest
// @Router /students/ [post]
func (c *Controller) register(ctx *fiber.Ctx) error {
  var (
    req dto.CreateStudentsRequest
  ) 

  err := common.DoCommonRequest(ctx, &req)
  if err != nil {
    return common.DoCommonErrorResponse(ctx, err)
  }

  return common.DoCommonSuccessResponse(ctx, req)
}

func NewController(shared shared.Holder, interfaces interfaces.Holder) Controller {
  return Controller{
    Shared: shared,
    Interfaces: interfaces,
  }
}
