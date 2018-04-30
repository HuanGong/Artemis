package uolo_lens

type (
	MysqlConfig struct {
		User     string
		Passwd   string
		HostPort string
		DBName   string
	}
	Config struct {
		ServerAddress string
		JWTSecretkey 	string
		LogLevel		string
		MysqlConfig
	}
)
