package repository

import (
	"database/sql"

	"belajar-golang-api-basic/models"
)

func GetUsers(db *sql.DB) ([]models.User, error) {

	query := `SELECT id, nama, posisi_id, perusahaan FROM users`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Nama,
			&user.PosisiID,
			&user.Perusahaan,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func InsertUser(db *sql.DB, user models.User) error {

	query := `
	INSERT INTO users (nama, posisi_id, perusahaan)
	VALUES (?, ?, ?)
	`
	//Exec untuk mengeksekusi INSERT, UPDATE, DELETE
	_, err := db.Exec(
		query,
		user.Nama,
		user.PosisiID,
		user.Perusahaan,
	) 

	return err
}