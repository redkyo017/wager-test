package server

import (
	"handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoute() *gin.Engine {
	r := gin.Default()

	// WAGER GROUP ROUTES
	wagerRoutes := r.Group("/wagers")
	wagerRoutes.GET("", handlers.GetWagers)
	wagerRoutes.POST("", handlers.CreateWager)

	r.POST("/buy/:wager_id", handlers.BuyWager)

	return r
}
