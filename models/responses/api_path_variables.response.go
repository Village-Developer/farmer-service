package responses

type ApiPathVariableResponse struct {
	ApiDocumentId     *string `json:"api_document_id"`
	Number            int     `json:"number"`
	PathVariableName  string  `json:"path_variable_name"`
	PathVariableValue string  `json:"path_variable_value"`
	Description       string  `json:"description"`
}
