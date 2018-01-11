package main

import (
	"EvelyApi/app"
	"github.com/goadesign/goa"
)

// EventsController implements the events resource.
type EventsController struct {
	*goa.Controller
}

// NewEventsController creates a events controller.
func NewEventsController(service *goa.Service) *EventsController {
	return &EventsController{Controller: service.NewController("EventsController")}
}

// Create runs the create action.
func (c *EventsController) Create(ctx *app.CreateEventsContext) error {
	// EventsController_Create: start_implement

	// Put your logic here

	return nil
	// EventsController_Create: end_implement
}

// Delete runs the delete action.
func (c *EventsController) Delete(ctx *app.DeleteEventsContext) error {
	// EventsController_Delete: start_implement

	// Put your logic here

	return nil
	// EventsController_Delete: end_implement
}

// List runs the list action.
func (c *EventsController) List(ctx *app.ListEventsContext) error {
	// EventsController_List: start_implement

	// Put your logic here

	res := app.EventCollection{}
	return ctx.OK(res)
	// EventsController_List: end_implement
}

// Modify runs the modify action.
func (c *EventsController) Modify(ctx *app.ModifyEventsContext) error {
	// EventsController_Modify: start_implement

	// Put your logic here

	res := &app.Event{}
	return ctx.OK(res)
	// EventsController_Modify: end_implement
}

// Nearby runs the nearby action.
func (c *EventsController) Nearby(ctx *app.NearbyEventsContext) error {
	// EventsController_Nearby: start_implement

	// Put your logic here

	res := app.EventCollection{}
	return ctx.OK(res)
	// EventsController_Nearby: end_implement
}

// Notify runs the notify action.
func (c *EventsController) Notify(ctx *app.NotifyEventsContext) error {
	// EventsController_Notify: start_implement

	// Put your logic here

	return nil
	// EventsController_Notify: end_implement
}

// Show runs the show action.
func (c *EventsController) Show(ctx *app.ShowEventsContext) error {
	// EventsController_Show: start_implement

	// Put your logic here

	res := app.EventCollection{}
	return ctx.OK(res)
	// EventsController_Show: end_implement
}

// Update runs the update action.
func (c *EventsController) Update(ctx *app.UpdateEventsContext) error {
	// EventsController_Update: start_implement

	// Put your logic here

	return nil
	// EventsController_Update: end_implement
}
