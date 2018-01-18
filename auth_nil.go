package swift

import (
	"net/http"
)

type nilAuth struct {
	storageUrl string
}
// v1 Authentication - make request
func (auth *nilAuth) Request(c *Connection) (*http.Request, error) {
	return nil, nil
}

// v1 Authentication - read response
func (auth *nilAuth) Response(resp *http.Response) error {
	return nil
}

// v1 Authentication - read storage url
func (auth *nilAuth) StorageUrl(Internal bool) string {
	return auth.storageUrl
}

// v1 Authentication - read auth token
func (auth *nilAuth) Token() string {
	return ""
}

// v1 Authentication - read cdn url
func (auth *nilAuth) CdnUrl() string {
	return ""
}
