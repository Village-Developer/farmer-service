package models

type ApiBody struct {
	ApiBodyId          *string `json:"api_body_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	ApiDocumentId      *string `json:"api_document_id"`
	Number             int     `json:"number"`
	ApiBodyName        string  `json:"api_body_name"`
	APiBodyValue       string  `json:"api_body_value"`
	ApiBodyDescription string  `json:"api_body_description"`
	ContentType        string  `json:"content_type"`
}
