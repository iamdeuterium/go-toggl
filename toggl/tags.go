package toggl

import (
	"fmt"
)

type TagsService struct {
	client *ApiClient
}

type Tag struct {
	Id 				uint 	`json:"id"`
	Name 			string	`json:"name"`
	WorkspaceId		uint	`json:"wid"`
}

type TagRequest struct {
	Tag				Tag		`json:"tag"`
}

type TagResponse struct {
	Tag				Tag		`json:"data"`
}

func (service *TagsService) Create(tag Tag) Tag {
	request 	:= TagRequest{Tag: tag}
	response	:= new(TagResponse)

	service.client.DoRequest("POST", "/tags", request, response)

	return response.Tag
}

func (service *TagsService) Update(tag Tag) Tag {
	request		:= TagRequest{Tag: tag}
	response 	:= new(TagResponse)

	service.client.DoRequest("PUT", fmt.Sprintf("/tags/%d", tag.Id), request, response)

	return response.Tag
}

func (service *TagsService) Delete(tag Tag) {
	request	:= TagRequest{Tag: tag}

	service.client.DoRequest("DELETE", fmt.Sprintf("/tags/%d", tag.Id), request, nil)
}