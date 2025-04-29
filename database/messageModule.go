package database

import "github.com/MalikSaddique/chat_application_go/models"

func (u *StorageImpl) SaveMessage(senderID string, msg models.SendMessageRequest) error {
	_, err := u.db.Exec(`
		INSERT INTO messages (sender_id, receiver_id, message, timestamp)
		VALUES ($1, $2, $3, NOW())`,
		senderID, msg.ReceiverID, msg.Message)

	return err
}
