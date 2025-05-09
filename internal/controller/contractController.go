package controller

import (
	"context"
	model "main/internal/models"
	"main/internal/service"
	errorx "main/internal/utils/myerror"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ContractController struct {
	serviceContract service.IContractService
}

func NewContractController() *ContractController {
	return &ContractController{
		serviceContract: service.NewContractService(),
	}
}

func (cc *ContractController) CreateContract(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var contract *model.Contract
	if err := c.ShouldBindJSON(&contract); err != nil {
		errorx.HandleError(c, errorx.New(http.StatusUnprocessableEntity, "Invalid request data", err))
		return
	}

	err := cc.serviceContract.CreateContract(ctx, contract)
	if errorx.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Contract created successfully"})
}

func (cc *ContractController) UpdateContract(c *gin.Context) {
	var filter model.Filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		errorx.HandleError(c, errorx.New(http.StatusBadRequest, "Invalid query", err))
		return
	}

	var contract *model.Contract
	if err := c.ShouldBindJSON(&contract); err != nil {
		errorx.HandleError(c, errorx.New(http.StatusUnprocessableEntity, "Invalid request data", err))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	err := cc.serviceContract.UpdateContract(ctx, filter, contract)
	if errorx.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contracts updated successfully"})

}

func (cc *ContractController) Delete(c *gin.Context) {
	var filter model.Filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		errorx.HandleError(c, errorx.New(http.StatusBadRequest, "Invalid query parameter", err))
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	err := cc.serviceContract.DeleteContract(ctx, filter)
	if errorx.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contract deleted successfully"})
}

func (cc *ContractController) Search(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var filter model.Filter
	// ShouldBindQuery Bind the query into a struct filter
	if err := c.ShouldBindQuery(&filter); err != nil {
		errorx.HandleError(c, errorx.New(http.StatusBadRequest, "Invalid query parameter", err))
		return
	}

	result, err := cc.serviceContract.Search(ctx, filter)

	if errorx.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, result)
}
