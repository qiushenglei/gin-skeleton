package main

import (
	"fmt"
	"github.com/anguloc/zet/pkg/safe"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

func CreateRWIsolateModels(v *viper.Viper) {
	// canal_test
	username := v.GetString("DB_CANAL_RW_ISOLATE_USERNAME")
	pw := v.GetString("DB_CANAL_RW_ISOLATE_PASSWORD")
	host := v.GetString("DB_SPEC_RW_HOST")
	port := v.GetString("DB_SPEC_RW_PORT")
	db := "rw_isolate"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=5s&readTimeout=5s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8", username, pw, host, port, db)

	conn, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建生成器
	c := gen.Config{
		OutPath:           safe.Path("/internal/app/data/mysql/rw_isolate/query"),
		ModelPkgPath:      safe.Path("/internal/app/data/mysql/rw_isolate/model"),
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode:              gen.WithDefaultQuery,
	}
	g := gen.NewGenerator(c)
	// 初始化生成器的db dial
	g.UseDB(conn)

	// 获取db表
	TableList, err := conn.Migrator().GetTables()
	if err != nil {
		panic(err)
	}

	// 生成model
	models := make([]interface{}, len(TableList))
	for i, v := range TableList {
		if strings.Contains(v, "order") {
			v = "order1"
		}
		models[i] = g.GenerateModel(v)
	}

	// 生成每个model的query
	g.ApplyBasic(models...)
	g.Execute()
}
