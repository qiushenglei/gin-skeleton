package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"

	"github.com/anguloc/zet/pkg/safe"
	"gorm.io/gen"
)

//go:generate go run .

var config string
var usage bool

type Querier interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id int) (gen.T, error)
}

func main() {
	// 获取flag
	flag.StringVar(&config, "config", safe.Path(".env.local"), "config path")
	flag.BoolVar(&usage, "help", false, "config path")
	flag.Parse()

	// help打印默认信息
	if usage {
		flag.PrintDefaults()
		return
	}
	v := viper.New()
	v.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
	v.SetConfigFile(config)
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	CreateCanalTestModels(v)
	CreateRWIsolateModels(v)

	return

	// canal_test
	//username := v.GetString("DB_SPEC_RW_USERNAME")
	//pw := v.GetString("DB_SPEC_RW_PASSWORD")
	//host := v.GetString("DB_SPEC_RW_HOST")
	//port := v.GetString("DB_SPEC_RW_PORT")
	//db := "canal_test"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=5s&readTimeout=5s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8", username, pw, host, port, db)
	//
	//conn, err := gorm.Open(mysql.Open(dsn))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//g.UseDB(conn)
	//
	//tables, err := conn.Migrator().GetTables()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//models := make([]interface{}, len(tables))
	//for i, v := range tables {
	//	models[i] = g.GenerateModel(v)
	//}
	//
	//g.ApplyBasic(models...)
	////g.ApplyInterface(func(Querier) {}, model.Rss{}, g.GenerateModel("score"))
	//
	//g.Execute()
}
