run:
	@cd day_${day} && go run main.go

test:
	@go test -count=1 -race ./...
