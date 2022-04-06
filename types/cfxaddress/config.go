package cfxaddress

type Config struct {
	AddressStringVerbose bool
}

func SetConfig(c Config) {
	config = c
}
func GetConfig() Config {
	return config
}

var config Config
