DB_URL=postgres://user:password@localhost:5432/burguer-db?sslmode=disable

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(NAME)

migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" down 1

.PHONY: migrate-create migrate-up migrate-down