package main

import (
	"goweather/config"
	"goweather/controllers"
	"goweather/routes"
	"goweather/services"
)

func main() {
	// Load configuration from .env file
	config.LoadConfig()

	// Initialize services and controllers
	weatherService := &services.WeatherService{}
	weatherController := &controllers.WeatherController{WeatherService: weatherService}

	// Set up routes and start the server
	router := routes.SetupRouter(weatherController)
	router.Run(":8080")
}
