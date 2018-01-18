package api

import (
	"EvelyApi/app"
	"EvelyApi/controller/parser"
	. "EvelyApi/middleware"
	"EvelyApi/model"
	. "EvelyApi/model/collection"
	. "EvelyApi/model/document"
	"labix.org/v2/mgo/bson"
	"github.com/goadesign/goa"
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

	// JWTからユーザー情報を取得
	uid, err := GetLoginID(ctx)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	user, _ := c.db.Users.FindOne(Keys{"id": uid})

	// レビューを保存するイベントを取得
	if !bson.IsObjectIdHex(ctx.EventID) {
        return ctx.BadRequest(goa.ErrBadRequest("\"" + ctx.EventID + "\" is not n event ID."))
    }
	keys := Keys{"_id": bson.ObjectIdHex(ctx.EventID)}
	e, err := c.db.Events.FindOne(keys)
	if err != nil {
		return ctx.NotFound(goa.ErrNotFound(err))
	}

	// レビューを作成
	p := ctx.Payload
	r := parser.ToReviewModel(p, bson.NewObjectId(), user)
	err = c.db.Reviews.Save(r, Keys{"_id": r.ID})
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	// 作成したレビューのIDをイベントのレビューID配列に追加保存する
	e.Reviews = append(e.Reviews, r.ID)
	err = c.db.Events.Save(e, keys)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	return ctx.Created(parser.ToReviewMedia(r))
}

// List runs the list action.
func (c *ReviewsController) List(ctx *app.ListReviewsContext) error {

	// イベントIDからレビューのID配列を取得する
	e, err := c.db.Events.FindOne(Keys{"_id": bson.ObjectIdHex(ctx.EventID)})
	if err != nil {
		return ctx.NotFound(goa.ErrNotFound(err))
	}
	// offsetからlimit数のレビューを取得
	var reviews []*ReviewModel
	for i := ctx.Offset; i < len(e.Reviews) && i < ctx.Limit; i++ {
		r, err := c.db.Reviews.FindOne(Keys{"_id": e.Reviews[i]})
		if err == nil {
			reviews = append(reviews, r)
		}
	}
	// レビュー情報をレスポンス形式に変換して返す
	res := make(app.ReviewCollection, len(reviews))
	for i := range reviews {
		res[i] = parser.ToReviewMedia(reviews[i])
	}
	return ctx.OK(res)
	// ReviewsController_List: end_implement
}
