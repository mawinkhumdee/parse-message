package config

type Config struct {
	DB     DB     `yaml:"db"`
	Gemini Gemini `yaml:"gemini"`
	Prompt Prompt `yaml:"prompt"`
	Kafka  Kafka  `yaml:"kafka"`
}

type DB struct {
	URI          string `yaml:"uri"`
	DatabaseName string `yaml:"database-name"`
}

type Gemini struct {
	ApiKey string `yaml:"api-key"`
	Model  string `yaml:"model"`
}

type Prompt struct {
	ExtractFields string `yaml:"extract-fields"`
}

type Kafka struct {
	Brokers []string               `yaml:"brokers"`
	Topics  map[string]TopicConfig `yaml:"topics"`
}

type TopicConfig struct {
	Topic   string `yaml:"topic"`
	GroupID string `yaml:"group-id"`
}
