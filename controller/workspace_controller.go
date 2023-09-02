package controller

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

	"shiftlab-go-rest-api/model"
	"shiftlab-go-rest-api/usecase"
)

type IWorkspaceController interface {
	GetAllWorkspaces(c echo.Context) error
	GetWorkspaceByWorkspaceId(c echo.Context) error
	CreateWorkspace(c echo.Context) error
	UpdateWorkspace(c echo.Context) error
	DeleteWorkspace(c echo.Context) error
}

type workspaceController struct {
	wu usecase.IWorkspaceUsecase
}

func NewWorkspaceController(wu usecase.IWorkspaceUsecase) IWorkspaceController {
	return &workspaceController{wu}
}

func (wc *workspaceController) GetAllWorkspaces(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	workspaceRes, err := wc.wu.GetAllWorkspaces(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, workspaceRes)
}

func (wc *workspaceController) GetWorkspaceByWorkspaceId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("workspaceId")
	workspaceId, _ := strconv.Atoi(id)
	workspaceRes, err := wc.wu.GetWorkspaceById(uint(userId.(float64)), uint(workspaceId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, workspaceRes)
}

func (wc *workspaceController) CreateWorkspace(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	workspace := model.Workspace{}
	if err := c.Bind(&workspace); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	workspace.UserId = uint(userId.(float64))
	workspaceRes, err := wc.wu.CreateWorkspace(workspace)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, workspaceRes)
}

func (wc *workspaceController) UpdateWorkspace(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("workspaceId")
	workspaceId, _ := strconv.Atoi(id)

	workspace := model.Workspace{}
	if err := c.Bind(&workspace); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	workspaceRes, err := wc.wu.UpdateWorkspace(workspace, uint(userId.(float64)), uint(workspaceId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, workspaceRes)
}

func (wc *workspaceController) DeleteWorkspace(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("workspaceId")
	workspaceId, _ := strconv.Atoi(id)

	err := wc.wu.DeleteWorkspace(uint(userId.(float64)), uint(workspaceId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
