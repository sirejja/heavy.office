
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

high-load-checkout:
	ghz --insecure --async \
--proto ./checkout/api/v1/checkout_v1.proto \
--call checkout.Checkout/ListCart \
-c 10 -n 1000000 --rps 200 \
-d '{"user": 8}' 0.0.0.0:8080

error-high-load-checkout:
	ghz --insecure --async \
--proto ./checkout/api/v1/checkout_v1.proto \
--call checkout.Checkout/ListCart \
-c 10 -n 1000000 --rps 200 \
-d '{"user": 8}' 0.0.0.0:8080

high-load-loms:
	ghz --insecure --async \
--proto ./loms/api/v1/loms_v1.proto \
--call loms.Loms/Stocks \
-c 10 -n 10000 \
--load-schedule=line --load-start=5 --load-step=5 \
-d '{"sku": 4487693}' 0.0.0.0:8081