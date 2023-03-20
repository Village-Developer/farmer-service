package models

type ApiBodyRaw struct {
	ApiBodyId     *string `json:"api_body_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	ApiDocumentId *string `json:"api_document_id"`
	Body          string  `json:"body"`
	Description   string  `json:"description"`
	ContentType   string  `json:"content_type"`
}
