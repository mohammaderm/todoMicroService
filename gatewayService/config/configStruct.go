package config

import "time"

type (
	Config struct {
		Logger  Logger  `yaml:"logger"`
		Service Service `yaml:"service"`
		Server  Server  `yaml:"server"`
		Token   Token   `yaml:"token"`
		Metrics Metrics `yaml:"metrics"`
	}
	Metrics struct {
		Port string `yaml:"port"`
	}
	Token struct {
		Secretkey string `yaml:"secretkey"`
		Issue     string `yaml:"issue"`
	}
	Server struct {
		Port                   string        `yaml:"port"`
		GracefulShutdownPeriod time.Duration `yaml:"gracefulShutdownPeriod"`
	}
	Auth struct {
		Host            string        `yaml:"host"`
		Port            string        `yaml:"port"`
		ContextDeadline time.Duration `yaml:"contextDeadline"`
	}
	Todo struct {
		Host            string        `yaml:"host"`
		Port            string        `yaml:"port"`
		ContextDeadline time.Duration `yaml:"contextDeadline"`
	}
	Service struct {
		Auth Auth `yaml:"auth"`
		Todo Todo `yaml:"todo"`
	}
	Logger struct {
		Max_Age          string `yaml:"max_age"`
		Max_Size         string `yaml:"max_size"`
		Filename_Pattern string `yaml:"filename_pattern"`
		Rotation_Time    string `yaml:"rotation_time"`
		Internal_Path    string `yaml:"internal_path"`
		Mode             string `yaml:"mode"`
	}
)
