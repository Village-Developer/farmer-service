package routers

import (
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func RoleRouterGroup(rg *gin.RouterGroup) {
	rg.POST("/roles", roleController.AddRole)
	rg.GET("/roles", roleController.GetAllRoles)
	rg.GET("/roles/:role_id", roleController.GetRoleById)
}
