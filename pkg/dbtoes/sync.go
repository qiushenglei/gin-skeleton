package dbtoes

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"regexp"
	"strings"
)

type SyncMapping map[TopicName]Index

type TopicName string

// Index es中的index索引定义
type Index struct {
	PrimaryTable string                // 主表，例如订单表，有可能分表
	Tables       []string              // 主子表，例如订单表和用户表和订单商品表
	ForeignKey   string                // 主表和子表的外键
	ES           *elasticsearch.Client // es客户端
	sync         Synchronizer          // 必须实现的同步接口

	Type           string
	Table          string
	IsPrimaryTable bool
	Body           []byte
	BodyJson       map[string]interface{}
}

// Synchronizer 必须实现的同步接口
type Synchronizer interface {
	InsertSync() error
	UpdateSync() error
	FindPrimaryTable(index *Index) error
}

type Option func(*Index)

func NewIndex(op ...Option) *Index {
	index := &Index{}
	for _, optionFunc := range op {
		optionFunc(index)
	}
	return index
}

func WithPrimaryTable1(pt string) Option {
	return func(i *Index) {
		i.PrimaryTable = pt
	}
}

func WithTables(tables []string) Option {
	return func(i *Index) {
		i.Tables = tables
	}
}

func WithForeignKey1(ForeignKey string) Option {
	return func(i *Index) {
		i.ForeignKey = ForeignKey
	}
}

func WithEsClient(es *elasticsearch.Client) Option {
	return func(i *Index) {
		i.ES = es
	}
}

func WithSynchronizer(sync Synchronizer) Option {
	return func(i *Index) {
		i.sync = sync
	}
}

func (i *Index) Start(msg []byte) error {
	i.Body = msg
	i.ParseMessage()

	return nil
}

func (i *Index) ParseMessage() {
	if err := json.Unmarshal(i.Body, i.BodyJson); err != nil {
		panic(err)
	}

	// 判断是否是主表(分表)
	i.ParseIsPrimaryTable()

}

func (i *Index) ParseIsPrimaryTable() {
	i.IsPrimaryTable = isMatch(i.PrimaryTable, (i.BodyJson)["table"].(string))
}

func isMatch(tablePattern string, tableName string) bool {
	// 判断是否满足系统定义的分表表达式"tablePrefix_num"
	if hasSplitSymbol := strings.Contains(tablePattern, "_\\d"); hasSplitSymbol == true {

		// 查看表名和正则表达式是否匹配
		pattern := fmt.Sprintf("/^%s_\\d*$/", tableName)
		matched, err := regexp.Match(pattern, []byte(tableName))
		if err != nil {
			panic(err)
		}
		if matched != true {
			return false
		}
		return true
	} else if tablePattern == tableName {
		return true
	}

	return false
}

func (i *Index) update() {
	// TODO::判断哪些表需要变(有外键绑定的才能知道)

	// TODO::修改
	//i.sync.updateSync()
}

func (i *Index) insert() {
	// TODO::不是主表，是主表直接插入。是主表，查找主表文档是否存在，不存在阻塞等待主表插入(sleep，不能使用chan因为是分布式系统)
	i.sync.FindPrimaryTable(i)

	// TODO::插入
	i.sync.InsertSync()
}
