package dto

type DbDirver string

const (
	DBDRIVER_MYSQL    DbDirver = "mysql"
	DBDRIVER_POSTGRES DbDirver = "postgres"
	DBDRIVER_SQLITE   DbDirver = "sqlite"
)
