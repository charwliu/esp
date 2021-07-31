.POSIX:
.PHONY: clean test security build run js-build

APP_NAME = esp
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgres://esp:mZZak_3rQ@localhost/esp_db?sslmode=disable
FRONTEND_DIR = frontend
COMMIT = $$( git rev-parse --short HEAD )
DIST_DIR = dist
YARN = yarn
YARN_FLAGS = --cwd $(FRONTEND_DIR)
YARN_INSTALL_FLAGS = $(YARN_FLAGS) --network-timeout 120000 --silent\
    --ignore-engines --ignore-optional --ignore-platform\
    --ignore-scripts

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

linter:
	golangci-lint run

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag.init build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

docker.run: docker.network docker.postgres swag.init docker.fiber docker.redis migrate.up

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name cgapp-fiber \
		--network dev-network \
		-p 5000:5000 \
		fiber

docker.postgres:
	docker run --rm -d \
		--name cgapp-postgres \
		--network dev-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

docker.redis:
	docker run --rm -d \
		--name cgapp-redis \
		--network dev-network \
		-p 6379:6379 \
		redis

docker.stop: docker.stop.fiber docker.stop.postgres docker.stop.redis

docker.stop.fiber:
	docker stop cgapp-fiber

docker.stop.postgres:
	docker stop cgapp-postgres

docker.stop.redis:
	docker stop cgapp-redis

swag.init:
	swag init

