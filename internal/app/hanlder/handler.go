package hanlder

import (
	"DiplomaWork/internal/app/model"
	"DiplomaWork/internal/app/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"

	"net/http"
)

type DipHandler struct {
	service service.DiplomaService
}

func NewDiplomaHandler(srv service.DiplomaService) *DipHandler {
	return &DipHandler{
		service: srv}

}

func (h *DipHandler) GetAllRestaraunts(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	restaraunts, err := h.service.GetAllRestaraunt(ctx)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, restaraunts)

	return nil

}
func (h *DipHandler) NewRestaraunts(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	body := c.Request().Body
	req, _ := io.ReadAll(body)
	defer body.Close()

	rest := &model.RestarauntsModel{}
	err := json.Unmarshal(req, rest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	restaraunts, err := h.service.NewRestaraunt(ctx, rest.RestarauntName, rest.RestarauntCategory)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusCreated, restaraunts)

}
