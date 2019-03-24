include .env

run:
	docker-compose up

test:
	go test muskooters/station