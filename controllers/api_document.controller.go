package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"village-developer.com/farmer/configs"
	"village-developer.com/farmer/services"
)

type ApiDocumentController struct{}

func (ApiDocumentController) AddApiDocumentProject(c *gin.Context) {
	response, status := services.ApiDocumentService{}.AddApiDocumentProjectService(c)
	if status != 200 {
		c.JSON(http.StatusBadRequest, new(configs.Error).ResponseFailed(response))
	} else {
		c.JSON(http.StatusOK, new(configs.Success).ResponseSuccess(response))
	}
}

func (ApiDocumentController) GetAllApiDocuments(c *gin.Context) {
	response, status := services.ApiDocumentService{}.GetAllApiDocumentService(c)
	if status == 200 {
		c.JSON(status, new(configs.SuccessData).ResponseSuccess("Success", response))
	} else {
		c.JSON(status, new(configs.Error).ResponseFailed("No data found"))
	}
}

func (ApiDocumentController) DeleteApiDocumentProject(c *gin.Context) {
	response, status := services.ApiDocumentService{}.DeleteApiDocumentProjectService(c)
	if status != 200 {
		c.JSON(http.StatusBadRequest, new(configs.Error).ResponseFailed(response))
	} else {
		c.JSON(http.StatusOK, new(configs.Success).ResponseSuccess(response))
	}
}

func (ApiDocumentController) AddApiDocumentGroup(c *gin.Context) {
	response, status := services.ApiDocumentService{}.AddApiDocumentGroupService(c)
	if status != 200 {
		c.JSON(http.StatusBadRequest, new(configs.Error).ResponseFailed(response))
	} else {
		c.JSON(http.StatusOK, new(configs.Success).ResponseSuccess(response))
	}
}

func (ApiDocumentController) AddApiDocuments(c *gin.Context) {
	response, status := services.ApiDocumentService{}.AddApiDocumentsService(c)
	if status != 200 {
		c.JSON(http.StatusBadRequest, new(configs.Error).ResponseFailed(response))
	} else {
		c.JSON(http.StatusOK, new(configs.Success).ResponseSuccess(response))
	}
}

func (ApiDocumentController) DeleteApiDocumentGroup(c *gin.Context) {
	response, status := services.ApiDocumentService{}.DeleteApiDocumentGroupService(c)
	if status != 200 {
		c.JSON(http.StatusBadRequest, new(configs.Error).ResponseFailed(response))
	} else {
		c.JSON(http.StatusOK, new(configs.Success).ResponseSuccess(response))
	}
}

func (ApiDocumentController) DeleteApiDocument(c *gin.Context) {
	response, status := services.ApiDocumentService{}.DeleteApiDocumentService(c)
	if status != 200 {
		c.JSON(http.StatusBadRequest, new(configs.Error).ResponseFailed(response))
	} else {
		c.JSON(http.StatusOK, new(configs.Success).ResponseSuccess(response))
	}
}

func (ApiDocumentController) GetApiDocument(c *gin.Context) {
	response, status := services.ApiDocumentService{}.GetApiDocumentService(c)
	if status == 200 {
		c.JSON(status, new(configs.SuccessData).ResponseSuccess("Success", response[0]))
	} else {
		c.JSON(status, new(configs.Error).ResponseFailed("No data found"))
	}
}
