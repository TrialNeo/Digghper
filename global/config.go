package global

type Config struct {
	Web struct {
		Port int
	}
	Database struct {
		DriverName     string
		DataSourceName string
		Host           string
		Port           int
		User           string
		Psw            string
		TimeZone       string
	}
	Redis struct {
		Addr     string
		Password string
	}
}
