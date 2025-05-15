package controller

import (
	"context"
	"main/internal/models"
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

// Why use pointer in gin.Context?
//
//	every request will create a context that store
// 	the Request data Response writer. Param, Header, Body, ... about it
//	use pointer to make sure the gin.context variable will save
// 	the modification when pass around handler
// 	it also make sure don't copy so much data when pass by value

func (cc *ContractController) CreateContract(ginContext *gin.Context) {
	ctx, cancel := context.WithTimeout(ginContext.Request.Context(), 5*time.Second)
	defer cancel()

	var contract models.Contract
	if err := ginContext.ShouldBindJSON(&contract); err != nil {
		ginContext.Error(errorx.NewMyError(http.StatusUnprocessableEntity, "Invalid request data", err, time.Now()))
		return
	}

	if err := cc.serviceContract.CreateContract(ctx, &contract); err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{
		"message": "Create contract successfully",
		"data":    contract,
		"time":    time.Now(),
	})
}

func (cc *ContractController) UpdateContract(ginContext *gin.Context) {
	var filter models.Filter
	if err := ginContext.ShouldBindQuery(&filter); err != nil {
		ginContext.Error(errorx.NewMyError(http.StatusBadRequest, "Invalid query", err, time.Now()))
		return
	}

	var contract models.Contract
	if err := ginContext.ShouldBindJSON(&contract); err != nil {
		ginContext.Error(errorx.NewMyError(http.StatusUnprocessableEntity, "Invalid request data", err, time.Now()))
		return
	}

	ctx, cancel := context.WithTimeout(ginContext.Request.Context(), 5*time.Second)
	defer cancel()

	if err := cc.serviceContract.UpdateContract(ctx, filter, &contract); err != nil {
		// Error attaches an error to the current context.
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{
		"message": "Update contract successfully",
		"data":    contract,
		"time":    time.Now(),
	})
}

func (cc *ContractController) Delete(ginContext *gin.Context) {
	var filter models.Filter
	if err := ginContext.ShouldBindQuery(&filter); err != nil {
		ginContext.Error(errorx.NewMyError(http.StatusBadRequest, "Invalid query parameter", err, time.Now()))
		return
	}

	ctx, cancel := context.WithTimeout(ginContext.Request.Context(), 5*time.Second)
	defer cancel()

	if err := cc.serviceContract.DeleteContract(ctx, filter); err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{
		"message": "Delete contract successfully",
	})
}

func (cc *ContractController) Search(ginContext *gin.Context) {
	ctx, cancel := context.WithTimeout(ginContext.Request.Context(), 5*time.Second)
	defer cancel()

	var filter models.Filter
	// ShouldBindQuery Bind the query into a struct filter
	if err := ginContext.ShouldBindQuery(&filter); err != nil {
		// Attach an error to the current context
		ginContext.Error(errorx.NewMyError(http.StatusBadRequest, "Invalid query parameter", err, time.Now()))
		return
	}

	result, err := cc.serviceContract.Search(ctx, filter)
	if err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{
		"result": result,
		"total":  len(result),
	})
}

func (cc *ContractController) SearchContractInRoom(ginContext *gin.Context) {
	ctx, cancel := context.WithTimeout(ginContext.Request.Context(), 5*time.Second)
	defer cancel()

	result, err := cc.serviceContract.GetTotalContractEachRoom(ctx)
	if err != nil {
		ginContext.Error(err)
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
