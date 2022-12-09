package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"messenger/internal"
)

func main() {
	// Echo instance
	e := echo.New()
	k := internal.Messenger{}
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", k.Hello)
	e.GET("/domoi", k.Dima)
	e.POST("/msg", k.Message)
	e.GET("/msg", k.History)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
