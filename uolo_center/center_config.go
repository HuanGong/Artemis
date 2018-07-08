package uolo_center

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
		LogLevel        string
		ServerAddress   string
		JWTSecretkey    string
		ProfilePhotoDir string
		MysqlConfig     MysqlConfig
		RedisConfig     []RedisConfig
	}
)
