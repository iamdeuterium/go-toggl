package toggl

import "fmt"

type WorkspacesService struct {
	client *ApiClient
}

type Workspace struct {
	Id		uint	`json:"id"`
	Name	string 	`json:"name"`
}

type WorkspaceResponse struct {
	Workspace Workspace `json:"data"`
}

type WorkspaceRequest struct {
	Workspace Workspace `json:"workspace"`
}

func (service *WorkspacesService) All() *[]Workspace {
	workspaces := new([]Workspace)

	service.client.DoRequest("GET", "/workspaces", nil, workspaces)

	return workspaces
}

func (service *WorkspacesService) GetByID(id uint) Workspace {
	response := new(WorkspaceResponse)

	service.client.DoRequest("GET", fmt.Sprintf("/workspaces/%d", id), nil, response)

	return response.Workspace
}

func (service *WorkspacesService) Update(workspace Workspace) Workspace {
	request := WorkspaceRequest{Workspace: workspace}
	response := new(WorkspaceResponse)

	service.client.DoRequest("PUT", fmt.Sprintf("/workspaces/%d", workspace.Id), request, response)

	return response.Workspace
}

func (service *WorkspacesService) GetUsersByWorkspace(workspace Workspace) *[]User {
	return service.GetUsersByWorkspaceId(workspace.Id)
}

func (service *WorkspacesService) GetUsersByWorkspaceId(id uint) *[]User {
	users := new([]User)

	service.client.DoRequest("GET", fmt.Sprintf("/workspaces/%d/users", id), nil, users)

	return users
}

func (service *WorkspacesService) GetClientsByWorkspace(workspace Workspace) *[]Client {
	return service.GetClientsByWorkspaceId(workspace.Id)
}

func (service *WorkspacesService) GetClientsByWorkspaceId(id uint) *[]Client {
	clients := new([]Client)

	service.client.DoRequest("GET", fmt.Sprintf("/workspaces/%d/clients", id), nil, clients)

	return clients
}

func (service *WorkspacesService) GetProjectsByWorkspace(workspace Workspace) *[]Project {
	return service.GetProjectsByWorkspaceId(workspace.Id)
}

func (service *WorkspacesService) GetProjectsByWorkspaceId(id uint) *[]Project {
	projects := new([]Project)

	service.client.DoRequest("GET", fmt.Sprintf("/workspaces/%d/projects", id), nil, projects)

	return projects
}

func (service *WorkspacesService) GetTagsByWorkspace(workspace Workspace) *[]Tag {
	return service.GetTagsByWorkspaceId(workspace.Id)
}

func (service *WorkspacesService) GetTagsByWorkspaceId(id uint) *[]Tag {
	tags := new([]Tag)

	service.client.DoRequest("GET", fmt.Sprintf("/workspaces/%d/tags", id), nil, tags)

	return tags
}