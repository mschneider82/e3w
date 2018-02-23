package conf

import (
	"gopkg.in/ini.v1"
        "os"
        "strings"
)

type Config struct {
	Port          string
	Auth          bool
	EtcdRootKey   string
	DirValue      string
	EtcdEndPoints []string
	EtcdUsername  string
	EtcdPassword  string
	CertFile      string
	KeyFile       string
	CAFile        string
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return defaultValue
    }
    return value
}

func getEnvStrList(key string, defaultValue string) []string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return strings.Split(defaultValue, ",")
    }
    return strings.Split(value, ",")
}

func Init(filepath string) (*Config, error) {
	cfg, err := ini.Load(filepath)
	if err != nil {
		return nil, err
	}

	c := &Config{}

        appSec := cfg.Section("app")
        c.Port = getEnv("APP_PORT", appSec.Key("port").Value())
        c.Auth = appSec.Key("auth").MustBool()

        etcdSec := cfg.Section("etcd")
        c.EtcdRootKey = getEnv("ETCD_ROOT_KEY", etcdSec.Key("root_key").Value())
        c.DirValue = etcdSec.Key("dir_value").Value()
        c.EtcdEndPoints = getEnvStrList("ETCD_ADDR", etcdSec.Key("addr").Value())
        c.EtcdUsername = getEnv("ETCD_USERNAME", etcdSec.Key("username").Value())
        c.EtcdPassword = getEnv("ETCD_PASSWORD", etcdSec.Key("password").Value())
        c.CertFile = etcdSec.Key("cert_file").Value()
        c.KeyFile = etcdSec.Key("key_file").Value()
        c.CAFile = etcdSec.Key("ca_file").Value()

	return c, nil
}
