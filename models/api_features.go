package models

type ApiFeature struct {
	ApiFeatureId  *string `json:"api_feature_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	ApiDocumentId *string `json:"api_document_id" gorm:"primaryKey;autoIncrement:false;type:uuid;default:uuid_generate_v4()"`
	Number        int     `json:"number"`
	Feature       string  `json:"feature"`
}
