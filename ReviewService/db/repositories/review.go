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
	GetAll() ([]*models.SingleReview , error)
	GetById(id string) (*models.SingleReview , error)
	DeleteById(id string) error
}

type ReviewRepositortyImpl struct {
	db *sql.DB
}

func NewReviewRepository(_db *sql.DB) ReviewRepositorty {
	return &ReviewRepositortyImpl{
		db: _db,
	}
}

func (r *ReviewRepositortyImpl) DeleteById(id string) error {
	query := "DELETE FROM REVIEW WHERE ID = ?"

	row , err := r.db.Exec(query , id)

	if err != nil {
		fmt.Println("Error in deleting the review" , err)
		return err
	}

	rowsAffected, err := row.RowsAffected()
    if err != nil {
        return fmt.Errorf("could not check affected rows: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("review with id %s not found", id) // Explicit "not found" error
    }
	fmt.Println("Review deleted successfully, rows affected:", rowsAffected)
	return nil
}

func (r *ReviewRepositortyImpl) GetById(id string) (*models.SingleReview , error) {
	query := "SELECT * FROM REVIEW WHERE ID = ?"

	row := r.db.QueryRow(query , id)

	review := &models.SingleReview{}
	err := row.Scan(&review.ID , &review.BookingID , &review.Comment , &review.Rating , &review.Created_At , &review.Updated_At)

	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("No Records Found!")
			return nil , err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
	}

	return review , nil	
}

func (r *ReviewRepositortyImpl) GetAll() ([]*models.SingleReview , error){
	query := "SELECT * FROM REVIEW;"

	rows , err := r.db.Query(query)

	if err != nil {
		fmt.Println("Error in fetching all review" , err) 
		return nil , err
	}

	defer rows.Close()

	var results []*models.SingleReview

	for rows.Next() {
		review :=  &models.SingleReview{}
		err := rows.Scan(&review.ID , &review.BookingID , &review.Comment , &review.Rating , &review.Created_At , &review.Updated_At)

		if err != nil {
			fmt.Println("Error fetching the review" , err)
			return nil , err
		}
		results = append(results, review)
	}

	if err := rows.Err() ; err != nil {
		fmt.Println("Error with rows") ; 
		return nil , err
	}

	return results , nil
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