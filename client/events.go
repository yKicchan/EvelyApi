// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "EvelyApi": events Resource Client
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

// CreateEventsPath computes a request path to the create action of events.
func CreateEventsPath() string {

	return fmt.Sprintf("/api/develop/v2/events")
}

// イベント作成
func (c *Client) CreateEvents(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateEventsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateEventsRequest create the request corresponding to the create action endpoint of the events resource.
func (c *Client) NewCreateEventsRequest(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Request, error) {
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

// DeleteEventsPath computes a request path to the delete action of events.
func DeleteEventsPath(eventID string) string {
	param0 := eventID

	return fmt.Sprintf("/api/develop/v2/events/%s", param0)
}

// イベント削除
func (c *Client) DeleteEvents(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteEventsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteEventsRequest create the request corresponding to the delete action endpoint of the events resource.
func (c *Client) NewDeleteEventsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
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

// ListEventsPath computes a request path to the list action of events.
func ListEventsPath() string {

	return fmt.Sprintf("/api/develop/v2/events")
}

// イベント複数取得
func (c *Client) ListEvents(ctx context.Context, path string, keyword *string, limit *int, offset *int) (*http.Response, error) {
	req, err := c.NewListEventsRequest(ctx, path, keyword, limit, offset)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListEventsRequest create the request corresponding to the list action endpoint of the events resource.
func (c *Client) NewListEventsRequest(ctx context.Context, path string, keyword *string, limit *int, offset *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if keyword != nil {
		values.Set("keyword", *keyword)
	}
	if limit != nil {
		tmp24 := strconv.Itoa(*limit)
		values.Set("limit", tmp24)
	}
	if offset != nil {
		tmp25 := strconv.Itoa(*offset)
		values.Set("offset", tmp25)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ModifyEventsPath computes a request path to the modify action of events.
func ModifyEventsPath(eventID string) string {
	param0 := eventID

	return fmt.Sprintf("/api/develop/v2/events/%s", param0)
}

// イベント編集
func (c *Client) ModifyEvents(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Response, error) {
	req, err := c.NewModifyEventsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewModifyEventsRequest create the request corresponding to the modify action endpoint of the events resource.
func (c *Client) NewModifyEventsRequest(ctx context.Context, path string, payload *EventPayload, contentType string) (*http.Request, error) {
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
	req, err := http.NewRequest("PUT", u.String(), &body)
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

// MyListEventsPath computes a request path to the my_list action of events.
func MyListEventsPath() string {

	return fmt.Sprintf("/api/develop/v2/events/my_list")
}

// 自分のイベント一覧を取得する
func (c *Client) MyListEvents(ctx context.Context, path string, limit *int, offset *int) (*http.Response, error) {
	req, err := c.NewMyListEventsRequest(ctx, path, limit, offset)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewMyListEventsRequest create the request corresponding to the my_list action endpoint of the events resource.
func (c *Client) NewMyListEventsRequest(ctx context.Context, path string, limit *int, offset *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if limit != nil {
		tmp26 := strconv.Itoa(*limit)
		values.Set("limit", tmp26)
	}
	if offset != nil {
		tmp27 := strconv.Itoa(*offset)
		values.Set("offset", tmp27)
	}
	u.RawQuery = values.Encode()
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

// NearbyEventsPath computes a request path to the nearby action of events.
func NearbyEventsPath() string {

	return fmt.Sprintf("/api/develop/v2/events/nearby")
}

// 近くのイベントを検索する
func (c *Client) NearbyEvents(ctx context.Context, path string, lat float64, lng float64, range_ int, limit *int, offset *int) (*http.Response, error) {
	req, err := c.NewNearbyEventsRequest(ctx, path, lat, lng, range_, limit, offset)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewNearbyEventsRequest create the request corresponding to the nearby action endpoint of the events resource.
func (c *Client) NewNearbyEventsRequest(ctx context.Context, path string, lat float64, lng float64, range_ int, limit *int, offset *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	tmp28 := strconv.FormatFloat(lat, 'f', -1, 64)
	values.Set("lat", tmp28)
	tmp29 := strconv.FormatFloat(lng, 'f', -1, 64)
	values.Set("lng", tmp29)
	tmp30 := strconv.Itoa(range_)
	values.Set("range", tmp30)
	if limit != nil {
		tmp31 := strconv.Itoa(*limit)
		values.Set("limit", tmp31)
	}
	if offset != nil {
		tmp32 := strconv.Itoa(*offset)
		values.Set("offset", tmp32)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// NotifyByInstanceIDEventsPath computes a request path to the notify_by_instance_id action of events.
func NotifyByInstanceIDEventsPath() string {

	return fmt.Sprintf("/api/develop/v2/events/notify/by_instance_id")
}

// 近くにイベントがあればインスタンスID宛に通知する
func (c *Client) NotifyByInstanceIDEvents(ctx context.Context, path string, payload *NotifyByInstanceIDPayload, contentType string) (*http.Response, error) {
	req, err := c.NewNotifyByInstanceIDEventsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewNotifyByInstanceIDEventsRequest create the request corresponding to the notify_by_instance_id action endpoint of the events resource.
func (c *Client) NewNotifyByInstanceIDEventsRequest(ctx context.Context, path string, payload *NotifyByInstanceIDPayload, contentType string) (*http.Request, error) {
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
	return req, nil
}

// NotifyByUserIDEventsPath computes a request path to the notify_by_user_id action of events.
func NotifyByUserIDEventsPath() string {

	return fmt.Sprintf("/api/develop/v2/events/notify/by_user_id")
}

// 近くにイベントがあればユーザーのデバイス全てに通知する
func (c *Client) NotifyByUserIDEvents(ctx context.Context, path string, payload *NotifyByUserIDPayload, contentType string) (*http.Response, error) {
	req, err := c.NewNotifyByUserIDEventsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewNotifyByUserIDEventsRequest create the request corresponding to the notify_by_user_id action endpoint of the events resource.
func (c *Client) NewNotifyByUserIDEventsRequest(ctx context.Context, path string, payload *NotifyByUserIDPayload, contentType string) (*http.Request, error) {
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

// PinEventsPath computes a request path to the pin action of events.
func PinEventsPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/api/develop/v2/events/pin/%s", param0)
}

// ユーザーのピンしたイベント一覧を取得する
func (c *Client) PinEvents(ctx context.Context, path string, limit *int, offset *int) (*http.Response, error) {
	req, err := c.NewPinEventsRequest(ctx, path, limit, offset)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPinEventsRequest create the request corresponding to the pin action endpoint of the events resource.
func (c *Client) NewPinEventsRequest(ctx context.Context, path string, limit *int, offset *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if limit != nil {
		tmp33 := strconv.Itoa(*limit)
		values.Set("limit", tmp33)
	}
	if offset != nil {
		tmp34 := strconv.Itoa(*offset)
		values.Set("offset", tmp34)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowEventsPath computes a request path to the show action of events.
func ShowEventsPath() string {

	return fmt.Sprintf("/api/develop/v2/events/detail")
}

// イベント情報取得
func (c *Client) ShowEvents(ctx context.Context, path string, ids []string) (*http.Response, error) {
	req, err := c.NewShowEventsRequest(ctx, path, ids)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowEventsRequest create the request corresponding to the show action endpoint of the events resource.
func (c *Client) NewShowEventsRequest(ctx context.Context, path string, ids []string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	for _, p := range ids {
		tmp35 := p
		values.Add("ids", tmp35)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateEventsPath computes a request path to the update action of events.
func UpdateEventsPath() string {

	return fmt.Sprintf("/api/develop/v2/events/update")
}

// イベントの開催フラグを更新する
func (c *Client) UpdateEvents(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewUpdateEventsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateEventsRequest create the request corresponding to the update action endpoint of the events resource.
func (c *Client) NewUpdateEventsRequest(ctx context.Context, path string) (*http.Request, error) {
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
