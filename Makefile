ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=test password=test dbname=test host=localhost port=5432 sslmode=disable
endif

INTERNAL_PKG_PATH=$(CURDIR)/internal/pkg
MIGRATION_FOLDER=$(INTERNAL_PKG_PATH)/db/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql
.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY:
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down


.PHONY: compose-up
compose-up:
	docker-compose build
	docker-compose up -d postgres

.PHONY: compose-rm
compose-rm:
	docker-compose down


test:
	go test -coverprofile=./internal/pkg/server/coverage.out ./internal/pkg/server
	go tool cover -html=./internal/pkg/server/coverage.out -o ./internal/pkg/server/coverage.html

generate_proto:
	protoc \
		--proto_path=internal/proto/ \
		--go_out=internal/pb/ \
		--go-grpc_out=internal/pb/ \
		internal/proto/*.proto

