#!/bin/bash

# Create configuration directory
mkdir -p /etc/service-exporter

# Copy configuration file
cp config/config.yaml /etc/service-exporter/config.yaml

# Copy binary to /usr/local/bin
cp build/service-exporter /usr/local/bin/service-exporter

# Copy service file to systemd
cp scripts/service-exporter.service /etc/systemd/system/

# Reload systemd and enable the service
systemctl daemon-reload
systemctl enable service-exporter
systemctl start service-exporter
