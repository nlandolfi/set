all:
	make check && make cover && make bench

# bench runs the benchmarks
bench:
	go test --bench=.

# check runs the tests
check:
	go test

# coverage runs the tests and generates coverage info
cover:
	go test --cover
