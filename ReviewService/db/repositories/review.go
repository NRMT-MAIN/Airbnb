package db

import (
	"database/sql"
	"fmt"
	"reviewservice/dtos"
	"reviewservice/models"
	"strconv"
)

type ReviewRepositorty interface {
	Create(payload *dtos.CreateReviewDTO) (*models.Review , error)
	// GetAll() ([]*models.Review , error)
	// GetById(id string) (*models.Review , error)
	// DeleteById(id string) (*models.Review , error)
}

type ReviewRepositortyImpl struct {
	db *sql.DB
}

func NewReviewRepository(_db *sql.DB) ReviewRepositorty {
	return &ReviewRepositortyImpl{
		db: _db,
	}
}

func (r *ReviewRepositortyImpl) Create(payload *dtos.CreateReviewDTO) (*models.Review ,error) {
	query := "INSERT INTO REVIEW (BOOKING_ID , COMMENT , RATING) VALUES(? , ? , ?);"

	result , err := r.db.Exec(query , payload.BookingID , payload.Comment , payload.Rating)

	if err != nil {
		fmt.Println("Error in creating the review :" , err)
		return nil , err
	}

	lastEnteredId , err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error in getting last entered id :" , lastEnteredId)
		return  nil , err
	}

	bookingid , _ := strconv.Atoi(payload.BookingID)

	review := &models.Review{
		ID: int(lastEnteredId),
		BookingID: bookingid,
		Comment: payload.Comment,
		Rating: payload.Rating,
	}

	fmt.Println("Review Created Succesfully!") ; 
	return review , nil
}