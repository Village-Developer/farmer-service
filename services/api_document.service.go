package services

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"village-developer.com/farmer/configs"
	"village-developer.com/farmer/models"
	"village-developer.com/farmer/models/responses"
)

type ApiDocumentService struct{}

var config = new(configs.Configs)
var db = config.Connect()

func (ApiDocumentService) AddApiDocumentProjectService(c *gin.Context) (string, int) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err.Error(), 400
	}
	var body map[string]interface{}
	json.Unmarshal(jsonData, &body)
	project := models.ApiProject{ApiProjectName: body["api_project_name"].(string)}
	err = db.Create(&project).Error
	if err != nil {
		return err.Error(), 400
	}
	test := []models.ApiProject{}
	db.Table("api_projects").Select("api_project_id, created_at").Order("created_at DESC").Limit(1).Scan(&test)
	user_id := c.Request.Header.Get("user_id")
	permission := models.ApiPermission{ApiProjectId: test[0].ApiProjectId, UserId: user_id, PermissionId: "da04d4c5-aeca-11ed-acea-0242ac130002"}
	err = db.Create(&permission).Error
	if err != nil {
		return err.Error(), 400
	} else {
		return "Insert Success", 200
	}
}

func (ApiDocumentService) GetAllApiDocumentService(c *gin.Context) ([]responses.ApiProjectResponse, int) {
	user_id := c.Request.Header.Get("user_id")

	project := []responses.ApiProjectResponse{}
	group := []responses.ApiGroupResponse{}
	api := []responses.ApiDocumentResponse{}
	db.Table("api_projects").Select("*").Joins("left join api_permissions on api_projects.api_project_id = api_permissions.api_project_id").Where("api_permissions.user_id = ?", user_id).Order("api_projects.api_project_name ASC").Scan(&project)
	db.Table("api_groups").Select("*").Joins("left join api_permissions on api_groups.api_project_id = api_permissions.api_project_id").Where("api_permissions.user_id = ?", user_id).Order("api_groups.api_group_name ASC").Scan(&group)
	db.Table("api_documents").Select("*").Joins("left join api_permissions on api_documents.api_project_id = api_permissions.api_project_id").Where("api_permissions.user_id = ?", user_id).Order("api_documents.api_document_name ASC").Scan(&api)
	for i := 0; i < len(project); i++ {
		project[i].Type = "project"
		if project[i].ApiGroups == nil {
			project[i].ApiGroups = []interface{}{}
		}
		for j := 0; j < len(group); j++ {
			group[j].Type = "group"
			if group[j].ApiGroups == nil {
				group[j].ApiGroups = []responses.ApiDocumentResponse{}
			}
			if *project[i].ApiProjectId == *group[j].ApiProjectId {
				for k := 0; k < len(api); k++ {
					api[k].Type = "api"
					if api[k].ApiGroupId != nil {
						if *group[j].ApiGroupId == *api[k].ApiGroupId {
							group[j].ApiGroups = append(group[j].ApiGroups, api[k])
						}
					}
				}
				project[i].ApiGroups = append(project[i].ApiGroups, group[j])
			}
		}
		for k := 0; k < len(api); k++ {
			api[k].Type = "api"
			if *project[i].ApiProjectId == *api[k].ApiProjectId && api[k].ApiGroupId == nil {
				project[i].ApiGroups = append(project[i].ApiGroups, api[k])
			}
		}
	}
	return project, 200
}

func (ApiDocumentService) DeleteApiDocumentProjectService(c *gin.Context) (string, int) {
	id := c.Param("api_project_id")
	db.Where("api_project_id = ?", id).Delete(&models.ApiDocument{})
	db.Where("api_project_id = ?", id).Delete(&models.ApiGroup{})
	db.Where("api_project_id = ?", id).Delete(&models.ApiPermission{})
	err := db.Where("api_project_id = ?", id).Delete(&models.ApiProject{}).Error
	if err != nil {
		return err.Error(), 400
	} else {
		return "Delete Success", 200
	}
}

func (ApiDocumentService) AddApiDocumentGroupService(c *gin.Context) (string, int) {
	id := c.Param("api_project_id")
	project_id := &id
	err := db.Where("api_project_id = ?", id).First(&models.ApiProject{}).Error
	if err != nil {
		return err.Error(), 400
	}
	var parent *string
	if c.Query("parent_id") == "" {
		parent = nil
	} else {
		id := c.Query("parent_id")
		parent = &id
	}
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err.Error(), 400
	}
	var body map[string]interface{}
	json.Unmarshal(jsonData, &body)
	group := models.ApiGroup{ApiGroupName: body["api_group_name"].(string), ApiProjectId: project_id, ParentGroupId: parent}
	err = db.Create(&group).Error
	if err != nil {
		return err.Error(), 400
	} else {
		return "Insert Success", 200
	}
}

func (ApiDocumentService) AddApiDocumentsService(c *gin.Context) (string, int) {
	id := c.Param("api_project_id")
	project_id := &id
	err := db.Where("api_project_id = ?", id).First(&models.ApiProject{}).Error
	if err != nil {
		return err.Error(), 400
	}
	parent := c.Query("parent_id")
	parent_id := &parent
	if parent == "" {
		parent_id = nil
	}
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err.Error(), 400
	}
	var body map[string]interface{}
	json.Unmarshal(jsonData, &body)
	fmt.Println(body)
	document := models.ApiDocument{ApiDocumentName: body["api_document_name"].(string), ApiProjectId: project_id, ApiGroupId: parent_id, Method: body["method"].(string), Path: body["path"].(string)}
	err = db.Create(&document).Error
	if err != nil {
		return err.Error(), 400
	} else {
		return "Insert Success", 200
	}
}

func (ApiDocumentService) DeleteApiDocumentGroupService(c *gin.Context) (string, int) {
	id := c.Param("api_group_id")
	db.Where("api_group_id = ?", id).Delete(&models.ApiDocument{})
	err := db.Where("api_group_id = ?", id).Delete(&models.ApiGroup{}).Error
	if err != nil {
		return err.Error(), 400
	} else {
		return "Delete Success", 200
	}
}

func (ApiDocumentService) DeleteApiDocumentService(c *gin.Context) (string, int) {
	id := c.Param("api_document_id")
	err := db.Where("api_document_id = ?", id).Delete(&models.ApiDocument{}).Error
	if err != nil {
		return err.Error(), 400
	} else {
		return "Delete Success", 200
	}
}

func (ApiDocumentService) GetApiDocumentService(c *gin.Context) ([]responses.ApiDocumentDetailResponse, int) {
	id := c.Param("api_document_id")
	var document models.ApiDocument
	features := []models.ApiFeature{}
	headers := []models.ApiHeader{}
	query_params := []models.ApiQueryParams{}
	path_variable := []models.ApiPathVariable{}
	body := []models.ApiBody{}
	body_raw := []models.ApiBodyRaw{}
	api_response := []models.ApiResponse{}
	err := db.Where("api_document_id = ?", id).First(&document).Error
	db.Where("api_document_id = ?", id).Find(&features)
	db.Where("api_document_id = ?", id).Find(&headers)
	db.Where("api_document_id = ?", id).Find(&query_params)
	db.Where("api_document_id = ?", id).Find(&path_variable)
	db.Where("api_document_id = ?", id).Find(&api_response)
	db.Where("api_document_id = ?", id).Table("api_body").Scan(&body)
	db.Where("api_document_id = ?", id).Table("api_body_raw").Scan(&body_raw)
	if err != nil {
		return nil, 400
	} else {
		var response responses.ApiDocumentDetailResponse
		featureResponses := []responses.ApiFeatureResponse{}
		headerResponses := []responses.ApiHeaderResponse{}
		queryParamResponses := []responses.ApiQueryParamsResponse{}
		pathVariableResponses := []responses.ApiPathVariableResponse{}
		apiResponses := []responses.ApiResponse{}
		bodyResponses := []responses.ApiBodyResponse{}
		bodyRawResponses := []responses.ApiBodyRawResponse{}

		for _, feature := range features {
			var featureResponseItem responses.ApiFeatureResponse
			featureResponseItem.ApiDocumentId = feature.ApiDocumentId
			featureResponseItem.Number = feature.Number
			featureResponseItem.Feature = feature.Feature
			featureResponses = append(featureResponses, featureResponseItem)
		}

		for _, header := range headers {
			var headerResponseItem responses.ApiHeaderResponse
			headerResponseItem.ApiDocumentId = header.ApiDocumentId
			headerResponseItem.Number = header.Number
			headerResponseItem.HeaderName = header.HeaderName
			headerResponseItem.HeaderValue = header.HeaderValue
			headerResponseItem.Description = header.Description
			headerResponses = append(headerResponses, headerResponseItem)
		}

		for _, query_param := range query_params {
			var queryParamResponseItem responses.ApiQueryParamsResponse
			queryParamResponseItem.ApiDocumentId = query_param.ApiDocumentId
			queryParamResponseItem.Number = query_param.Number
			queryParamResponseItem.QueryParamName = query_param.QueryParamName
			queryParamResponseItem.QueryParamValue = query_param.QueryParamValue
			queryParamResponseItem.Description = query_param.Description
			queryParamResponses = append(queryParamResponses, queryParamResponseItem)
		}

		for _, path_variable := range path_variable {
			var pathVariableResponseItem responses.ApiPathVariableResponse
			pathVariableResponseItem.ApiDocumentId = path_variable.ApiDocumentId
			pathVariableResponseItem.Number = path_variable.Number
			pathVariableResponseItem.PathVariableName = path_variable.PathVariableName
			pathVariableResponseItem.PathVariableValue = path_variable.PathVariableValue
			pathVariableResponseItem.Description = path_variable.Description
			pathVariableResponses = append(pathVariableResponses, pathVariableResponseItem)
		}

		for _, response := range api_response {
			var responseResponseItem responses.ApiResponse
			responseResponseItem.ApiDocumentId = response.ApiDocumentId
			responseResponseItem.Number = response.Number
			responseResponseItem.StatusCode = response.StatusCode
			responseResponseItem.Body = response.ResponseBody
			responseResponseItem.Description = response.Description
			apiResponses = append(apiResponses, responseResponseItem)
		}

		for _, body := range body {
			var bodyResponseItem responses.ApiBodyResponse
			bodyResponseItem.ApiDocumentId = body.ApiDocumentId
			bodyResponseItem.Number = body.Number
			bodyResponseItem.ApiBodyName = body.ApiBodyName
			bodyResponseItem.ApiBodyValue = body.APiBodyValue
			bodyResponseItem.Description = body.ApiBodyDescription
			bodyResponses = append(bodyResponses, bodyResponseItem)
		}

		for _, body_raw := range body_raw {
			var bodyRawResponseItem responses.ApiBodyRawResponse
			bodyRawResponseItem.ApiDocumentId = body_raw.ApiDocumentId
			bodyRawResponseItem.Body = body_raw.Body
			bodyRawResponseItem.Description = body_raw.Description
			bodyRawResponseItem.ContentType = body_raw.ContentType
			bodyRawResponses = append(bodyRawResponses, bodyRawResponseItem)
		}

		response.ApiDocumentId = document.ApiDocumentId
		response.ApiDocumentName = document.ApiDocumentName
		response.Method = document.Method
		response.Url = document.Path
		response.Features = featureResponses
		response.Headers = headerResponses
		response.QueryParams = queryParamResponses
		response.PathVariables = pathVariableResponses
		response.Responses = apiResponses
		response.Body = bodyResponses
		response.BodyRaw = bodyRawResponses

		return []responses.ApiDocumentDetailResponse{response}, 200
	}
}
