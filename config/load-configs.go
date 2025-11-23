package config

func LoadConfig() Config {
	return Config{
		DB: DB{
			URI:          "mongodb://vector:secretvector@localhost:27017/?authSource=admin&directConnection=true",
			DatabaseName: "vectors",
		},
	}
}
