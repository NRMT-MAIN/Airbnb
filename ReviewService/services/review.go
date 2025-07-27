package services

import (
	"fmt"
	db "reviewservice/db/repositories"
	"reviewservice/dtos"
	"reviewservice/models"
)

type ReviewService interface {
	CreateReview(payload *dtos.CreateReviewDTO) (*models.Review , error)
	GetAllReviews() ([]*models.SingleReview , error)
	GetByIdReview(id string) (*models.SingleReview , error)
	DeleteByIdReview(id string) error
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

func (rs *ReviewServiceImpl) GetAllReviews() ([]*models.SingleReview , error) {
	reviews , err := rs.ReviewRepositorty.GetAll()

	if err != nil {
		fmt.Println("Error in Getting reviews" , err)
		return nil , err
	}

	return reviews , nil 
}

func (rs *ReviewServiceImpl) GetByIdReview(id string) (*models.SingleReview , error) {
	review , err := rs.ReviewRepositorty.GetById(id)

	if err != nil {
		fmt.Println("Error in getting the review by id" , err)
		return nil , err
	}
	return review , err
}

func (rs *ReviewServiceImpl) DeleteByIdReview(id string) error {
	err := rs.ReviewRepositorty.DeleteById(id)
	if err != nil {
		fmt.Println("Error in deleting the review" , err)
		return err
	}
	return nil 
}