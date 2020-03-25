package tuneuptechnology

// Setup the HTTP client
const APIBaseUrl = "https://app.tuneuptechnology.com/api/"

type Client struct {
	Auth 		string `json:"auth"`
	APIKey		string `json:"api_key"`
}
