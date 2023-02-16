package services

import (
	"village-developer.com/farmer/configs"
	"village-developer.com/farmer/models"
)

type RoleService struct{}

func (RoleService) AddRoleService(role models.Role) string {
	db := new(configs.Configs)
	connect := db.Connect()
	err := connect.Create(&role).Error
	if err != nil {
		return err.Error()
	} else {
		return "Success"
	}
}

func (RoleService) GetAllRolesService(role []models.Role) []models.Role {
	db := new(configs.Configs)
	connect := db.Connect()
	connect.Find(&role)
	return role
}

func (RoleService) GetRoleByIdService(role models.Role, id string) (models.Role, int) {
	db := new(configs.Configs)
	connect := db.Connect()
	err := connect.First(&role, id).Error
	if err != nil {
		return role, 400
	} else {
		return role, 200
	}
}
