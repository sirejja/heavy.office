
build-all:
	 cd checkout && GOOS=linux make build
	 cd loms && GOOS=linux make build
	 cd notifications && GOOS=linux make build

run-all: build-all
	sudo docker compose -f docker-compose.yml up --force-recreate --build

precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit

prepare-grpc-deps:
	cd checkout && make prepare-grpc-deps
	cd loms && make prepare-grpc-deps

run-kafka:
	sudo docker compose -f docker-compose-kafka.yml up --build

logs:
	mkdir -p logs/data
	touch logs/data/log.txt
	touch logs/data/offsets.yaml
	sudo chmod -R 777 logs/data
	cd graylog && sudo docker compose up
