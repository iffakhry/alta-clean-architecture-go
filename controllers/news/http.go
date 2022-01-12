package news

import (
	"alta/cleanarch/businesses/news"
	"alta/cleanarch/controllers"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type NewsController struct {
	newsUseCase news.Usecase
}

func NewNewsController(newsUC news.Usecase) *NewsController {
	return &NewsController{
		newsUseCase: newsUC,
	}
}

func (newsController *NewsController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := newsController.newsUseCase.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, result)
}
