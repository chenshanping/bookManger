package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
}
