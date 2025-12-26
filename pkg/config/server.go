package config

type ServerConfig struct {
	Listen     string `yaml:"listen"`
	JWTSecrete string `yaml:"jwt_secrete"`
}
