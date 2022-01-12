package routes

import (
	"alta/cleanarch/controllers/news"

	echo "github.com/labstack/echo/v4"
)

type ControllerList struct {
	NewsController news.NewsController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	news := e.Group("news")
	news.GET("/list", cl.NewsController.GetAll)
	// news.POST("/store", cl.NewsController.Store, middlewareApp.RoleValidation("admin"))
	// news.PUT("/update", cl.NewsController.Update)
}
