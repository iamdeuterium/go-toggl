package toggl

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ApiClient struct {
	httpClient *http.Client
	apiToken string

	Users *UsersService
	Projects *ProjectsService
	Workspaces *WorkspacesService
	Clients *ClientsService
	Tags *TagsService
}

func NewClient(apiToken string, httpClient *http.Client) *ApiClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &ApiClient{apiToken: apiToken, httpClient: httpClient}

	client.Users		= &UsersService{client: client}
	client.Projects 	= &ProjectsService{client: client}
	client.Workspaces	= &WorkspacesService{client: client}
	client.Clients		= &ClientsService{client: client}
	client.Tags			= &TagsService{client: client}

	return client
}

func (client *ApiClient) NewRequest(method string, url string, body interface{}) *http.Request {
	var bodyBuffer io.ReadWriter

	if body != nil {
		bodyBuffer = new(bytes.Buffer)
		json.NewEncoder(bodyBuffer).Encode(body)
	}

	req, _ := http.NewRequest(method, "https://www.toggl.com/api/v8" + url, bodyBuffer)

	req.SetBasicAuth(client.apiToken, "api_token")
	req.Header.Add("Content-Type", "application/json")

	return req
}

func (client *ApiClient) DoRequest(method string, url string, body interface{}, responseStruct interface{}) {
	request := client.NewRequest(method, url, body)

	resp, _ := client.httpClient.Do(request)

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(responseStruct)
}