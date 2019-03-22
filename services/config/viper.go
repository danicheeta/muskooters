package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init(appName string) {
	viper.SetEnvPrefix(appName)
	viper.AutomaticEnv()
}

func MustString(s string) string {
	confVal := viper.GetString(s)
	if confVal == "" {
		logrus.Panicf("env var %s cannot be empty", s)
	}

	return confVal
}
