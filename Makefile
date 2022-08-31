up:
	docker-compose up --build

db:
	go run cmd/fake_db/main.go

server:
	go run cmd/fake_server/main.go

# make curl request to test
test:
	curl http://localhost:4444/redirect
