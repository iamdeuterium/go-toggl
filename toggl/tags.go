package toggl

type TagsService struct {
	client *ApiClient
}

type Tag struct {
	Id 				uint 	`json:"id"`
	Name 			string	`json:"name"`
	WorkspaceId		uint	`json:"wid"`
}