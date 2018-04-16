package uolo_center

type (
	MysqlConfig struct {
		User     string
		Passwd   string
		HostPort string
		DBName   string
	}

	Config struct {
		LogLevel 		string
		ServerAddress 	string
		JWTSecretkey	string
		MysqlConfig		MysqlConfig
	}
)
