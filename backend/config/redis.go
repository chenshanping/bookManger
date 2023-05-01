package config

import "time"

type Redis struct {
	Host        string        `yaml:"host"`
	Port        int64         `yaml:"port"`
	Password    string        `yaml:"password"`
	DB          int           `yaml:"db"`
	Prefix      string        `yaml:"prefix"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
}
