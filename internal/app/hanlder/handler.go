package hanlder

import (
	"DiplomaWork/internal/app/service"
	"github.com/labstack/echo/v4"
)

type DipHandler struct {
	service service.DiplomaService
}

func NewDiplomaHandler(srv service.DiplomaService) *DipHandler {
	return &DipHandler{
		service: srv}

}

func (h *DipHandler) GetAllRestaurants(c echo.Context) error {

	return nil

}
