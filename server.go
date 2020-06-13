package main

import (
	"github.com/Masato4556/youtubemanagergo/middlewares"
	"github.com/Masato4556/youtubemanagergo/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main(){
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.YoutubeService())

	// routes
	routes.Init(e)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}