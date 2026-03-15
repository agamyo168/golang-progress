package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Service struct {
	Systemd string `json:"systemd"`
	Display string `json:"display"`
	Port    int    `json:"port"`
}

type Config struct {
	Host     string    `json:"host"`
	Services []Service `json:"services"`
}

var DefaultServices = []Service{
	{Systemd: "sonarr", Display: "Sonarr", Port: 8989},
	{Systemd: "radarr", Display: "Radarr", Port: 7878},
	{Systemd: "lidarr", Display: "Lidarr", Port: 8686},
	{Systemd: "readarr", Display: "Readarr", Port: 8787},
	{Systemd: "prowlarr", Display: "Prowlarr", Port: 9696},
}

var DefaultHost = "localhost"

func ConfigDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".arr-dashboard")
}

func ConfigPath() string {
	return filepath.Join(ConfigDir(), "config.json")
}

func GetHost() string {
	cfg, err := Load()
	if err != nil || cfg.Host == "" {
		return DefaultHost
	}
	return cfg.Host
}

func Load() (*Config, error) {
	path := ConfigPath()
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func EnsureConfig() error {
	path := ConfigPath()
	if _, err := os.Stat(path); err == nil {
		return nil
	}


	dir := ConfigDir()
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	host := detectLocalIP()
	if host == "" {
		host = DefaultHost
	}

	cfg := Config{
		Host:     host,
		Services: DefaultServices,
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func detectLocalIP() string {
	// Simple approach: try to connect to a known address and get local IP
	// This is a basic implementation - can be improved later
	return ""
}
