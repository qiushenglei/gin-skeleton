package conf

import "github.com/spf13/viper"

func InitConf(data interface{}, fileName string) error {
	// 初始化viper
	v := viper.New()

	// 设置file
	v.SetConfigFile(fileName)

	// 开始读取文件
	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	// 映射到结构体内
	err = v.Unmarshal(data)
	if err != nil {
		return err
	}
	return nil
}
