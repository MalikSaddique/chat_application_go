package db

import "github.com/MalikSaddique/chat_application_go/models"

func (u *StorageImpl) FindUserByEmail(email string) (*models.UserLogin, error) {
	var user models.UserLogin

	err := u.db.QueryRow("SELECT id, email, password FROM users WHERE email=$1", email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		log.Fatal("Error: User not exist")
		return nil, err
	}

	return &user, nil

}
