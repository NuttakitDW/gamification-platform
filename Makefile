compose:
	@docker-compose -f docker-compose.yml up --build
gen:
	swag init -g ./cmd/api/main.go

