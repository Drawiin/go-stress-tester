run:
	go run main.go --url="http://localhost:3000/ac6bb669/unwrap" --requests=50000 --concurrency=8

build-img:
	docker build -f Dockerfile -t vini65599/go-stress-test:latest .

run-img:
	docker run -it --rm vini65599/go-stress-test:latest --url="http://google.com.br" --requests=50 --concurrency=8

.PHONY: build-img run-img