Name: service-exporter
Version: 1.0
Release: 1%{?dist}
Summary: Service Exporter for OS Metrics and Service Discovery

License: MIT
Source0: %{name}-%{version}.tar.gz

Requires: systemd

%description
Service Exporter for collecting OS metrics and self-discovering running services.

%prep
%setup -q

%build

%install
mkdir -p %{buildroot}/usr/local/bin
mkdir -p %{buildroot}/etc/service-exporter
cp service-exporter %{buildroot}/usr/local/bin/
cp config/config.yaml %{build
