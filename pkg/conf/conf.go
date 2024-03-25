package conf

import "github.com/spf13/viper"

func InitConf(data interface{}, fileName string) {
	// 初始化viper
	v := viper.New()

	// 设置file
	v.SetConfigFile(fileName)

	// 开始读取文件
	err := v.ReadInConfig()
	if err != nil {
		panic("read config err:" + err.Error())
	}

	// 映射到结构体内
	v.Unmarshal(data)
}
