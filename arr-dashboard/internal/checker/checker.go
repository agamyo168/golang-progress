package checker

import (
	"arr-dashboard/internal/config"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type ServiceStatus struct {
	Name    string `json:"name"`
	Systemd string `json:"systemd"`
	Port    int    `json:"port"`
	URL     string `json:"url"`
	Status  string `json:"status"`
}

type ServicesResponse struct {
	Services    []ServiceStatus `json:"services"`
	LastRefresh time.Time        `json:"lastRefresh"`
}

func GetServices() (*ServicesResponse, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	host := config.GetHost()
	services := make([]ServiceStatus, len(cfg.Services))

	for i, svc := range cfg.Services {
		status := checkSystemd(svc.Systemd)

		services[i] = ServiceStatus{
			Name:    svc.Display,
			Systemd: svc.Systemd,
			Port:    svc.Port,
			URL:     fmt.Sprintf("http://%s:%d", host, svc.Port),
			Status:  status,
		}
	}

	return &ServicesResponse{
		Services:    services,
		LastRefresh: time.Now(),
	}, nil
}

func checkSystemd(serviceName string) string {
	cmd := exec.Command("systemctl", "is-active", serviceName)
	output, err := cmd.Output()

	if err != nil {
		return "inactive"
	}

	result := string(output)
	if result == "active\n" {
		return "active"
	}
	return "inactive"
}

func Refresh() (*ServicesResponse, error) {
	return GetServices()
}

func MarshalJSON(resp *ServicesResponse) ([]byte, error) {
	return json.Marshal(resp)
}
