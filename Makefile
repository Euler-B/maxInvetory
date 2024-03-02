include .env

run:
	go run main.go

database-up:
	docker run \
	-d \
	--name mariaDB \
	-p 3306:3306 \
	-e MARIADB_ROOT_PASSWORD=${DB_PASSWORD} \
	-e MARIADB_DATABASE=${DB_NAME} \
	mariadb:10.7.4
database-down:
	docker rm -f mariaDB


start-db:
	make database-up
	echo Espere unos segundos mientras se levanta la instancia
	sleep 15
	make migrate-db


migrate-db:
	migrate -path database/migrations/ -database ${DB_URL} up

migrate-db-rollback:
	migrate -path database/migrations/ -database ${DB_URL} down

migrate-fix:
	migrate -path database/migrations/ -database ${DB_URL} force 1

tests:
	 go test -v ./...

# TO_DO:
# 1.- refactoizar star-db
# 2.- refactorizar sentencias migrate down

