package controller

import (
	model "main/internal/models"
	"main/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContractController struct{}

func (cc ContractController) CreateContract(c *gin.Context) {
	serviceContract := service.GetContractService()
	var contract model.Contract

	err := c.ShouldBindJSON(&contract)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Some field are not right"})
		return
	}

	_, err = serviceContract.Search(contract.ID)
	if err == nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "This contract has already exist"})
		return
	}

	err = serviceContract.CreateContract(&contract)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Can not create contract"})
		return
	}

	data, err := serviceContract.Search(contract.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve created contract"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":          "Your created contract",
		"updated contract": data,
	})
}

func (cc ContractController) UpdateContract(c *gin.Context) {
	var contract model.Contract
	serviceContract := service.GetContractService()

	// Ép data từ JSON từ request HTTP sang type struct
	err := c.ShouldBindJSON(&contract)
	if err != nil {
		// Có một lỗi cú pháp trong yêu cầu và yêu cầu bị từ chối.
		c.JSON(http.StatusBadRequest, gin.H{"message": "Some field are not right"})
		return
	}

	_, err = serviceContract.Search(contract.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "This contract does not exist"})
		return
	}

	err = serviceContract.UpdateContract(contract.ID, &contract)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not update contract"})
		return
	}

	data, err := serviceContract.Search(contract.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not retrieve updated contract"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "Your updated contract",
		"updated contract": data,
	})
}

func (cc ContractController) Delete(c *gin.Context) {
	id := c.Param("id")
	serviceContract := service.GetContractService()

	role, err := serviceContract.Search(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	err = serviceContract.DeleteContract(role.ID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "role does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete successfully"})
}

func (cc ContractController) Search(c *gin.Context) {
	id := c.Param("id")
	var contract model.Contract
	serviceContract := service.GetContractService()

	contract, err := serviceContract.Search(id)

	// gin.H is a type alias for map[string]interface{}.
	// It's a shorthand notation used to create maps that store data
	// with string keys and can hold values of any type.
	// This type allows you to easily define and use dynamic
	// JSON-like structures within your Gin applications.

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "contract found",
		"data":    contract,
	})
}

func (cc ContractController) SearchAll(c *gin.Context) {
	serviceContract := service.GetContractService()
	contracts, err := serviceContract.SearchAll()

	if err.Error() == "dont have any contracts in table" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empty contract table"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "contracts found",
		"data":    contracts,
	})
}
