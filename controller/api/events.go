package api

import (
	"EvelyApi/app"
	"EvelyApi/model"
    "EvelyApi/controller/parser"
	"github.com/goadesign/goa"
	"labix.org/v2/mgo"
	"log"
)

// EventsController implements the events resource.
type EventsController struct {
	*goa.Controller
	db *model.EventDB
}

// NewEventsController creates a events controller.
func NewEventsController(service *goa.Service, db *mgo.Database) *EventsController {
	return &EventsController{
		Controller: service.NewController("EventsController"),
		db:         model.NewEventDB(db),
	}
}

// Create runs the create action.
func (c *EventsController) Create(ctx *app.CreateEventsContext) error {
	// EventsController_Create: start_implement

	// Put your logic here
	claims := GetJWTClaims(ctx)
	user := &model.UserModel{
		ID:   claims["id"].(string),
		Name: claims["name"].(string),
	}
	payload := ctx.Payload

	eventID, err := c.db.NewEvent(user.ID, payload.UpcomingDate.StartDate)
	if err != nil {
		log.Printf("[EvelyApi] failed to create event: %s", err)
	}

	event := parser.ToEventModel(payload, eventID, user)

	err = c.db.SaveEvent(event)
	if err != nil {
		log.Printf("[EvelyApi] faild to save event: %s", err)
		return ctx.BadRequest()
	}

	return ctx.Created(parser.ToEventMedia(event))
	// EventsController_Create: end_implement
}

// Delete runs the delete action.
func (c *EventsController) Delete(ctx *app.DeleteEventsContext) error {
	// EventsController_Delete: start_implement

	// Put your logic here
	claims := GetJWTClaims(ctx)
	userID := claims["id"].(string)

	if userID != ctx.UserID {
		log.Printf("[EvelyApi] permission error")
		return ctx.Forbidden()
	}

	err := c.db.DeleteEvent(ctx.UserID, ctx.EventID)
	if err != nil {
		log.Printf("[EvelyApi] failed to delete event: %s", err)
		return ctx.NotFound()
	}
	return ctx.OK([]byte("Seccess!!"))
	// EventsController_Delete: end_implement
}

// List runs the list action.
func (c *EventsController) List(ctx *app.ListEventsContext) error {
	// EventsController_List: start_implement

	// Put your logic here
	events, err := c.db.GetEvents(ctx.Limit, ctx.Offset, model.WithKeyword(ctx.Keyword), model.WithUserID(ctx.UserID))
	if err != nil {
		log.Printf("[EvelyApi] faild to search events: %s", err)
		return ctx.NotFound()
	}

	res := make(app.EventTinyCollection, len(events))
	for i := range events {
		res[i] = parser.ToEventTinyMedia(events[i])
	}
	return ctx.OKTiny(res)
	// EventsController_List: end_implement
}

// Show runs the show action.
func (c *EventsController) Show(ctx *app.ShowEventsContext) error {
	// EventsController_Show: start_implement

	// Put your logic here
	event, err := c.db.GetEvent(ctx.UserID, ctx.EventID)
	if err != nil {
		log.Printf("[EvelyApi] faild to find event: %s", err)
		return ctx.NotFound()
	}
	return ctx.OK(parser.ToEventMedia(event))
	// EventsController_Show: end_implement
}

// Update runs the update action.
func (c *EventsController) Update(ctx *app.UpdateEventsContext) error {
	// EventsController_Update: start_implement

	// Put your logic here
	claims := GetJWTClaims(ctx)
	user := &model.UserModel{
		ID:   claims["id"].(string),
		Name: claims["name"].(string),
	}
	if user.ID != ctx.UserID {
		log.Printf("[EvelyApi] permission error")
		return ctx.Forbidden()
	}

	event := parser.ToEventModel(ctx.Payload, ctx.EventID, user)
	err := c.db.SaveEvent(event)
	if err != nil {
		log.Printf("[EvelyApi] faild to save event: %s", err)
		return ctx.NotFound()
	}
	return ctx.OK(parser.ToEventMedia(event))
	// EventsController_Update: end_implement
}
