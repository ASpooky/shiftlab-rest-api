package controller

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

	"shiftlab-go-rest-api/model"
	"shiftlab-go-rest-api/usecase"
)

type IShiftController interface {
	GetAllShifts(c echo.Context) error
	GetShiftByWorkspaceId(c echo.Context) error
	CreateShift(c echo.Context) error
	UpdateShift(c echo.Context) error
	DeleteShift(c echo.Context) error
}

type shiftController struct {
	su usecase.IShiftUsecase
}

func NewShiftController(su usecase.IShiftUsecase) IShiftController {
	return &shiftController{su}
}

func (sc *shiftController) GetAllShifts(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	shiftsRes, err := sc.su.GetAllShifts(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, shiftsRes)
}

func (sc *shiftController) GetShiftByWorkspaceId(c echo.Context) error {
	id := c.Param("workspaceId")
	workspaceId, _ := strconv.Atoi(id)
	shiftRes, err := sc.su.GetShiftByWorkspaceId(uint(workspaceId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, shiftRes)
}

func (sc *shiftController) CreateShift(c echo.Context) error {
	shift := model.Shift{}
	if err := c.Bind(&shift); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	shiftRes, err := sc.su.CreateShift(shift)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, shiftRes)
}

func (sc *shiftController) UpdateShift(c echo.Context) error {
	id := c.Param("shiftId")
	shiftId, _ := strconv.Atoi(id)

	shift := model.Shift{}
	if err := c.Bind(&shift); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	shiftRes, err := sc.su.UpdateShift(shift, uint(shiftId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, shiftRes)
}

func (sc *shiftController) DeleteShift(c echo.Context) error {

	id := c.Param("shiftId")
	shiftId, _ := strconv.Atoi(id)

	err := sc.su.DeleteShift(uint(shiftId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
