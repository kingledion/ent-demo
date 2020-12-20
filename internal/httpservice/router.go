package httpservice

import (
	"fmt"

	"github.com/kingledion/ent-demo/internal/config"

	"github.com/gin-gonic/gin"
)

// Run launches a Gin webserver to listen and serve on the config port
func Run(h Handler, cfg config.HttpConfig) {

	router := initRoutes(h)

	router.Run(fmt.Sprintf(":%s", cfg.Port))

}

// InitRoutes creates all of the routes for our application and returns a router
func initRoutes(h Handler) *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	v1.GET("/:userid/ordered-at", h.GetOrderedAtByUser)
	v1.POST("/order", h.AddOrder)

	return router

}
