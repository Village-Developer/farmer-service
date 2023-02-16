package models

type Role struct {
	RoleId   int `gorm:"primaryKey"`
	RoleName string
}
