package controller

import (
	"log"
	"time"
	"EvelyApi/app"
	"EvelyApi/model"
	"github.com/goadesign/goa"
	mgo "gopkg.in/mgo.v2"
)

const (
	Lng = 0
	Lat = 1
)

/**
 * イベント情報をAPIのレスポンス形式に変換する
 * @param e イベント情報
 * @return  レスポンス形式に変換したイベント情報
 */
func ToEventMedia(e *model.EventModel) *app.Event {
	return &app.Event{
		ID:    e.ID,
		Title: e.Title,
		Host: &app.UserTiny{
			ID:   e.Host.ID,
			Name: e.Host.Name,
		},
		Body: e.Body,
		Place: &app.Location{
			Name: e.Place.Name,
			Lat:  e.Place.LngLat[Lat],
			Lng:  e.Place.LngLat[Lng],
		},
		UpdateDate: e.UpdateDate,
		UpcomingDate: &app.UpcomingDate{
			StartDate: e.UpcomingDate.StartDate,
			EndDate:   e.UpcomingDate.EndDate,
		},
		URL:  e.URL,
		Mail: e.Mail,
		Tel:  e.Tel,
	}
}

/**
 * イベント情報をAPIのレスポンス形式に変換する
 * @param e イベント情報
 * @return  レスポンス形式に変換したイベント情報
 */
func ToEventTinyMedia(e *model.EventModel) *app.EventTiny {
	return &app.EventTiny{
		ID:    e.ID,
		Title: e.Title,
		Host: &app.UserTiny{
			ID:   e.Host.ID,
			Name: e.Host.Name,
		},
		Place: &app.Location{
			Name: e.Place.Name,
			Lat:  e.Place.LngLat[Lat],
			Lng:  e.Place.LngLat[Lng],
		},
		UpcomingDate: &app.UpcomingDate{
			StartDate: e.UpcomingDate.StartDate,
			EndDate:   e.UpcomingDate.EndDate,
		},
	}
}

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
	user := GetLoginUser(ctx)
	eventID, err := c.db.NewEvent(user.ID)
	if err != nil {
		log.Printf("failed to create event: %s", err)
	}

	p := ctx.Payload
	event := &model.EventModel{
		ID:    eventID,
		Title: p.Title,
		Host: model.Host{
			ID:   user.ID,
			Name: user.Name,
		},
		Body: p.Body,
		Place: model.Location{
			Name:   p.Place.Name,
			LngLat: [2]float64{p.Place.Lng, p.Place.Lat},
		},
		UpdateDate: time.Now(),
		UpcomingDate: model.UpcomingDate{
			StartDate: p.UpcomingDate.StartDate,
			EndDate:   p.UpcomingDate.EndDate,
		},
		URL:  p.URL,
		Mail: p.Mail,
		Tel:  p.Tel,
	}

	err = c.db.SaveEvent(event)
	if err != nil {
		log.Printf("faild to save event: %s", err)
		return ctx.BadRequest()
	}

	return ctx.Created(ToEventMedia(event))
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
	events, err := c.db.GetEvents(ctx.Limit, ctx.Offset, ctx.Keyword, ctx.UserID)
	if err != nil {
		log.Printf("faild to search events: %s", err)
		return ctx.NotFound()
	}

	res := make(app.EventTinyCollection, len(events))
	for i := range events {
		res[i] = ToEventTinyMedia(&events[i])
	}
	return ctx.OKTiny(res)
	// EventsController_List: end_implement
}

// Show runs the show action.
func (c *EventsController) Show(ctx *app.ShowEventsContext) error {
	// EventsController_Show: start_implement

	// Put your logic here

	res := &app.Event{}
	return ctx.OK(res)
	// EventsController_Show: end_implement
}

// Update runs the update action.
func (c *EventsController) Update(ctx *app.UpdateEventsContext) error {
	// EventsController_Update: start_implement

	// Put your logic here

	res := &app.Event{}
	return ctx.OK(res)
	// EventsController_Update: end_implement
}
