package dto

type AppMode string

const (
	APPMODE_DEV  AppMode = "dev"
	APPMODE_TEST AppMode = "test"
	APPMODE_PROD AppMode = "prod"
)
