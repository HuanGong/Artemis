package mem_notifier

type (
	MysqlConfig struct {
		User     string
		Passwd   string
		HostPort string
		DBName   string
	}

	Config struct {
		MysqlConfig		MysqlConfig
		ServerAddress 	string
	}
)
