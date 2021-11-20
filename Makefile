build:
	go build -o ./bin/subnet-calculator cmd/subnet-calculator/main.go
run: build
	./bin/subnet-calculator