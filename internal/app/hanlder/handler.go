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

	if rest.RestarauntName == "" || rest.CategoryID == 0 {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("missing parameters").Error())
	}

	restaraunts, err := h.service.NewRestaraunt(ctx, rest.RestarauntName, rest.CategoryID)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusCreated, restaraunts)

}
func (h *DipHandler) GetCategories(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	categories, err := h.service.GetCattegories(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(json.Marshal(categories))

	return c.JSON(http.StatusOK, categories)

}
