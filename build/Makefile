all: build

build:
    go build -o build/service-exporter src/main.go src/metrics.go src/service_discovery.go

clean:
    rm -rf build/service-exporter

.PHONY: all build clean
