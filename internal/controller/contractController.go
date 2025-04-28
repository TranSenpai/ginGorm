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

	if !errorx.BindJSONOrAbort(c, &contract) {
		return
	}

	err := serviceContract.CreateContract(ctx, &contract)
	if errorx.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Contract created successfully"})
}

func (cc ContractController) UpdateContract(c *gin.Context) {
	studentCode := c.Param("studentcode")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	serviceContract := service.GetContractService()
	var contract model.Contract

	if !errorx.BindJSONOrAbort(c, &contract) {
		return
	}

	_, err := serviceContract.Search(ctx, studentCode)
	if errorx.HandleError(c, err) {
		return
	}

	err = serviceContract.UpdateContract(ctx, studentCode, &contract)
	if errorx.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contract updated successfully"})
}

func (cc ContractController) Delete(c *gin.Context) {
	studentCode := c.Param("studentcode")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	serviceContract := service.GetContractService()

	_, err := serviceContract.Search(ctx, studentCode)
	if errorx.HandleError(c, err) {
		return
	}

	err = serviceContract.DeleteContract(ctx, studentCode)
	if errorx.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contract deleted successfully"})
}

func (cc ContractController) Search(c *gin.Context) {
	studentCode := c.Param("studentcode")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	serviceContract := service.GetContractService()

	contract, err := serviceContract.Search(ctx, studentCode)
	if errorx.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contract})
}

func (cc ContractController) SearchAll(c *gin.Context) {
	serviceContract := service.GetContractService()

	contracts, err := serviceContract.SearchAll()
	if errorx.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contracts})
}

func (cc ContractController) SearchByName(c *gin.Context) {
	fullName := c.Param("fullname")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	serviceContract := service.GetContractService()

	contract, err := serviceContract.SearchByName(ctx, fullName)
	if errorx.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contract})
}
