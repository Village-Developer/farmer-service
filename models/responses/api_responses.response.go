package responses

type ApiResponse struct {
	ApiDocumentId *string `json:"api_document_id"`
	Number        int     `json:"number"`
	StatusCode    int     `json:"statuc_code"`
	Body          string  `json:"response_body"`
	Description   string  `json:"description"`
}
