package config

import "time"

type (
	Config struct {
		Logger   Logger   `yaml:"logger"`
		Auth     Auth     `yaml:"auth"`
		Database Database `yaml:"database"`
		Grpc     Grpc     `yaml:"grpc"`
	}
	Database struct {
		Postgresql Postgresql `yaml:"postgresql"`
	}
	Grpc struct {
		Port string `yaml:"port"`
	}
	Logger struct {
		Max_Age          string `yaml:"max_age"`
		Max_Size         string `yaml:"max_size"`
		Filename_Pattern string `yaml:"filename_pattern"`
		Rotation_Time    string `yaml:"rotation_time"`
		Internal_Path    string `yaml:"internal_path"`
		Mode             string `yaml:"mode"`
	}
	Auth struct {
		Secretkey    string        `yaml:"secretkey"`
		Issue        string        `yaml:"issue"`
		ExpireTime_a time.Duration `yaml:"expire_time_a"`
		ExpireTime_r time.Duration `yaml:"expire_time_r"`
	}
	Postgresql struct {
		Database  string `yaml:"database"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		Adabter   string `yaml:"adabter"`
		Time_zone string `yaml:"time_zone"`
		Charset   string `yaml:"charset"`
	}
)
