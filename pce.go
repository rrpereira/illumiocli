package illumiocli

import "github.com/go-resty/resty/v2"

type PCE struct {
	Host    string
	User    string
	Key     string
	Version string
}

func NewPCE(host, user, key, version string) *PCE {
	return &PCE{
		Host:    host,
		User:    user,
		Key:     key,
		Version: version,
	}
}

func (c *PCE) NewRestyClient() *resty.Client {
	return resty.
		New().
		SetBaseURL(c.Host+"/api/"+c.Version).
		SetBasicAuth(c.User, c.Key)
}
