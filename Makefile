
run_all:
	docker-compose up


run_app:
	docker rmi pahan_app -f
	swag init -g /cmd/main/main.go --parseInternal --parseDependency
	docker-compose up -d app

run_pg:
	docker-compose up -d psql

swag:
	swag init -g cmd/main/ --parseInternal --parseDependency