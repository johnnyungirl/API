package route

import (
	"API/controller"
	"API/repository"
	config "API/util"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()
	custRepo := repository.NewCustomerRepository(db)
	custHandler := controller.NewCustomerHandler(custRepo)
	apiRoutes := httpRouter.Group("/customer")
	apiRoutes.GET("/partnerlist", custHandler.ListParter)
	httpRouter.Run()
	config.LoadConfig()

}
