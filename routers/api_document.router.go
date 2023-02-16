package routers

import (
	"github.com/gin-gonic/gin"
)

type ApiDocumentRouter struct{}

func ApiDocumentRouterGroup(rg *gin.RouterGroup) {
	rg.GET("/ping", apiDocumentController.Ping)
	rg.GET("/pong", apiDocumentController.Ping)
	rg.GET("/pang", apiDocumentController.Ping)
}
