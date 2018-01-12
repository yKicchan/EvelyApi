package api

import (
	"EvelyApi/app"
	"EvelyApi/model"
	"github.com/goadesign/goa"
)

// PinsController implements the pins resource.
type PinsController struct {
	*goa.Controller
	db *model.EvelyDB
}

// NewPinsController creates a pins controller.
func NewPinsController(service *goa.Service, db *model.EvelyDB) *PinsController {
	return &PinsController{
        Controller: service.NewController("PinsController"),
		db:         db,
    }
}

// Off runs the off action.
func (c *PinsController) Off(ctx *app.OffPinsContext) error {
	// PinsController_Off: start_implement

	// Put your logic here

	return nil
	// PinsController_Off: end_implement
}

// On runs the on action.
func (c *PinsController) On(ctx *app.OnPinsContext) error {
	// PinsController_On: start_implement

	// Put your logic here

	return nil
	// PinsController_On: end_implement
}
