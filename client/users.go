// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "EvelyApi": users Resource Client
//
// Command:
// $ goagen
// --design=EvelyApi/design
// --out=$(GOPATH)/src/EvelyApi
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ShowUsersPath computes a request path to the show action of users.
func ShowUsersPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/api/develop/v2/users/%s", param0)
}

// アカウント情報取得
func (c *Client) ShowUsers(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowUsersRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowUsersRequest create the request corresponding to the show action endpoint of the users resource.
func (c *Client) NewShowUsersRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateUsersPath computes a request path to the update action of users.
func UpdateUsersPath() string {

	return fmt.Sprintf("/api/develop/v2/users/update/token")
}

// インスタンスIDの登録・更新
// 認証ありで登録ユーザーを、認証なしでゲストユーザーを登録・更新する
func (c *Client) UpdateUsers(ctx context.Context, path string, payload *TokenPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateUsersRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateUsersRequest create the request corresponding to the update action endpoint of the users resource.
func (c *Client) NewUpdateUsersRequest(ctx context.Context, path string, payload *TokenPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.OptionalJWTSigner != nil {
		if err := c.OptionalJWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
