package configs

import (
	"github.com/spf13/viper"
)

var (
	//系统变量
	BasePath   string // 定义项目的根目录
	LogPath    string // 定义项目的日志目录
	EnvFile    string // 定义配置文件名称
	HttpPort   string // 服务端口
	AppRunMode string // 程序运行模式

	// EnvConfig ... 全局配置文件
	EnvConfig *viper.Viper
)
