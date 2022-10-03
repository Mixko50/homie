package config

type config struct {
	Address     string   `yaml:"address"`
	AutoMigrate bool     `yaml:"auto_migrate"`
	Cors        []string `yaml:"cors"`
	JwtSecret   string   `yaml:"jwt_secret"`
	MySqlConfig string   `yaml:"mysql_config"`
}
