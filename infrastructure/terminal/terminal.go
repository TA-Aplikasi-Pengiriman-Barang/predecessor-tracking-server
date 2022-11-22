package terminal

import (
	"tracking-server/interfaces"
	"tracking-server/shared"
	"tracking-server/shared/common"
	"tracking-server/shared/dto"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces interfaces.Holder
	Shared     shared.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	terminal := app.Group("/terminal")
	terminal.Get("/:id", c.get)
	terminal.Post("/allTerminal", c.allTerminal)
	terminal.Post("/twoClosest", c.twoClosestTerminal)
}

// All godoc
// @Tags Terminal
// @Summary Get terminal info
// @Description Put all mandatory parameter
// @Param id path string true "terminal ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.GetTerminalInfoResponse
// @Failure 200 {object} dto.GetTerminalInfoResponse
// @Router /terminal/{id} [get]
func (c *Controller) get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	c.Shared.Logger.Infof("get terminal info, data: %s", id)

	res, err := c.Interfaces.TerminalViewsService.GetTerminalInfo(id)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Terminal
// @Summary Get all terminal sorted by distance
// @Description Put all mandatory parameter
// @Param GetAllTerminalDto body dto.GetAllTerminalDto true "GetAllTerminalDto"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.GetAllTerminalResponse
// @Failure 200 {object} dto.GetAllTerminalResponse
// @Router /terminal/allTerminal [post]
func (c *Controller) allTerminal(ctx *fiber.Ctx) error {
	var (
		body     dto.GetAllTerminalDto
		response dto.GetAllTerminalResponse
	)

	err := common.DoCommonRequest(ctx, &body)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("get all terminal, data: %s", body)

	response, err = c.Interfaces.TerminalViewsService.GetAllTerminalSorted(body)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, response)
}

// All godoc
// @Tags Terminal
// @Summary Get two closestterminal
// @Description Put all mandatory parameter
// @Param GetAllTerminalDto body dto.GetAllTerminalDto true "GetAllTerminalDto"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.GetAllTerminalResponse
// @Failure 200 {object} dto.GetAllTerminalResponse
// @Router /terminal/twoClosest [post]
func (c *Controller) twoClosestTerminal(ctx *fiber.Ctx) error {
	var (
		body     dto.GetAllTerminalDto
		response dto.GetAllTerminalResponse
	)

	err := common.DoCommonRequest(ctx, &body)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("get all terminal, data: %s", body)

	response, err = c.Interfaces.TerminalViewsService.GetTwoClosesTerminal(body)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, response)
}

func NewController(interfaces interfaces.Holder, shared shared.Holder) Controller {
	return Controller{
		Interfaces: interfaces,
		Shared:     shared,
	}
}
