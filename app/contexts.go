// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "EvelyApi": Application Contexts
//
// Command:
// $ goagen
// --design=EvelyApi/design
// --out=$(GOPATH)/src/EvelyApi
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// PingActionsContext provides the actions ping action context.
type PingActionsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewPingActionsContext parses the incoming request URL and body, performs validations and creates the
// context used by the actions controller ping action.
func NewPingActionsContext(ctx context.Context, r *http.Request, service *goa.Service) (*PingActionsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := PingActionsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *PingActionsContext) OK(resp []byte) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	}
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *PingActionsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// SigninAuthContext provides the auth signin action context.
type SigninAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *LoginPayload
}

// NewSigninAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller signin action.
func NewSigninAuthContext(ctx context.Context, r *http.Request, service *goa.Service) (*SigninAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SigninAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SigninAuthContext) OK(r *Token) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.token+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *SigninAuthContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}
