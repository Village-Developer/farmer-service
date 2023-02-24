package responses

type ApiProjectResponse struct {
	ApiProjectId   *string       `json:"id"`
	ApiProjectName *string       `json:"name"`
	ApiGroups      []interface{} `json:"children" gorm:"-"`
}
