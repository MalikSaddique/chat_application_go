DB_URL=postgres://postgres:1122@localhost:5433/chat_app_go?sslmode=disable

migrate:
		goose -dir db/migrations postgres "$(DB_URL)" up
