package models

type ApiHeader struct {
	ApiHeaderId   *string `json:"api_header_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	ApiDocumentId *string `json:"api_document_id"`
	Number        int     `json:"number"`
	HeaderName    string  `json:"header_name"`
	HeaderValue   string  `json:"header_value"`
	Description   string  `json:"description"`
}
