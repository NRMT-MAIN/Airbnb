package controllers

import (
	"fmt"
	"net/http"
	"reviewservice/dtos"
	"reviewservice/services"
	"reviewservice/utils"
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