package funcs

import (
	"net/url"
	"path"
)

// GetComponent ...
func GetComponent(s string) (*url.URL, error) {
	u, err := url.Parse(s)
	return u, err
}

// GetHostName get host name
func GetHostName(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return u.Host, nil
}

// URLJoin join url
func URLJoin(source, target string) string {
	return path.Join(source, target)
}
