package configs

type Config struct {
	JWTSecret             string
	MSSQLConnectionString string
	DivineShieldBaseUrl   string
	SlackToken            string
}

func InitConfig() *Config {
	return &Config{
		JWTSecret:             "02262025_secret",
		MSSQLConnectionString: "sqlserver://sa:Testing.1@localhost:1433?encrypt=disable&trustServerCertificate=true",
		DivineShieldBaseUrl:   "http://localhost:4000",
		SlackToken:            "xoxb-8563173054387-8562932784514-9mpXhTY7I2QMsAIOiybDMPe4",
	}
}
