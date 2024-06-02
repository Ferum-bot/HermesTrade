install-all:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

generate-assets-storage:
	mkdir -p internal/assets-storage/generated/schema
	oapi-codegen -package dto -generate chi-server,types,spec api/assets-storage/schema.yaml > internal/assets-storage/generated/schema/dto.gen.go

generate-spreads-storage:
	mkdir -p internal/spreads-storage/generated/schema
	oapi-codegen -package dto -generate chi-server,types,spec api/spreads-storage/schema.yaml > internal/spreads-storage/generated/schema/dto.gen.go

launch-dev-env:
	sh ./dev/mongo/start.sh
	sh ./dev/kafka/start.sh
	sh ./dev/postgres/start.sh

stop-dev-env:
	sh ./dev/mongo/stop.sh
	sh ./dev/kafka/stop.sh
	sh ./dev/postgres/stop.sh

test:
	 go test ./...
lint:
	golangci-lint run

build-all-docker-images:
	docker build -t hermes-trade-asset-spread-hunter:latest -f ./docker/asset-spread-hunter/Dockerfile .

	docker build -t hermes-trade-assets-storage:latest -f ./docker/assets-storage/Dockerfile .

	docker build -t hermes-trade-spreads-storage:latest -f ./docker/spreads-storage/Dockerfile .

	docker build -t hermes-trade-connector-telegram:latest -f ./docker/connector/telegram/Dockerfile .

	docker build -t hermes-trade-scrapper-binance:latest -f ./docker/scrapper/binance/Dockerfile .
	docker build -t hermes-trade-scrapper-by-bit:latest -f ./docker/scrapper/by-bit/Dockerfile .
	docker build -t hermes-trade-scrapper-coinbase:latest -f ./docker/scrapper/coinbase/Dockerfile .
	docker build -t hermes-trade-scrapper-kraken:latest -f ./docker/scrapper/kraken/Dockerfile .
	docker build -t hermes-trade-scrapper-okx:latest -f ./docker/scrapper/okx/Dockerfile .
	docker build -t hermes-trade-scrapper-upbit:latest -f ./docker/scrapper/upbit/Dockerfile .

launch-asset-spread-hunter:
	go run ./cmd/asset-spread-hunter/main.go

launch-assets-storage:
	go run ./cmd/assets-storage/main.go

launch-spreads-storage:
	go run ./cmd/spreads-storage/main.go

launch-connector-telegram:
	go run ./cmd/connector/telegram/main.go

launch-all-scrappers:
	go run ./cmd/scrapper/binance/main.go
	go run ./cmd/scrapper/by-bit/main.go
	go run ./cmd/scrapper/coinbase/main.go
	go run ./cmd/scrapper/kraken/main.go
	go run ./cmd/scrapper/okx/main.go
	go run ./cmd/scrapper/upbit/main.go
