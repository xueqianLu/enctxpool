package config

var gconfig = new(Config)

type Config struct {
	// chain server
	ChainServerAddr string

	//
	EncTxpoolAddr string
}



func GetConfig() *Config {
	return gconfig
}