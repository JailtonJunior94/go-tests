test_bench: 
	cd internal/infra/database/ \
	go test -bench=.

test_fuzz: 
	cd internal/entities/ \
	go test -fuzz .

test:
    go test --coverprofile test/coverage.out ./...
	go tool cover -html=test/coverage.out