package uolo_lens

type (
	MysqlConfig struct {
		User     string
		Passwd   string
		HostPort string
		DBName   string
	}
	RedisConfig struct {
		Name     string
		Addr     string
		DbIndex  int32
		Password string
	}
	Config struct {
		ServerAddress string
		JWTSecretkey  string
		LogLevel      string
		MysqlConfig   MysqlConfig
		PostDataDir   string
		RedisConfig   []RedisConfig
		CleanMarkPath string
	}
)
