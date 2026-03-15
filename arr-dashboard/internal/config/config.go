package config

import (
	"encoding/json"
	"net"
	"os"
	"path/filepath"
)

type Service struct {
	Systemd string `json:"systemd"`
	Display string `json:"display"`
	Port    int    `json:"port"`
	Enabled bool   `json:"enabled"`
}

type Config struct {
	Host     string    `json:"host"`
	Services []Service `json:"services"`
}

var DefaultServices = []Service{
	{Systemd: "sonarr", Display: "Sonarr", Port: 8989, Enabled: true},
	{Systemd: "radarr", Display: "Radarr", Port: 7878, Enabled: true},
	{Systemd: "lidarr", Display: "Lidarr", Port: 8686, Enabled: true},
	{Systemd: "readarr", Display: "Readarr", Port: 8787, Enabled: true},
	{Systemd: "prowlarr", Display: "Prowlarr", Port: 9696, Enabled: true},
	{Systemd: "syncthing", Display: "Syncthing", Port: 8384, Enabled: true},
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

	for i := range cfg.Services {
		if !cfg.Services[i].Enabled {
			cfg.Services[i].Enabled = true
		}
	}

	return &cfg, nil
}

func Save(cfg *Config) error {
	path := ConfigPath()
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}

func UpdateServices(services []Service) error {
	cfg, err := Load()
	if err != nil {
		return err
	}
	cfg.Services = services
	return Save(cfg)
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

	return os.WriteFile(path, data, 0600)
}

func detectLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
