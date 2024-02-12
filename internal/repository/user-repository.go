package repository

import (
	"errors"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (r *postgresRepository) GetUserByEmailAndPassword(
	email string,
	testPassword string,
) (int, string, error) {
	var id int
	var hashedPassword string

	query := `
		SELECT 
			id,
			password
		FROM users
		WHERE email = $1
	`

	row := r.DB.SQL.QueryRow(query, email)

	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}

func (r *postgresRepository) GetUserByID(id int) (models.User, error) {
	var user models.User

	query := `
						SELECT 
							id, 
							first_name, 
							last_name, 
							email, 
							phone,
							password, 
							access_level, 
							created_at, 
							updated_at
						FROM users 
						WHERE id = $1
					`

	row := r.DB.SQL.QueryRow(query, id)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.AccessLevel,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *postgresRepository) UpdateUser(user models.User) error {
	query := `
		UPDATE users 
		SET first_name = $1, last_name = $2, email = $3, access_level = $4, updated_at = $5
		WHERE id = $6
	`

	_, err := r.DB.SQL.Exec(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.AccessLevel,
		user.UpdatedAt,
		user.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
