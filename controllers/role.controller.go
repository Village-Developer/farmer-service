package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"village-developer.com/farmer/configs"
	"village-developer.com/farmer/models"
	"village-developer.com/farmer/services"
)

type RoleController struct{}

func (h *RoleController) AddRole(c *gin.Context) {
	role := models.Role{RoleId: 8, RoleName: "admin"}
	response := services.RoleService{}.AddRoleService(role)
	db := new(configs.Configs)
	if response != "Success" {
		c.JSON(http.StatusBadRequest, db.ResponseFailed(response))
	} else {
		c.JSON(http.StatusOK, db.ResponseSuccess(response))
	}
}

func (h *RoleController) GetAllRoles(c *gin.Context) {
	roles := []models.Role{}
	response := services.RoleService{}.GetAllRolesService(roles)
	c.JSON(http.StatusOK, response)
}

func (h *RoleController) GetRoleById(c *gin.Context) {
	role := models.Role{}
	response, status := services.RoleService{}.GetRoleByIdService(role, c.Param("role_id"))
	c.JSON(status, response)
}
