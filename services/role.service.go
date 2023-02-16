package services

import (
	"village-developer.com/farmer/configs"
	"village-developer.com/farmer/models"
)

type RoleService struct{}

var config = new(configs.Configs)
var db = config.Connect()

func (RoleService) AddRoleService(role models.Role) string {
	err := db.Create(&role).Error
	if err != nil {
		return err.Error()
	} else {
		return "Success"
	}
}

func (RoleService) GetAllRolesService(role []models.Role) ([]models.Role, int) {
	err := db.Find(&role).Error
	if err != nil {
		return role, 204
	} else {
		return role, 200
	}
}

func (RoleService) GetRoleByIdService(role models.Role, id string) (models.Role, int) {
	err := db.First(&role, id).Error
	if err != nil {
		return role, 400
	} else {
		return role, 200
	}
}
