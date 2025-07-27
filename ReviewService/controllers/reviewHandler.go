package controllers

import (
	"fmt"
	"net/http"
	"reviewservice/dtos"
	"reviewservice/services"
	"reviewservice/utils"

	"github.com/go-chi/chi/v5"
)

type ReviewController struct {
	ReviewService services.ReviewService
}

func NewReviewController(_ReviewService services.ReviewService) *ReviewController{
	return &ReviewController{
		ReviewService: _ReviewService,
	}
}

func (rc *ReviewController) CreateReview(w http.ResponseWriter , r *http.Request){
	fmt.Println("Create Review called in User Controller.")
	
	var payload dtos.CreateReviewDTO

	if jsonErr := utils.ReadJsonBody(r , &payload) ; jsonErr != nil {
		utils.WriteErrorJsonResponse(w , "JSON Reading Error" , http.StatusInternalServerError , jsonErr)
		return
	}

	review , err := rc.ReviewService.CreateReview(&payload)

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error creating the review" , http.StatusInternalServerError , err)
		return
	}

	utils.WriteSuccessJsonResponse(w , "Review Created!" , http.StatusCreated , review)
}

func (rc *ReviewController) GetAllReviews(w http.ResponseWriter , r *http.Request){
	reviews , err := rc.ReviewService.GetAllReviews()

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in getting reviews" , http.StatusInternalServerError  , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Reviews fetched successfully!" , http.StatusOK , reviews)
}

func (rc *ReviewController) GetByIdReview(w http.ResponseWriter , r *http.Request){
	id := chi.URLParam(r , "id")
	review , err := rc.ReviewService.GetByIdReview(id)

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in getting review" , http.StatusInternalServerError  , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Review fetched successfully!" , http.StatusOK , review)
}


func (rc *ReviewController) DeleteById(w http.ResponseWriter , r *http.Request) {
	id := chi.URLParam(r , "id")
	err := rc.ReviewService.DeleteByIdReview(id)

	if err != nil {
		utils.WriteErrorJsonResponse(w , "Error in deleting review" , http.StatusInternalServerError  , err)
		return
	}
	utils.WriteSuccessJsonResponse(w , "Review deleted successfully!" , http.StatusOK ,"Deleted")
}