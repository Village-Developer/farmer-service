package models

import "time"

type ApiProject struct {
	ApiProjectId   *string       `json:"api_project_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	ApiProjectName string        `json:"api_project_name"`
	CreatedAt      time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	ApiGroups      []interface{} `json:"api_children" gorm:"-"`
}
