package toggl

type ProjectsService struct {
	client *ApiClient
}

type Project struct {
	Id 			uint 	`json:"id"`
	Name 		string 	`json:"name"`
	WorkspaceID	uint	`json:"wid"`
	ClientID    uint	`json:"cid"`
}