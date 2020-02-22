package category

import (
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbcategory"
	"github.com/gkeele21/ldsmediaAPI/pkg/log"
	"github.com/labstack/echo"
	"net/http"
)

// RegisterRoutes sets up routes on a given nova.Server instance
func RegisterRoutes(g *echo.Group) {
	g.GET("/categories", getCategories)
}

// Response is the json representation of a user
//type Response struct {
//	User dbuser.User
//}

// getCategories retrieves all the categories
func getCategories(req echo.Context) error {
	var err error

	log.LogRequestData(req)
	categories, err := dbcategory.ReadAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "couldn't get categories", err)
	}

	return req.JSON(http.StatusOK, categories)
}

