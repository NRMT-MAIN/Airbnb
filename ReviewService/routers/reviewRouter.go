package routers

import (
	"reviewservice/controllers"

	"github.com/go-chi/chi/v5"
)

type ReviewRouter struct {
	ReviewController *controllers.ReviewController
}

func NewReviewRouter(_ReviewController *controllers.ReviewController) Router {
	return &ReviewRouter{
		ReviewController: _ReviewController,
	}
}

func (rr *ReviewRouter) Register(r chi.Router){
	r.Post("/review" , rr.ReviewController.CreateReview)
	r.Get("/review" , rr.ReviewController.GetAllReviews)
	r.Get("/review/{id}" , rr.ReviewController.GetByIdReview)
	r.Delete("/review/{id}" , rr.ReviewController.DeleteById)
} 