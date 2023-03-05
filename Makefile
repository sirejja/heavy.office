
build-all:
	 cd checkout && GOOS=linux make build
	 cd loms && GOOS=linux make build
	 cd notifications && GOOS=linux make build

run-all: build-all
	sudo docker compose up --force-recreate --build

precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit

prepare-grpc-deps:
	cd checkout && make prepare-grpc-deps
	cd loms && make prepare-grpc-deps