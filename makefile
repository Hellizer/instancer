include .env
default: debug
	

build: clean
	@echo "building"
	-@mkdir build || echo "build folder exists"
	go build -o build/$(FNAME) cmd/main.go
	cp .env build/

run: 
	build/$(FNAME)

clean:
	@echo "clening"
	-@rm -rf build || echo "nothing to delete"

debug:
	-@go run cmd/main.go

migrate:
	migrate -source file://migration/ -database "postgresql://${DBU}:${DBP}@localhost:5432/${DBN}?sslmode=disable" -verbose up

migrate-drop:
	migrate -source file://migration/ -database "postgresql://${DBU}:${DBP}@localhost:5432/${DBN}?sslmode=disable" drop

.PHONY: build