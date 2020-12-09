.PHONY: day_1 day_2

run:
	@cd day_${day} && go run main.go

test:
	@go test -count=1 -race ./...
