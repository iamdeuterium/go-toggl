package toggl

type ClientsService struct {
	client *ApiClient
}

type Client struct {
	Id 				uint 	`json:"id"`
	Name 			string	`json:"name"`
	WorkspaceId		uint	`json:"wid"`
}

func (service *ClientsService) test() {

}