package handler

import (
	"net/http"

	"github.com/PongponZ/go-book-store/internal/delivery/http/handler/response"
	"github.com/PongponZ/go-book-store/internal/domain/usecase"
	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	Create(c echo.Context) error
	FindByID(c echo.Context) error
	Update(c echo.Context) error
	DeleteByID(c echo.Context) error
}

type userHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserHandler(userUsecase usecase.IUserUsecase) IUserHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (h *userHandler) Create(c echo.Context) error {
	return nil
}

func (h *userHandler) FindByID(c echo.Context) error {
	userID := c.Param("userID")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, "user id is required")
	}

	user, err := h.userUsecase.FindByID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response.FindByID{
		Status: response.StatusOk,
		Data:   *user,
	})
}

func (h *userHandler) Update(c echo.Context) error {
	return nil
}

func (h *userHandler) DeleteByID(c echo.Context) error {
	return nil
}
