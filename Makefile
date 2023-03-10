up:
	docker-compose -f dev/docker-compose.yml up
down:
	docker-compose -f dev/docker-compose.yml down
app:
	docker exec -it bench_api /bin/bash
run:
	go run ./cmd/main.go
bench_default:
	go-wrk -d 5 -T 10000 -M POST http://127.0.0.1:3000/default
bench_concurrent:
	go-wrk -d 5 -T 10000 -M POST http://127.0.0.1:3000/concurrent
