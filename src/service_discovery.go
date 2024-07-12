package main

import (
    "bufio"
    "os/exec"
    "strings"
)

func discoverServices() {
    cmd := exec.Command("systemctl", "list-units", "--type=service", "--no-pager", "--all")
    output, err := cmd.Output()
    if err != nil {
        return
    }

    scanner := bufio.NewScanner(strings.NewReader(string(output)))
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, ".service") {
            fields := strings.Fields(line)
            service := fields[0]
            status := checkServiceStatus(service)
            serviceStatus.WithLabelValues(service).Set(status)
        }
    }
}

func checkServiceStatus(service string) float64 {
    cmd := exec.Command("systemctl", "is-active", service)
    output, err := cmd.Output()
    if err != nil {
        return 0
    }

    if strings.TrimSpace(string(output)) == "active" {
        return 1
    }
    return 0
}
