package api

import (
	"EvelyApi/app"
	"github.com/goadesign/goa"
	"EvelyApi/model"
)

// ReviewsController implements the reviews resource.
type ReviewsController struct {
	*goa.Controller
	db *model.EvelyDB
}

// NewReviewsController creates a reviews controller.
func NewReviewsController(service *goa.Service, db *model.EvelyDB) *ReviewsController {
	return &ReviewsController{
        Controller: service.NewController("ReviewsController"),
		db:         db,
    }
}

// Create runs the create action.
func (c *ReviewsController) Create(ctx *app.CreateReviewsContext) error {
	// ReviewsController_Create: start_implement

	// Put your logic here

	return nil
	// ReviewsController_Create: end_implement
}

// List runs the list action.
func (c *ReviewsController) List(ctx *app.ListReviewsContext) error {
	// ReviewsController_List: start_implement

	// Put your logic here

	res := app.ReviewCollection{}
	return ctx.OK(res)
	// ReviewsController_List: end_implement
}
