package controllers

import (
	"goweather/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherController struct {
	WeatherService *services.WeatherService
}

// GetWeather handles the request to fetch weather information for a given city
func (c *WeatherController) GetWeather(ctx *gin.Context) {
	city := ctx.Query("city")
	if city == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "City is required"})
		return
	}

	weather, err := c.WeatherService.GetWeather(city)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"weather": weather})
}
