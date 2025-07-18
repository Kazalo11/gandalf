.PHONY: run-backend
run-backend:
	go run cmd/app/main.go

.PHONY: run-backend-test
run-backend-test:
	go test -v ./... -coverprofile=coverage.out