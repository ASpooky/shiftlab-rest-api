package router

import (
	"net/http"
	"os"
	"shiftlab-go-rest-api/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, wc controller.IWorkspaceController, sc controller.IShiftController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		//CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge: 60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	w := e.Group("/workspaces")
	w.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	w.GET("", wc.GetAllWorkspaces)
	w.GET("/:workspaceId", wc.GetWorkspaceByWorkspaceId)
	w.POST("", wc.CreateWorkspace)
	w.PUT("/:workspaceId", wc.UpdateWorkspace)
	w.DELETE("/:workspaceId", wc.DeleteWorkspace)

	s := e.Group("/shifts")
	s.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	s.GET("", sc.GetAllShifts)
	s.GET("/:workspaceId", sc.GetShiftByWorkspaceId)
	s.POST("", sc.CreateShift)
	s.PUT("/:shiftId", sc.UpdateShift)
	s.DELETE("/:shiftId", sc.DeleteShift)

	return e
}
