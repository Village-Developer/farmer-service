package responses

type ApiBodyResponse struct {
	ApiDocumentId *string `json:"api_document_id"`
	Number        int     `json:"number"`
	ApiBodyName   string  `json:"api_body_name"`
	ApiBodyValue  string  `json:"api_body_value"`
	Description   string  `json:"description"`
	ContentType   string  `json:"content_type"`
}
