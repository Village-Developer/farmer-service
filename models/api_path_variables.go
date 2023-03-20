package models

type ApiPathVariable struct {
	ApiPathVariableId *string `json:"api_path_variable_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	ApiDocumentId     *string `json:"api_document_id"`
	Number            int     `json:"number"`
	PathVariableName  string  `json:"path_variable_name"`
	PathVariableValue string  `json:"path_variable_value"`
	Description       string  `json:"description"`
}
