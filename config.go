package emailer

import "log"

type Config struct {
	Host     string
	Port     int
	address  string
	Password string
	Email    string
	Name     string
}

func CreateConfig(host string, port int, address, password, email string, name string) *Config {
	if name == "" {
		log.Print("emailer has no name")
		return nil
	}
	return &Config{
		Host:     host,
		Port:     port,
		address:  address,
		Password: password,
		Email:    email,
		Name:     name,
	}
}
