package config

type (
	Config struct {
		Database Database `yaml:"database"`
		Logger   Logger   `yaml:"logger"`
		Server   Server   `yaml:"server"`
	}
	Database struct {
		Mysql Mysql `yaml:"mysql"`
	}
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Mysql struct {
		Database  string `yaml:"database"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		Adabter   string `yaml:"adabter"`
		Time_zone string `yaml:"time_zone"`
		Charset   string `yaml:"charset"`
	}
	Logger struct {
		Max_Age          string `yaml:"max_age"`
		Max_Size         string `yaml:"max_size"`
		Filename_Pattern string `yaml:"filename_pattern"`
		Rotation_Time    string `yaml:"rotation_time"`
		Internal_Path    string `yaml:"internal_path"`
	}
)
