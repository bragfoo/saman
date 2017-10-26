package config

type ConfType struct {
	Api *ApiType
}

type ApiType struct {
	Enable  bool   `default: true`
	Prefix  string `default: "api"`
	Version string `default: "v1"`
	Port    int    `default: 8233`
}
