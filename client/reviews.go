// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "EvelyApi": reviews Resource Client
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
	"strconv"
)

// CreateReviewsPath computes a request path to the create action of reviews.
func CreateReviewsPath(eventID string) string {
	param0 := eventID

	return fmt.Sprintf("/api/develop/v2/reviews/%s", param0)
}

// レビュー投稿
func (c *Client) CreateReviews(ctx context.Context, path string, payload *ReviewPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateReviewsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateReviewsRequest create the request corresponding to the create action endpoint of the reviews resource.
func (c *Client) NewCreateReviewsRequest(ctx context.Context, path string, payload *ReviewPayload, contentType string) (*http.Request, error) {
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
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ListReviewsPath computes a request path to the list action of reviews.
func ListReviewsPath(eventID string) string {
	param0 := eventID

	return fmt.Sprintf("/api/develop/v2/reviews/%s", param0)
}

// レビューの一覧取得
func (c *Client) ListReviews(ctx context.Context, path string, limit *int, offset *int) (*http.Response, error) {
	req, err := c.NewListReviewsRequest(ctx, path, limit, offset)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListReviewsRequest create the request corresponding to the list action endpoint of the reviews resource.
func (c *Client) NewListReviewsRequest(ctx context.Context, path string, limit *int, offset *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if limit != nil {
		tmp38 := strconv.Itoa(*limit)
		values.Set("limit", tmp38)
	}
	if offset != nil {
		tmp39 := strconv.Itoa(*offset)
		values.Set("offset", tmp39)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
