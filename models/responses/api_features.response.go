package responses

type ApiFeatureResponse struct {
	ApiDocumentId *string `json:"api_document_id"`
	Number        int     `json:"number"`
	Feature       string  `json:"feature"`
}
