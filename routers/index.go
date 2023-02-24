package routers

import (
	"github.com/gin-gonic/gin"
	"village-developer.com/farmer/controllers"
	"village-developer.com/farmer/middlewares"
)

type IndexRouter struct {
	router *gin.Engine
}

var middleware = new(middlewares.Middleware)
var apiDocumentController = new(controllers.ApiDocumentController)

/* Router Group */
func (h *IndexRouter) SetupRouter() IndexRouter {
	r := IndexRouter{
		router: gin.Default(),
	}
	r.router.Use(middleware.CORSMiddleware())
	r.router.Use(middleware.VerifyToken())
	v1 := r.router.Group("/api/v1")
	ApiDocumentRouterGroup(v1)
	return r
}

/* Run */
func (r IndexRouter) Run(addr string) error {
	return r.router.Run(addr)
}
