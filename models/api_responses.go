package models

type ApiResponse struct {
	ApiResponseId *string `json:"api_respone_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	ApiDocumentId *string `json:"api_document_id"`
	Number        int     `json:"number"`
	StatusCode    int     `json:"statuc_code"`
	ResponseBody  string  `json:"response_body"`
	Description   string  `json:"description"`
}
