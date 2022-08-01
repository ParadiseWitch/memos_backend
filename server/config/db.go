package config

type Db struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type DbDirver string

const (
	DBDRIVER_MYSQL    DbDirver = "mysql"
	DBDRIVER_POSTGRES DbDirver = "postgres"
	DBDRIVER_SQLITE   DbDirver = "sqlite"
)
