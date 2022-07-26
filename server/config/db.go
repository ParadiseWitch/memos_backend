package config

type Db struct {
	Driver DbDirver `mapstructure:"driver"`
	Url    string   `mapstructure:"url"`
}

type DbDirver string

const (
	DBDRIVER_MYSQL    DbDirver = "mysql"
	DBDRIVER_POSTGRES DbDirver = "postgres"
	DBDRIVER_SQLITE   DbDirver = "sqlite"
)
