package models

type ApiDocument struct {
	ApiDocumentId   *string `json:"api_document_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	ApiDocumentName string  `json:"api_document_name"`
	ApiProjectId    *string `json:"api_project_id"`
	ApiGroupId      *string `json:"api_group_id"`
	Method          string  `json:"method"`
	Path            string  `json:"path"`
}
