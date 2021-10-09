package emailer

type Config struct {
	Host string
	Port int
	address string
	Password string
	Email string
}

func CreateConfig(host string, port int, address, password, email string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		address:  address,
		Password: password,
		Email: email,
	}
}
