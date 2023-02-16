package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"village-developer.com/farmer/configs"
	"village-developer.com/farmer/models"
)

type ApiDocumentController struct{}

func (h *ApiDocumentController) Ping(c *gin.Context) {
	db := new(configs.Configs)
	connect := db.Connect()
	// role := models.Role{RoleName: "admins"}
	role := models.Role{}
	// connect.Create(&role)
	connect.Last(&role)

	c.JSON(http.StatusOK, role)
}
