package main

import (
	"goweather/config"
	"goweather/controllers"
	"goweather/routes"
	"goweather/services"
)

func main() {
	config.LoadConfig()

	weatherService := &services.WeatherService{}
	weatherController := &controllers.WeatherController{WeatherService: weatherService}

	router := routes.SetupRouter(weatherController)
	router.Run(":8080")
}
