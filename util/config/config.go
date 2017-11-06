package config

type ConfType struct {
	Api *ApiType
	Db  *DbType
	constant * ConstantType
}

type ApiType struct {
	Enable  bool   `default: true`
	Prefix  string `default: "api"`
	Version string `default: "v1"`
	Port    int    `default: 8233`
}

type DbType struct {
	Username string `default root`
	Password string `default 123..com`
	Protocol string `default tcp`
	Address  string `default 127.0.0.1`
	Port     string `defalut 3306`
	Dbname   string `default saman`
	MaxConnections int `default 100`
	MaxIdleConns int `default 100`
}

type ConstantType struct {
	IdGeneratorURL string `default http://127.0.0.1:1273/id`
} 