package configs

type Config struct {
	JWTSecret             string
	MSSQLConnectionString string
	DivineShieldBaseUrl   string
}

func InitConfig() *Config {
	return &Config{
		JWTSecret:             "02262025_secret",
		MSSQLConnectionString: "sqlserver://sa:Testing.1@localhost:1433?encrypt=disable&trustServerCertificate=true",
		DivineShieldBaseUrl:   "http://localhost:4000",
	}
}
