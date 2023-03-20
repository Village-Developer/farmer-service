package responses

type ApiBodyRawResponse struct {
	ApiDocumentId *string `json:"api_document_id"`
	Body          string  `json:"body"`
	Description   string  `json:"description"`
	ContentType   string  `json:"content_type"`
}
