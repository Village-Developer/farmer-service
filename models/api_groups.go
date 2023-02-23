package models

type ApiGroup struct {
	ApiGroupId    *string       `json:"api_group_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	ApiGroupName  string        `json:"api_group_name"`
	ApiProjectId  *string       `json:"api_project_id"`
	ParentGroupId *string       `json:"parent_group_id"`
	ApiGroups     []ApiDocument `json:"api_children" gorm:"-"`
}
