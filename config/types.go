package config

type Config struct {
	Bind string `yaml:"bind"`
	PingTime int `yaml:"ping_time"`
	Timeout int `yaml:"timeout"`
	Servers map[string]string `yaml:"servers"`
}

func (c *Config) NewDefault() {
	c.Bind = "0.0.0.0:8080"
	c.PingTime = 30
	c.Timeout = 10
	c.Servers = map[string]string{
		"test1": "127.0.0.1:25565",
		"test2": "127.0.0.1:25566",
	}
}