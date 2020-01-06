package fetch

import (
	"crypto/tls"
	"time"
)

// Options http client options
// default option
//  method: GET
//  body: nil
//  header: {"Accept-Encoding": "gzip,deflate", "Accept": "*/*"}
//  timeout: 20s
type Options struct {
	Method    string
	Body      []byte
	Header    map[string]string
	Timeout   time.Duration
	TLSConfig *tls.Config
}

// NewDefaultOptions create a default options
func NewDefaultOptions() Options {
	return Options{
		Method: "GET",
		Header: map[string]string{
			"Accept-Encoding": "gzip,deflate",
			"Accept":          "*/*",
		},
		Body:      nil,
		Timeout:   20 * time.Second,
		TLSConfig: &tls.Config{},
	}
}
