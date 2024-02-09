run:
		go run cmd/main.go
docker-run:
		docker run -d -p 27018:27017 --name mongodb_container mongo
docker-stop:
		docker stop mongodb_container
docker-rm:
		docker rm mongodb_container
	