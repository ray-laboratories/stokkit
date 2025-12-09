package main

import (
	stockApp "stokkit/application/stock"
	"stokkit/infra/memcache"
	stockRoutes "stokkit/routes/stock"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	memCache := memcache.NewGarmentMemCache()
	gSvc := stockApp.NewService(memCache)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	group := router.Group("/")
	stockRoutes.SetupStockRoutes(group, gSvc)

	router.Run(":8080")
}
