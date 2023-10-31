build:
	docker compose build 

start: 
	docker compose up

start-dev:
	docker compose run -d -p 5432:5432 postgres
	air

start-detatch: build
	docker compose up -d

stop:
	docker compose down

