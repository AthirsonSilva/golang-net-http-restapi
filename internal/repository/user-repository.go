package repository

import (
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
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

func (r *postgresRepository) InsertUser(user models.User) (int, error) {
	var id int
	query := `
		INSERT INTO users (first_name, last_name, email, phone, password, access_level, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	_, err := r.DB.SQL.Exec(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Password,
		user.AccessLevel,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return id, err
	}

	id, _, err = r.GetUserByEmailAndPassword(user.Email, user.Password)
	if err != nil {
		return id, err
	}

	return id, nil
}
