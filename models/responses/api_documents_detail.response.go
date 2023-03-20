package responses

type ApiDocumentDetailResponse struct {
	ApiDocumentId   *string                   `json:"id"`
	ApiDocumentName string                    `json:"name"`
	Url             string                    `json:"url"`
	Method          string                    `json:"method"`
	Description     string                    `json:"description"`
	Features        []ApiFeatureResponse      `json:"features"`
	QueryParams     []ApiQueryParamsResponse  `json:"query_params"`
	PathVariables   []ApiPathVariableResponse `json:"path_variables"`
	Headers         []ApiHeaderResponse       `json:"headers"`
	Body            []ApiBodyResponse         `json:"body"`
	BodyRaw         []ApiBodyRawResponse      `json:"body_raw"`
	Responses       []ApiResponse             `json:"responses"`
}
