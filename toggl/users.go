package toggl

type UsersService struct {
	client *ApiClient
}

type User struct {
	Id 		uint 	`json:"id"`
	Email 	string 	`json:"email"`
}

func (c *UsersService) Me() (User, error) {
	return User{}, nil
}