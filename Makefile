.DEFAULT_GOAL := buildnrun

BIN_FILE=iban_validator
BIN_FILE_MAC=iban_validator_darwin

build:
	go build -o "$(BIN_FILE)"

test:
	go test ./... -cover

run:
	./"$(BIN_FILE)"

clean:
	go clean
	rm --force "$(BIN_FILE)" "$(BIN_FILE_MAC)"

buildnrun:
	go run main.go

buildmac:
	GOOS=darwin GOARCH=amd64 go build -o "$(BIN_FILE_MAC)"

runmac:
	chmod +x ./"$(BIN_FILE_MAC)"
	./"$(BIN_FILE_MAC)"