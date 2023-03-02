package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"merchant/internal/handler"
	"merchant/internal/repository"
	"merchant/internal/service"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	prodRepo := repository.NewProductRepository(db)
	prodService := service.NewProductService(prodRepo)
	prodHandler := handler.NewProductHandler(prodService)

	product := router.Group("merchant")
	{
		product.POST("/editStore", prodHandler.EditStore)
	}

	return router
}
