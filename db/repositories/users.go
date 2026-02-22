package db

import (
	"Auth/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById() (*models.User, error)
	Create() error
}

type UserRepositoryImp struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImp{db: _db}
}

func (u *UserRepositoryImp) GetById() (*models.User, error) {
	fmt.Println("Fetching user in UserRepository")
	//prepare query
	query := "SELECT id,username,email,created_at,updated_at FROM users WHERE id=?"
	//execute query
	row := u.db.QueryRow(query, 1) //ready row
	//process the result
	user := &models.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {

			fmt.Println("No user found with given ID")
			return nil, err
		} else {
			fmt.Println("Error scanning user: ", err)
			return nil, err

		}

	}
	fmt.Println("User fetched successfully", user)

	return user, nil
}

func (u *UserRepositoryImp) Create() error {
	query := "INSERT INTO users (username,email,password) VALUES (?,?,?)"
	result, err := u.db.Exec(query, "testuser", "test@gmail.com", "test")
	if err != nil {
		fmt.Println("Error inserting user: ", err)
		return err
	}
	rowsAffected, rowErr := result.RowsAffected()
	if rowErr != nil {
		fmt.Println("Error getting affected rows: ", rowErr)
		return rowErr
	}
	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not created:")
		return nil

	}
	fmt.Println("User created successfully :")

	return nil

}
