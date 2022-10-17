package v1

import (
	"github.com/gin-gonic/gin"
)

func InitGroup(group *gin.RouterGroup) {
	group.POST("/wagers", PlaceOneWager)
	group.POST("/wagers/:wager_id", BuyFullOrPartOneWager)
	group.GET("/wagers", ListWagers)
}
