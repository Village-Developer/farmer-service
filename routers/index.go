package routers

import (
	"github.com/gin-gonic/gin"
	"village-developer.com/farmer/controllers"
)

type IndexRouter struct {
	router *gin.Engine
}

var apiDocumentController = new(controllers.ApiDocumentController)
var roleController = new(controllers.RoleController)

func (h *IndexRouter) SetupRouter() IndexRouter {
	r := IndexRouter{
		router: gin.Default(),
	}
	v1 := r.router.Group("/api/v1")
	ApiDocumentRouterGroup(v1)
	RoleRouterGroup(v1)
	return r
}

func (r IndexRouter) Run(addr string) error {
	return r.router.Run(addr)
}
