package controller

import (
	"github.com/topazur/leesin/internal/service"
	"github.com/topazur/leesin/pkg/log"
)

type Controller struct {
	logger *log.Logger
	store  service.Store
}

func NewController(logger *log.Logger, store service.Store) (controller *Controller) {
	return &Controller{
		logger: logger,
		store:  store,
	}
}
