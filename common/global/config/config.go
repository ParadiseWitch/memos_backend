package config

type Config struct {
	Application Application `json:"application"`
	Db          Db          `json:"db"`
	Log         Log         `json:"log"`
}
