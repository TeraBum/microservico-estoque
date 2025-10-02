package stockmoves

import (
	stockmovesSrvc "api-estoque/internal/service/stock_moves"

	"github.com/sirupsen/logrus"
)

type Controller struct {
	Service *stockmovesSrvc.Service
	Logger  *logrus.Logger
}

func New(service *stockmovesSrvc.Service, logger *logrus.Logger) *Controller {
	return &Controller{
		Service: service,
		Logger:  logger,
	}
}
