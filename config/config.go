package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DbConfig struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	DBname         string `yaml:"db-name"`
	ConnectionPool struct {
		MaxIdleConnections int   `yaml:"max-idle-connections"`
		MaxOpenConnections int   `yaml:"max-open-connections"`
		ConnectionLifetime int64 `yaml:"connection-lifetime"`
	} `yaml:"connection-pool"`
}

type LocalStorageConfig struct {
	MonitoredDirectory string `yaml:"monitored-directory"`
}

type SchedulerConfig struct {
	Period    int64 `yaml:"period"`
	BatchSize int   `yaml:"batch-size"`
	DaysDelta int   `yaml:"days-delta"`
}

type Config struct {
	ServerConfig       ServerConfig       `yaml:"server"`
	DbConfig           DbConfig           `yaml:"db"`
	LocalStorageConfig LocalStorageConfig `yaml:"local-storage"`
	SchedulerConfig    SchedulerConfig    `yaml:"scheduler"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
