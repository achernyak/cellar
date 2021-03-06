// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/achernyak/cellar/design
// --out=$(GOPATH)/src/github.com/achernyak/cellar
// --version=v1.1.0-dirty
//
// API "cellar": bottle Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// ShowBottlePath computes a request path to the show action of bottle.
func ShowBottlePath(bottleID int) string {
	param0 := strconv.Itoa(bottleID)

	return fmt.Sprintf("/bottles/%s", param0)
}

// Get bottle by id
func (c *Client) ShowBottle(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBottleRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBottleRequest create the request corresponding to the show action endpoint of the bottle resource.
func (c *Client) NewShowBottleRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
