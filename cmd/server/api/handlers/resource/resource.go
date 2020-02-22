package resource

import (
	"fmt"
	"github.com/gkeele21/ldsmediaAPI/internal/app/database/dbresource"
	"github.com/gkeele21/ldsmediaAPI/pkg/log"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// RegisterRoutes sets up routes on a given nova.Server instance
func RegisterRoutes(g *echo.Group) {
	g.GET("/resources/:userId", getUserResources)
}

// Response is the json representation of a user
//type Response struct {
//	User dbuser.User
//}

// getUserResources retrieves all the resources for the user with the route parameter :userId
func getUserResources(req echo.Context) error {
	var err error

	log.LogRequestData(req)

	resources, err := dbresource.ReadAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad user ID given")
	}

	fmt.Printf("Resources : %#v\n", resources)


	searchID := req.Param("userId")
	num, err := strconv.ParseInt(searchID, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad user ID given")
	}

	userResources, err := dbresource.ReadByUserID(num)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "couldn't get resources", err)
	}

	fmt.Printf("# of UserResources : %s", len(userResources))

	return req.JSON(http.StatusOK, userResources)
}

