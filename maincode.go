package main

import (
	"cab-management/database"
	"cab-management/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.Connect()

	router.POST("/cab/register", handlers.RegisterCab)
	router.POST("/city/add", handlers.AddCity)
	router.PUT("/cab/change-location", handlers.ChangeCabLocation)
	router.PUT("/cab/change-state", handlers.ChangeCabState)
	router.POST("/cab/book", handlers.BookCab)
	router.POST("/cab/idle-time", handlers.CabIdleTime)
	router.POST("/cab/history", handlers.GetCabHistory)

	router.Run(":8080")
}
