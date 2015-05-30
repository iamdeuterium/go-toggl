package toggl

import (
	"fmt"
)

type ClientsService struct {
	client *ApiClient
}

type Client struct {
	ID 				uint 	`json:"id"`
	Name 			string	`json:"name"`
	WorkspaceId		uint	`json:"wid"`
}

type ClientResponse struct {
	Client Client `json:"data"`
}

type ClientRequest struct {
	Client Client  `json:"client"`
}

func (service *ClientsService) All() *[]Client {
	clients := new([]Client)

	service.client.DoRequest("GET", "/clients", nil, clients)

	return clients
}

func (service *ClientsService) GetByID(id uint) Client {
	response := new(ClientResponse)

	service.client.DoRequest("GET", fmt.Sprintf("/clients/%d", id), nil, response)

	return response.Client
}

func (service *ClientsService) Create(client Client) Client {
	request 	:= ClientRequest{Client: client}
	response	:= new(ClientResponse)

	service.client.DoRequest("POST", "/clients", request, response)

	return response.Client
}

func (service *ClientsService) Update(client Client) Client {
	request		:= ClientRequest{Client: client}
	response 	:= new(ClientResponse)

	service.client.DoRequest("PUT", fmt.Sprintf("/clients/%d", client.ID), request, response)

	return response.Client
}

func (service *ClientsService) Delete(client Client) {
	request	:= ClientRequest{Client: client}

	service.client.DoRequest("DELETE", fmt.Sprintf("/clients/%d", client.ID), request, nil)
}


func (service *ClientsService) GetProjectsByClient(client Client) *[]Project {
	return service.GetProjectsByClientId(client.ID)
}

func (service *ClientsService) GetProjectsByClientId(id uint) *[]Project {
	projects := new([]Project)

	service.client.DoRequest("GET", fmt.Sprintf("/clients/%d/projects", id), nil, projects)

	return projects
}