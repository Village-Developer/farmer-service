package responses

type ApiDocumentResponse struct {
	ApiDocumentId   *string `json:"id"`
	ApiDocumentName string  `json:"name"`
	ApiProjectId    *string `json:"project_id"`
	ApiGroupId      *string `json:"group_id"`
	Method          string  `json:"method"`
	Path            string  `json:"-"`
	Description     string  `json:"-"`
	Feature         string  `json:"-"`
}
