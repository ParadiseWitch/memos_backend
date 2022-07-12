package config

type Application struct {
	Mode AppMode `mapstructure:"mode"`
	Host string  `mapstructure:"host"`
	Name string  `mapstructure:"name"`
	Port int     `mapstructure:"port"`
}

type AppMode string

const (
	APPMODE_DEV  AppMode = "dev"
	APPMODE_TEST AppMode = "test"
	APPMODE_PROD AppMode = "prod"
)
