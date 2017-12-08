// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "EvelyApi": actions Resource Client
//
// Command:
// $ goagen
// --design=EvelyApi/design
// --out=$(GOPATH)/src/EvelyApi
// --version=v1.3.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// PingActionsPath computes a request path to the ping action of actions.
func PingActionsPath() string {

	return fmt.Sprintf("/develop/v1/actions/ping")
}

// 導通確認
func (c *Client) PingActions(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewPingActionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPingActionsRequest create the request corresponding to the ping action endpoint of the actions resource.
func (c *Client) NewPingActionsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
