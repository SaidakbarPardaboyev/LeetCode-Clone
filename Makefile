DATABASE_URL="postgres://postgres:root@localhost:5432/leetcode?sslmode=disable"

mig-create:
	@if [ -z "$(name)" ]; then \
		read -p "Enter migration name: " name; \
	fi; \
	migrate create -ext sql -dir migrations -seq $$name

mig-up:
	@migrate -database "$(DATABASE_URL)" -path migrations up

mig-down:
	@migrate -database "$(DATABASE_URL)" -path migrations down

mig-force:
	@if [ -z "$(version)" ]; then \
		read -p "Enter migration version: " version; \
	fi; \
	migrate -database "$(DATABASE_URL)" -path migrations force $$version

runS:
	@go run cmd/mainS.go

tidy:
	@go mod tidy
	@go mod vendor

drop-tables:
	@psql $(DATABASE_URL) -a -f drop_objects.sql
