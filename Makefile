dev:
	go run cmd/main.go example $(ARGS)

exec:
	go run cmd/main.go real $(ARGS)