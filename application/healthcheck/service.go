package healthcheck

import (
	"github.com/gofiber/fiber/v2"
	"tracking-server/shared"
	"tracking-server/shared/dto"
)

type (
	Service interface {
		HttpHealthcheck(app *fiber.App) dto.Status
	}

	service struct {
		shared shared.Holder
	}
)

func (h *service) HttpHealthcheck(app *fiber.App) dto.Status {
	data := dto.HCData{
		HandlerCount: app.HandlersCount(),
	}
	return dto.Status{
		Name:   dto.HTTP,
		Status: dto.OK,
		Data:   data,
	}
}

func NewHealthcheckService(shared shared.Holder) Service {
	return &service{
		shared: shared,
	}
}
