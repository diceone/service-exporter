package main

import (
    "net/http"
    "time"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

type Config struct {
    Interval         string   `yaml:"interval"`
    MetricsEndpoint  string   `yaml:"metrics_endpoint"`
    Services         []string `yaml:"services"`
}

var config Config

func loadConfig() {
    data, err := ioutil.ReadFile("/etc/service-exporter/config.yaml")
    if err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        log.Fatalf("Error parsing config file: %v", err)
    }
}

func main() {
    loadConfig()
    http.Handle(config.MetricsEndpoint, promhttp.Handler())

    go func() {
        interval, _ := time.ParseDuration(config.Interval)
        for {
            collectMetrics()
            time.Sleep(interval)
        }
    }()

    go func() {
        for {
            discoverServices()
            time.Sleep(60 * time.Second)
        }
    }()

    http.ListenAndServe(":5555", nil)
}
