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

type ContractController struct{}

func (cc ContractController) CreateContract(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	serviceContract := service.GetContractService()
	var contract model.Contract

	if err := c.ShouldBindJSON(&contract); err != nil {
		errorx.HandleError(c, errorx.New(errorx.StatusUnprocessableEntity, "Invalid request data", err))
		return
	}

	err := serviceContract.CreateContract(ctx, &contract)
	errorx.HandleError(c, err)

	c.JSON(http.StatusCreated, gin.H{"message": "Contract created successfully"})
}

func (cc ContractController) UpdateContract(c *gin.Context) {
	studentCode := c.Param("studentcode")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	serviceContract := service.GetContractService()
	var contract model.Contract

	if err := c.ShouldBindJSON(&contract); err != nil {
		errorx.HandleError(c, errorx.New(errorx.StatusUnprocessableEntity, "Invalid request data", err))
		return
	}

	_, err := serviceContract.Search(ctx, studentCode)
	errorx.HandleError(c, err)

	err = serviceContract.UpdateContract(ctx, studentCode, &contract)
	errorx.HandleError(c, err)

	c.JSON(http.StatusOK, gin.H{"message": "Contract updated successfully"})
}

func (cc ContractController) Delete(c *gin.Context) {
	studentCode := c.Param("studentcode")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	serviceContract := service.GetContractService()

	_, err := serviceContract.Search(ctx, studentCode)
	errorx.HandleError(c, err)

	err = serviceContract.DeleteContract(ctx, studentCode)
	errorx.HandleError(c, err)

	c.JSON(http.StatusOK, gin.H{"message": "Contract deleted successfully"})
}

func (cc ContractController) Search(c *gin.Context) {
	studentCode := c.Param("studentcode")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	serviceContract := service.GetContractService()

	contract, err := serviceContract.Search(ctx, studentCode)
	errorx.HandleError(c, err)

	c.JSON(http.StatusOK, gin.H{"data": contract})
}

func (cc ContractController) SearchAll(c *gin.Context) {
	serviceContract := service.GetContractService()

	contracts, err := serviceContract.SearchAll()
	errorx.HandleError(c, err)

	c.JSON(http.StatusOK, gin.H{"data": contracts})
}

func (cc ContractController) SearchByName(c *gin.Context) {
	fullName := c.Param("fullname")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	serviceContract := service.GetContractService()

	contract, err := serviceContract.SearchByName(ctx, fullName)
	errorx.HandleError(c, err)

	c.JSON(http.StatusOK, gin.H{"data": contract})
}
