package statisticians


type (
	RedisConfig struct {
		Address	string
		Passwd	string
		DBIndex int32
	}
	Config struct {
		JWTSecretkey 	string //= "TestIv8^vow%Jf4gPfppyt*q"
		ServerAddress 	string
		LogLevel		string
		RedisConf 		*RedisConfig
	}
)