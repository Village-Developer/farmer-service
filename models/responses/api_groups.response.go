package responses

type ApiGroupResponse struct {
	ApiGroupId    *string               `json:"id"`
	ApiGroupName  string                `json:"name"`
	ApiProjectId  *string               `json:"project_id"`
	ParentGroupId *string               `json:"group_id"`
	ApiGroups     []ApiDocumentResponse `json:"children" gorm:"-"`
}
