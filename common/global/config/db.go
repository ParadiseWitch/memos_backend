package config

type Db struct {
	Driver   DbDirver `json:"driver"`
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type DbDirver string

const (
	DBDRIVER_MYSQL    DbDirver = "mysql"
	DBDRIVER_POSTGRES DbDirver = "postgres"
	DBDRIVER_SQLITE   DbDirver = "sqlite"
)
