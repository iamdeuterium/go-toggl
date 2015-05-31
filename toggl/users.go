package toggl

type UsersService struct {
	client *ApiClient
}

type User struct {
	Id 			uint		`json:"id"`
	Email 		string 		`json:"email"`
	ApiToken	string		`json:"api_token"`

	DefaultWorkspaceID uint `json:"default_wid"`
	Workspaces	[]Workspace
}

type UserResponse struct {
	User		User		`json:"data"`
}

type UserRequest struct {
	User		User		`json:"user"`
}

func (service *UsersService) Current() User {
	response := new(UserResponse)

	service.client.DoRequest("GET", "/me", nil, response)

	return response.User
}

func (service *UsersService) UpdateCurrent(user User) User {
	request		:= UserRequest{User: user}
	response 	:= new(UserResponse)

	service.client.DoRequest("PUT", "/me", request, response)

	return response.User
}