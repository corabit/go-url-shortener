dependencies:
	go mod download

run: dependencies
	go run main.go

build: dependencies
	go build main.go

build-docker: 
	docker build -t go-url-shortener .

run-docker: 
	docker run --name url-shortener  -p 8080:8080 -d koralbit:url-shortener

stop-docker:
	docker container stop go-url-shortener || true

remove-docker-container:
	docker container rm go-url-shortener || true

remove-docker-image:
	docker image rm go-url-shortener || true

clean: stop-docker remove-docker-container remove-docker-image