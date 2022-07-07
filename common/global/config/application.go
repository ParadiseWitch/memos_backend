package config

type Application struct {
	Mode AppMode `json:"mode"`
	Host string  `json:"host"`
	Name string  `json:"name"`
	Port int     `json:"port"`
}

type AppMode string

const (
	APPMODE_DEV  AppMode = "dev"
	APPMODE_TEST AppMode = "test"
	APPMODE_PROD AppMode = "prod"
)
