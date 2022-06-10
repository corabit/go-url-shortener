run: build-docker
	docker-compose -p url-shortener up -d

build-docker: 
	docker build -t koralbit/go-url-shortener .

stop:
	docker-compose -p url-shortener stop

clean:
	docker-compose -p url-shortener down