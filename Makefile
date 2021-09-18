build:
	go build -o ./bin/fibonacci ./cmd
start:
	sudo ./bin/fibonacci

docker-build:
	docker build -t fibonacci:latest -f Dockerfile .

docker-run:
	docker run --link memcached:mc -d -p 80:80 -p 8080:8080 --rm --name fibonacci fibonacci

docker-memcached:
	docker pull memcached
	docker run -d --rm --name memcached memcached