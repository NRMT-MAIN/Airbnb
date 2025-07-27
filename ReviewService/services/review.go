package services

import (
	"fmt"
	db "reviewservice/db/repositories"
	"reviewservice/dtos"
	"reviewservice/models"
)

type ReviewService interface {
	CreateReview(payload *dtos.CreateReviewDTO) (*models.Review , error)
}

type ReviewServiceImpl struct {
	ReviewRepositorty db.ReviewRepositorty
}

func NewReviewService(_ReviewRepository db.ReviewRepositorty) ReviewService{
	return &ReviewServiceImpl{
		ReviewRepositorty: _ReviewRepository,
	}
}

func (rs *ReviewServiceImpl) CreateReview(payload *dtos.CreateReviewDTO) (*models.Review , error) {
	review , err := rs.ReviewRepositorty.Create(payload)

	if err != nil {
		fmt.Println("Error in creating review in review service")
		return nil , err
	}

	return review , nil 
}