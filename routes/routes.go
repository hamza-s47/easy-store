package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hamza-s47/easy-store/handlers"
)

func Routes_init() *gin.Engine {
	route := gin.Default()
	api := route.Group("/api")
	{
		api.GET("/product", handlers.ProdHandler)
	}

	return route
}
