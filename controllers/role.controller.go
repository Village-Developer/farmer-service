package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"village-developer.com/farmer/configs"
	"village-developer.com/farmer/models"
	"village-developer.com/farmer/services"
)

type RoleController struct{}

func (r *RoleController) AddRole(c *gin.Context) {
	role := models.Role{RoleId: 8, RoleName: "admin"}
	response := services.RoleService{}.AddRoleService(role)
	if response != "Success" {
		c.JSON(http.StatusBadRequest, new(configs.Error).ResponseFailed(response))
	} else {
		c.JSON(http.StatusOK, new(configs.Success).ResponseSuccess(response))
	}
}

func (r *RoleController) GetAllRoles(c *gin.Context) {
	roles := []models.Role{}
	response, status := services.RoleService{}.GetAllRolesService(roles)
	if status == 200 {
		c.JSON(status, new(configs.SuccessData).ResponseSuccess("Success", response))
	} else {
		c.JSON(status, new(configs.Error).ResponseFailed("No data found"))
	}
}

func (r *RoleController) GetRoleById(c *gin.Context) {
	role := models.Role{}
	response, status := services.RoleService{}.GetRoleByIdService(role, c.Param("role_id"))
	if status == 200 {
		c.JSON(status, new(configs.SuccessData).ResponseSuccess("Success", response))
	} else {
		c.JSON(status, new(configs.Error).ResponseFailed("No data found"))
	}
}
