MIGRATIONS_FOLDER = ./internal/database/migrations
DATABASE_URL = mysql://root:password@(localhost:3306)/egoist

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database $(DATABASE_URL) -verbose up
migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database $(DATABASE_URL) -verbose down
dev:
	docker compose --env-file ./local.env up -d
down:
	docker compose down
docker.build:
	docker build -t egoist .
	