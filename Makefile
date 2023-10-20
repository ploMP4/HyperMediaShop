build:
	docker-compose build 

start: 
	docker-compose up

start-detatch: build
	docker-compose up -d

stop:
	docker-compose down

