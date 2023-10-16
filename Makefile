build:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css
	templ generate -path ./templates/ 
	docker-compose build 

start: 
	docker-compose up

start-detatch: build
	docker-compose up -d

stop:
	docker-compose down

