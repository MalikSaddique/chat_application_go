package database

import "github.com/MalikSaddique/chat_application_go/models"

func (u *StorageImpl) SaveMessage(senderID string, msg models.SendMessageRequest) error {
	_, err := u.db.Exec(`
		INSERT INTO messages (sender_id, receiver_id, message, timestamp)
		VALUES ($1, $2, $3, NOW())`,
		senderID, msg.ReceiverID, msg.Message)

	return err
}

func (u *StorageImpl) FetchMessages(senderID, receiverID string) ([]models.Message, error) {
	query := `
		SELECT sender_id, receiver_id, message, timestamp
		FROM messages
		WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1)
		ORDER BY timestamp ASC
	`
	rows, err := u.db.Query(query, senderID, receiverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.Sender, &msg.Receiver, &msg.Message, &msg.Timestamp)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
