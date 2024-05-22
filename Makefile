install-all:
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

generate-assets-storage:
	mkdir -p internal/assets-storage/generated/schema
	oapi-codegen -package dto -generate chi-server,types,spec api/assets-storage/schema.yaml > internal/assets-storage/generated/schema/dto.gen.go

generate-spreads-storage:
	mkdir -p internal/spreads-storage/generated/schema
	oapi-codegen -package dto -generate chi-server,types,spec api/spreads-storage/schema.yaml > internal/spreads-storage/generated/schema/dto.gen.go