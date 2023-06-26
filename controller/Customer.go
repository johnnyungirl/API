package controller

import (
	"API/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	custHandler repository.CustomerRepository
}
type CustomerHandler interface {
	ListParter(*gin.Context)
	ListMenu(*gin.Context)
	GetPartner(*gin.Context)
	GetMenu(*gin.Context)
}

func NewCustomerHandler(custHandler repository.CustomerRepository)(*customerHandler){
	return &customerHandler{
		custHandler: custHandler,
	}
}
func (handle customerHandler) ListParter(ctx *gin.Context) {
	partner, err := handle.custHandler.ListPartner()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"code":    -1,
			"message": "something wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "succes",
		"partner": partner,
	})

}
