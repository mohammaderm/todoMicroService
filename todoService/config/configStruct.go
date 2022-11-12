package config

type (
	Config struct {
		Logger   Logger   `yaml:"logger"`
		Database Database `yaml:"database"`
		Grpc     Grpc     `yaml:"grpc"`
		Redis    Redis    `yaml:"redis"`
	}
	Database struct {
		Mysql Mysql `yaml:"mysql"`
	}
	Grpc struct {
		Port string `yaml:"port"`
	}
	Mysql struct {
		Database  string `yaml:"database"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		Adabter   string `yaml:"adabter"`
		Time_zone string `yaml:"time_zone"`
		Charset   string `yaml:"charset"`
	}
	Redis struct {
		Server   string `yaml:"server"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
		Port     string `yaml:"port"`
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
