package models

type ApiQueryParams struct {
	ApiQueryParamsId *string `json:"api_query_params_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	ApiDocumentId    *string `json:"api_document_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	Number           int     `json:"number"`
	QueryParamName   string  `json:"query_param_name"`
	QueryParamValue  string  `json:"query_param_value"`
	Description      string  `json:"description"`
}
