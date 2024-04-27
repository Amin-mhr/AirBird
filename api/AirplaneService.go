package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Airplane struct {
	name         string
	seatCount    int
	airPlaneType int
}

type AirplaneServiceInterface interface {
	createAirplane(*Airplane) error
	updateairplane(*Airplane) error
}

type AirplaneService struct{}

func (*AirplaneService) createAirplane(airplane *Airplane) error {
	return nil
}

func (s *AirplaneService) updateairplane(airplane *Airplane) error {
	return nil
}

func NewAirplaneService() AirplaneServiceInterface {
	return &AirplaneService{}
}

func createAirplane(service AirplaneServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		airPlane := new(Airplane)
		if err := c.Bind(airPlane); err != nil {
			return err
		}
		if err := service.createAirplane(airPlane); err != nil {
			fmt.Println(err)
		}
		return c.JSON(http.StatusOK, airPlane)
	}
}

func AirPlaneRoutes(server *echo.Echo) {
	airplaneService := NewAirplaneService()
	server.POST("/AirPlane", createAirplane(airplaneService))
}
