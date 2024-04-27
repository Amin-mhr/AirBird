package api

import "github.com/labstack/echo/v4"

func SetRouting(server *echo.Echo) {
	AirPlaneRoutes(server)
}
