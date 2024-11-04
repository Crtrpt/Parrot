package host

import (
	"context"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	Api      *ApiCfg             `json:"api"`
	Crontab  map[string]*Crontab `json:"crontab"`
	Database *Database           `json:"database"`
	Redis    *RedisCfg           `json:"redis"`
}

func Start(ctx context.Context) (err error) {
	err = initConfig()
	if err != nil {
		return
	}
	err = initLog()
	if err != nil {
		return
	}

	err = initDB()
	if err != nil {
		return
	}

	err = initRedis()
	if err != nil {
		return
	}
	return
}

func initConfig() (err error) {
	Cfg = &Config{}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/parrot/")
	viper.AddConfigPath("$HOME/.parrot")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		return
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		err = viper.Unmarshal(&Cfg)
		if err != nil {
			logrus.Info(err.Error())
		}
	})
	return nil
}
