package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Init 初始化viper
func Init() (err error) {
	viper.SetConfigName("config")
	// viper.SetConfigType("yaml") 配合远程配置中心使用的，告诉viper当前的配置使用什么格式来解析。（如何解析从远程来的字节流）
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了 ...")
	})
	return
}
