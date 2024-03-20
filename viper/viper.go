package viper

import (
	"flag"
	"fmt"
	"os"

	"github.com/championlong/go-backend-common/slog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type CliOptions interface {
	String() string
	GetConfigType() string
}

var viperGroup *viper.Viper

func InitConfig(path string, configValue CliOptions) error {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				panic("config路径为空，请设置")
			} else {
				config = configEnv
			}
		} else {
		}
	} else {
		config = path
	}
	slog.Infof("您正在使用func Viper()传递的值,config的路径为%v\n", config)

	viperGroup = viper.New()
	viperGroup.SetConfigFile(config)
	viperGroup.SetConfigType(configValue.GetConfigType())
	err := viperGroup.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viperGroup.WatchConfig()

	viperGroup.OnConfigChange(func(e fsnotify.Event) {
		if err := viperGroup.Unmarshal(&configValue); err != nil {
			slog.Errorf("config 动态序列号失败 %v", err)
		}
	})
	if err = viperGroup.Unmarshal(&configValue); err != nil {
		return err
	}
	return nil
}

func GetViper() *viper.Viper {
	return viperGroup
}
