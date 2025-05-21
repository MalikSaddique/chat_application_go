package db

import (
	"fmt"

	"github.com/MalikSaddique/chat_application_go/models"
)

func (u *StorageImpl) SearchUser(email string) ([]models.UserResponse, error) {
	query := `
		SELECT id, email
		FROM users
		WHERE email ILIKE $1
	`
	rows, err := u.db.Query(query, fmt.Sprintf("%%%s%%", email))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserResponse

	for rows.Next() {
		var user models.UserResponse
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
