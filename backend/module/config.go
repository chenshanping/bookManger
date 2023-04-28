package module

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	System System `yaml:"system"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type System struct {
	Env  string `yaml:"env"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
