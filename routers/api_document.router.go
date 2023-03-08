package routers

import (
	"github.com/gin-gonic/gin"
)

type ApiDocumentRouter struct{}

func ApiDocumentRouterGroup(rg *gin.RouterGroup) {
	rg.POST("/api_documents/project", apiDocumentController.AddApiDocumentProject)
	rg.GET("/api_documents", apiDocumentController.GetAllApiDocuments)
	rg.DELETE("/api_documents/project/:api_project_id", apiDocumentController.DeleteApiDocumentProject)
	rg.POST("/api_documents//project/:api_project_id/group", apiDocumentController.AddApiDocumentGroup)
	rg.POST("/api_documents/project/:api_project_id/api", apiDocumentController.AddApiDocuments)
	rg.DELETE("/api_documents/group/:api_group_id", apiDocumentController.DeleteApiDocumentGroup)
	rg.DELETE("/api_documents/:api_document_id", apiDocumentController.DeleteApiDocument)
}
