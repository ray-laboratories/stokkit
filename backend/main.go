package main

import (
	"github.com/gin-gonic/gin"
	stockApp "stokkit/application/stock"
	"stokkit/infra/memcache"
	stockRoutes "stokkit/routes/stock"
)

func main() {
	router := gin.Default()
	memCache := memcache.NewGarmentMemCache()
	gSvc := stockApp.NewService(memCache)
	group := router.Group("/")
	stockRoutes.SetupStockRoutes(group, gSvc)
	router.Run(":8080")
}
