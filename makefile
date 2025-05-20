.PHONY: clean build run-client run-server run-multiple

clean:  
	docker-compose down --volumes --remove-orphans
	docker system prune -f

build:
	docker-compose up --build

create:
	docker-compose up -d

run:
	docker start grpc_server

run-client:
	docker start grpc_client

# only run grpc client
# run multiple client here 3 instance

run-multiple:
	docker-compose up --scale client=3

run-server:
	docker run --rm --network grpc_default -it grpc_server

run-client:
	docker run --rm --network grpc_default -it grpc_client

# .DEFAULT_GOAL := clean build 