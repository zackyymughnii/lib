package db

import (
	"fmt"
	"net/url"
)

type Config struct {
	Driver       string `yaml:"driver"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Name         string `yaml:"name"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Locale       string `yaml:"locale"`
	MaxOpenConns int    `yaml:"maxopenconns"`
}

func (d Config) String() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
		url.QueryEscape(d.Locale),
	)
}
