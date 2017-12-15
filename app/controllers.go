// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "EvelyApi": Application Controllers
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
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Muxer
	SendMail(*SendMailAuthContext) error
	Signin(*SigninAuthContext) error
	Signup(*SignupAuthContext) error
	VerifyToken(*VerifyTokenAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service *goa.Service, ctrl AuthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/develop/v1/auth/signup/send_mail", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/develop/v1/auth/signin", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/develop/v1/auth/signup", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/develop/v1/auth/signup/verify_token", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSendMailAuthContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*SignupPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.SendMail(rctx)
	}
	h = handleAuthOrigin(h)
	service.Mux.Handle("POST", "/api/develop/v1/auth/signup/send_mail", ctrl.MuxHandler("send_mail", h, unmarshalSendMailAuthPayload))
	service.LogInfo("mount", "ctrl", "Auth", "action", "SendMail", "route", "POST /api/develop/v1/auth/signup/send_mail")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSigninAuthContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*LoginPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Signin(rctx)
	}
	h = handleAuthOrigin(h)
	service.Mux.Handle("POST", "/api/develop/v1/auth/signin", ctrl.MuxHandler("signin", h, unmarshalSigninAuthPayload))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Signin", "route", "POST /api/develop/v1/auth/signin")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSignupAuthContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Signup(rctx)
	}
	h = handleAuthOrigin(h)
	service.Mux.Handle("POST", "/api/develop/v1/auth/signup", ctrl.MuxHandler("signup", h, unmarshalSignupAuthPayload))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Signup", "route", "POST /api/develop/v1/auth/signup")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewVerifyTokenAuthContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.VerifyToken(rctx)
	}
	h = handleAuthOrigin(h)
	service.Mux.Handle("GET", "/api/develop/v1/auth/signup/verify_token", ctrl.MuxHandler("verify_token", h, nil))
	service.LogInfo("mount", "ctrl", "Auth", "action", "VerifyToken", "route", "GET /api/develop/v1/auth/signup/verify_token")
}

// handleAuthOrigin applies the CORS response headers corresponding to the origin.
func handleAuthOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8888/swaggerui") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalSendMailAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalSendMailAuthPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &signupPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalSigninAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalSigninAuthPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &loginPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalSignupAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalSignupAuthPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &userPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// EventsController is the controller interface for the Events actions.
type EventsController interface {
	goa.Muxer
	Create(*CreateEventsContext) error
	Delete(*DeleteEventsContext) error
	List(*ListEventsContext) error
	Show(*ShowEventsContext) error
	Update(*UpdateEventsContext) error
}

// MountEventsController "mounts" a Events resource controller on the given service.
func MountEventsController(service *goa.Service, ctrl EventsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/develop/v1/events", ctrl.MuxHandler("preflight", handleEventsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/develop/v1/events/:user_id/:event_id", ctrl.MuxHandler("preflight", handleEventsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/develop/v1/events/:user_id", ctrl.MuxHandler("preflight", handleEventsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateEventsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*EventPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleEventsOrigin(h)
	service.Mux.Handle("POST", "/api/develop/v1/events", ctrl.MuxHandler("create", h, unmarshalCreateEventsPayload))
	service.LogInfo("mount", "ctrl", "Events", "action", "Create", "route", "POST /api/develop/v1/events", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteEventsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleEventsOrigin(h)
	service.Mux.Handle("DELETE", "/api/develop/v1/events/:user_id/:event_id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Events", "action", "Delete", "route", "DELETE /api/develop/v1/events/:user_id/:event_id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListEventsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleEventsOrigin(h)
	service.Mux.Handle("GET", "/api/develop/v1/events", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Events", "action", "List", "route", "GET /api/develop/v1/events")
	service.Mux.Handle("GET", "/api/develop/v1/events/:user_id", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Events", "action", "List", "route", "GET /api/develop/v1/events/:user_id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowEventsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleEventsOrigin(h)
	service.Mux.Handle("GET", "/api/develop/v1/events/:user_id/:event_id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Events", "action", "Show", "route", "GET /api/develop/v1/events/:user_id/:event_id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateEventsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*EventPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("jwt", h, "api:access")
	h = handleEventsOrigin(h)
	service.Mux.Handle("PUT", "/api/develop/v1/events/:user_id/:event_id", ctrl.MuxHandler("update", h, unmarshalUpdateEventsPayload))
	service.LogInfo("mount", "ctrl", "Events", "action", "Update", "route", "PUT /api/develop/v1/events/:user_id/:event_id", "security", "jwt")
}

// handleEventsOrigin applies the CORS response headers corresponding to the origin.
func handleEventsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8888/swaggerui") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateEventsPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateEventsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &eventPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateEventsPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateEventsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &eventPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swaggerui/*filepath", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swaggerui/*filepath", "public/swaggerui")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swaggerui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "public/swaggerui", "route", "GET /swaggerui/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swaggerui/", "public/swaggerui/index.html")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swaggerui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "public/swaggerui/index.html", "route", "GET /swaggerui/")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8888/swaggerui") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// UsersController is the controller interface for the Users actions.
type UsersController interface {
	goa.Muxer
	Show(*ShowUsersContext) error
}

// MountUsersController "mounts" a Users resource controller on the given service.
func MountUsersController(service *goa.Service, ctrl UsersController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/develop/v1/users/:user_id", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleUsersOrigin(h)
	service.Mux.Handle("GET", "/api/develop/v1/users/:user_id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Users", "action", "Show", "route", "GET /api/develop/v1/users/:user_id")
}

// handleUsersOrigin applies the CORS response headers corresponding to the origin.
func handleUsersOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8888/swaggerui") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Expose-Headers", "X-Time")
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
