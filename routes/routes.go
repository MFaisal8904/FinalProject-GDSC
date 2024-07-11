package routes

import (
	"goweather/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the routes for the application
func SetupRouter(weatherController *controllers.WeatherController) *gin.Engine {
	router := gin.Default()

	weatherRoutes := router.Group("/weather")
	{
		weatherRoutes.GET("/current", weatherController.GetWeather)
	}

	return router
}
