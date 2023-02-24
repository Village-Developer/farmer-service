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
	db.Table("api_projects").Select("*").Joins("left join api_permissions on api_projects.api_project_id = api_permissions.api_project_id").Where("api_permissions.user_id = ?", user_id).Scan(&project)
	db.Table("api_groups").Select("*").Joins("left join api_permissions on api_groups.api_project_id = api_permissions.api_project_id").Where("api_permissions.user_id = ?", user_id).Scan(&group)
	db.Table("api_documents").Select("*").Joins("left join api_permissions on api_documents.api_project_id = api_permissions.api_project_id").Where("api_permissions.user_id = ?", user_id).Scan(&api)
	for i := 0; i < len(project); i++ {
		for j := 0; j < len(group); j++ {
			if *project[i].ApiProjectId == *group[j].ApiProjectId {
				for k := 0; k < len(api); k++ {
					if api[k].ApiGroupId != nil {
						if *group[j].ApiGroupId == *api[k].ApiGroupId {
							group[j].ApiGroups = append(group[j].ApiGroups, api[k])
							project[i].ApiGroups = append(project[i].ApiGroups, group[j])
						}
					}
				}
			}
		}
		for k := 0; k < len(api); k++ {
			if *project[i].ApiProjectId == *api[k].ApiProjectId && api[k].ApiGroupId == nil {
				project[i].ApiGroups = append(project[i].ApiGroups, api[k])
			}
		}
	}
	return project, 200
}

func (ApiDocumentService) DeleteApiDocumentService(c *gin.Context) (string, int) {
	id := c.Param("id")
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
	document := models.ApiDocument{ApiDocumentName: body["api_document_name"].(string), ApiProjectId: project_id, ApiGroupId: parent_id, Method: body["method"].(string), Path: body["path"].(string), Description: body["description"].(string), Feature: body["feature"].(string)}
	err = db.Create(&document).Error
	if err != nil {
		return err.Error(), 400
	} else {
		return "Insert Success", 200
	}
}
