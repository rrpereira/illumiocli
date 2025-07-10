package illumiocli

// Client represents an Illumio PCE client
type Client struct {
	Host    string
	User    string
	Key     string
	Version string
}

// NewClient creates a new Illumio PCE client
func NewClient(host, user, key, version string) *Client {
	return &Client{
		Host:    host,
		User:    user,
		Key:     key,
		Version: version,
	}
}
