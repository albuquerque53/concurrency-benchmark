# Concurrency Benchmark :bicyclist:

## :mag: About the project

This is an simple API that searches mocked users and send them to an external API.<br>
> Each API call lasts 3 seconds.

**There are two endpoints:**
1. `/default`: Will send the users to external API in a procedural way.
2. `/concurrent`: Will send the users to external API using concurrency.

I've added some shortcuts in Makefile, so that you can run benchmarks commands simply.
> btw, I'm using the [go-wrk](https://github.com/tsliwowicz/go-wrk) to run benchmarks the endpoints.

## :runner: How to run it?

1. Build container:
```
make start
```

2. Get into app & run API:
```
make app && make run
```

## :watch: Benchmark

1. To benchmark the "default" route:
```
make bench_default
```

2. To benchmark the "concurrent" route:
```
make bench_concurrent
```
