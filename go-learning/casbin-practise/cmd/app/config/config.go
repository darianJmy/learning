package config

type Config struct {
	Default DefaultOptions `yaml:"default"`
	Mysql   MysqlOptions   `yaml:"mysql"`
}

type DefaultOptions struct {
	Listen int    `yaml:"listen"`
	LogDir string `yaml:"log_dir"`
}

type MysqlOptions struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

