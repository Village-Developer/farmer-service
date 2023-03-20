package responses

type ApiQueryParamsResponse struct {
	ApiDocumentId   *string `json:"api_document_id"`
	Number          int     `json:"number"`
	QueryParamName  string  `json:"query_param_name"`
	QueryParamValue string  `json:"query_param_value"`
	Description     string  `json:"description"`
}
