package toggl

import (
	"fmt"
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