test_bench: 
	cd internal/infra/database/ \
	go test -bench=.

test_fuzz: 
	cd internal/entities/ \
	go test -fuzz .