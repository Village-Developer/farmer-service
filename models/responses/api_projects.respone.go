package responses

type ApiProjectResponse struct {
	ApiProjectId   *string       `json:"id"`
	ApiProjectName *string       `json:"name"`
	Type           string        `json:"type" gorm:"-"`
	ApiGroups      []interface{} `json:"children" gorm:"-"`
}
