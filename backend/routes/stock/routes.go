package stock

import (
	"github.com/gin-gonic/gin"
	"net/http"
	stockApp "stokkit/application/stock"
	stockTypes "stokkit/domain/stock"
	"stokkit/routes"
	"strconv"
)

func SetupStockRoutes(router *gin.RouterGroup, svc *stockApp.Service) {
	router.GET("/stock", func(c *gin.Context) {
		garments, err := svc.GetAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, routes.NewResponseFromError(err))
			return
		}
		c.JSON(http.StatusOK, routes.NewResponse(garments))
		return
	})

	router.GET("/stock/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.NewResponseFromError(err))
			return
		}
		garment, err := svc.Get(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.NewResponseFromError(err))
			return
		}
		c.JSON(http.StatusOK, routes.NewResponse(garment))
		return
	})

	router.POST("/stock", func(c *gin.Context) {
		var garmentVal stockTypes.GarmentValues
		if err := c.BindJSON(&garmentVal); err != nil {
			c.JSON(http.StatusBadRequest, routes.NewResponseFromError(err))
			return
		}
		var garment stockTypes.Garment
		garment.GarmentValues = garmentVal
		garment, err := svc.Save(c.Request.Context(), garment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, routes.NewResponseFromError(err))
			return
		}
		c.JSON(http.StatusOK, routes.NewResponse(garment))
		return
	})

	router.POST("/stock/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, routes.NewResponseFromError(err))
			return
		}
		var garmentVal stockTypes.GarmentValues
		if err := c.ShouldBindJSON(&garmentVal); err != nil {
			c.JSON(http.StatusBadRequest, routes.NewResponseFromError(err))
			return
		}
		var garment stockTypes.Garment
		garment.GarmentValues = garmentVal
		garment.ID = id
		garment, err = svc.Save(c.Request.Context(), garment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, routes.NewResponseFromError(err))
			return
		}
		c.JSON(http.StatusOK, routes.NewResponse(garment))
		return
	})
}
