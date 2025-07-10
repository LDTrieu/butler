run:
	go run cmd/main.go

docker-build:
	rm -rf build
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/main cmd/main.go
	docker build -t butler . --no-cache

docker-run:
	docker run -p 5000:5000 --name butler --detach butler