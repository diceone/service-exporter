package main

import (
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/process"
    "github.com/prometheus/client_golang/prometheus"
)

var (
    cpuUsage = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "os_cpu_usage",
            Help: "CPU usage percentage.",
        },
        []string{"cpu"},
    )
    memUsage = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "os_memory_usage",
            Help: "Memory usage percentage.",
        },
    )
    diskUsage = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "os_disk_usage",
            Help: "Disk usage percentage.",
        },
        []string{"path"},
    )
    serviceStatus = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "service_status",
            Help: "Status of system services.",
        },
        []string{"service"},
    )
    processCount = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "process_count",
            Help: "Number of running processes.",
        },
    )
    processCPU = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "process_cpu_usage",
            Help: "CPU usage percentage of each process.",
        },
        []string{"pid", "name"},
    )
    processMemory = prometheus.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "process_memory_usage",
            Help: "Memory usage percentage of each process.",
        },
        []string{"pid", "name"},
    )
)

func init() {
    prometheus.MustRegister(cpuUsage)
    prometheus.MustRegister(memUsage)
    prometheus.MustRegister(diskUsage)
    prometheus.MustRegister(serviceStatus)
    prometheus.MustRegister(processCount)
    prometheus.MustRegister(processCPU)
    prometheus.MustRegister(processMemory)
}

func collectMetrics() {
    cpus, _ := cpu.Percent(0, true)
    for i, c := range cpus {
        cpuUsage.WithLabelValues(string(i)).Set(c)
    }

    vm, _ := mem.VirtualMemory()
    memUsage.Set(vm.UsedPercent)

    partitions, _ := disk.Partitions(true)
    for _, p := range partitions {
        usage, _ := disk.Usage(p.Mountpoint)
        diskUsage.WithLabelValues(p.Mountpoint).Set(usage.UsedPercent)
    }

    procs, _ := process.Processes()
    processCount.Set(float64(len(procs)))

    for _, proc := range procs {
        pid := proc.Pid
        name, _ := proc.Name()
        cpuPercent, _ := proc.CPUPercent()
        memPercent, _ := proc.MemoryPercent()
        
        processCPU.WithLabelValues(string(pid), name).Set(cpuPercent)
        processMemory.WithLabelValues(string(pid), name).Set(float64(memPercent))
    }
}
