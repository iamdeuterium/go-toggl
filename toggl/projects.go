package toggl

import (
	"fmt"
	"strings"
)

type ProjectsService struct {
	client *ApiClient
}

type Project struct {
	ID 			uint 	`json:"id"`
	Name 		string 	`json:"name"`
	WorkspaceID	uint	`json:"wid"`
	ClientID    uint	`json:"cid"`
}

type ProjectResponse struct {
	Project		Project `json:"data"`
}

type ProjectRequest struct {
	Project 	Project `json:"project"`
}

func (service *ProjectsService) GetByID(id uint) Project {
	response := new(ProjectResponse)

	service.client.DoRequest("GET", fmt.Sprintf("/projects/%d", id), nil, response)

	return response.Project
}

func (service *ProjectsService) GetByName(name string, workspaceID uint) (Project, bool) {
	projects := service.client.Workspaces.GetProjectsByWorkspaceId(workspaceID)

	for i := 0; i < len(*projects); i++ {
		project := (*projects)[i]

		if project.Name == name {
			return project, true
		}
	}

	return Project{}, false
}

func (service *ProjectsService) GetByNamePrefix(name string, workspaceID uint) ([]Project, bool) {
	projects := service.client.Workspaces.GetProjectsByWorkspaceId(workspaceID)

	result := []Project{}

	for i := 0; i < len(*projects); i++ {
		project := (*projects)[i]

		if strings.HasPrefix(strings.ToLower(project.Name), strings.ToLower(name)) {
			result = append(result, project)
		}
	}

	return result, false
}

func (service *ProjectsService) Create(project Project) Project {
	request 	:= ProjectRequest{Project: project}
	response	:= new(ProjectResponse)

	fmt.Println(request)

	service.client.DoRequest("POST", "/projects", request, response)

	return response.Project
}

func (service *ProjectsService) Update(project Project) Project {
	request		:= ProjectRequest{Project: project}
	response 	:= new(ProjectResponse)

	service.client.DoRequest("PUT", fmt.Sprintf("/projects/%d", project.ID), request, response)

	return response.Project
}

func (service *ProjectsService) Delete(project Project) {
	request	:= ProjectRequest{Project: project}

	service.client.DoRequest("DELETE", fmt.Sprintf("/projects/%d", project.ID), request, nil)
}