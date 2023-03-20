package responses

type ApiHeaderResponse struct {
	ApiDocumentId *string `json:"api_document_id"`
	Number        int     `json:"number"`
	HeaderName    string  `json:"header_name"`
	HeaderValue   string  `json:"header_value"`
	Description   string  `json:"description"`
}
