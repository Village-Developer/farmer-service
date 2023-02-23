package models

type ApiPermission struct {
	ApiProjectId *string `json:"api_project_id"`
	UserId       string  `json:"user_id"`
	PermissionId string  `json:"permission_id"`
}
