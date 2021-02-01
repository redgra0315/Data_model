/*
2 @Auth: Linux
3 @Date: 2021/1/27 16:00
4 */
package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Address string `mapstructure:"address"`
	Service string `mapstructure:"service"`
	Token   string `mapstructure:"token"`
	Secret  string `mapstructrre:"secret"`
	Path    string `mapstructrre:"path"`
	Phone   int8   `mapstructrre:"phone"`
	Port    int    `mapstructure:"port"`
}

var Conf = new(Config)

//  检测配置文件的路径
func CheckPath() (err error, config interface{}) {

	viper.SetConfigFile("config/config.yaml") //
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return err, nil
	}

	// 将读取的配置信息保存至全局变量Conf
	if err = viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		return
	}
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})
	//return Conf, nil
	//fmt.Println(Conf.Address,Conf.Port,Conf.Service,Conf.Token)
	return
}
