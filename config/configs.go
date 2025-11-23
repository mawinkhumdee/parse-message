package config

type Config struct {
	DB DB `yaml:"db"`
}
type DB struct {
	URI          string `yaml:"uri"`
	DatabaseName string `yaml:"database-name"`
}
