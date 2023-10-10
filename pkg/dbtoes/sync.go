package dbtoes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type SyncMapping map[TopicName]Index

type TopicName string

// Index es中的index索引定义
type Index struct {
	PrimaryTable string                     // 主表，例如订单表，有可能分表
	Tables       []string                   // 主子表，例如订单表和用户表和订单商品表
	ForeignKey   string                     // 主表和子表的外键
	ESConn       *elasticsearch.Client      // es客户端
	TypedESConn  *elasticsearch.TypedClient // es客户端
	sync         Synchronizer               // 必须实现的同步接口

	SQLType        string
	SyncTable      string
	IsPrimaryTable bool
	Body           []byte
	BodyJson       map[string]interface{}
	BodyData       []interface{}
	BodyFirstData  map[string]interface{}
	PrimaryID      string
	PrimarySource  interface{}
}

func (i *Index) SetPrimarySource(PrimarySource interface{}) {
	i.PrimarySource = PrimarySource
}

func (i *Index) SetPrimaryID(PrimaryID string) {
	i.PrimaryID = PrimaryID
}

// Synchronizer 必须实现的同步接口
type Synchronizer interface {
	InsertSync(i *Index) error
	UpdateSync(i *Index) error
	FindPrimaryTableByPForeignKey(i *Index) error
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

func WithForeignKey(ForeignKey string) Option {
	return func(i *Index) {
		i.ForeignKey = ForeignKey
	}
}

func WithEsClient(es *elasticsearch.Client) Option {
	return func(i *Index) {
		i.ESConn = es
	}
}

func WithEsTypedClient(es *elasticsearch.TypedClient) Option {
	return func(i *Index) {
		i.TypedESConn = es
	}
}

func WithSynchronizer(sync Synchronizer) Option {
	return func(i *Index) {
		i.sync = sync
	}
}

func (i *Index) Start(msg []byte) error {
	i.Body = msg

	// 解析msg
	i.parseMessage()

	// 执行同步
	i.proc()

	return nil
}

func (i *Index) parseMessage() {
	if err := json.Unmarshal(i.Body, &i.BodyJson); err != nil {
		panic(err)
	}

	// 同步的表命
	i.parseSyncTable()

	// SQL类型
	i.parseSQLType()

	i.parseBodyData()

	// 判断是否是主表(分表)
	i.parseIsPrimaryTable()
}

func (i *Index) parseSyncTable() {
	i.SyncTable = (i.BodyJson)["table"].(string)
}

func (i *Index) parseSQLType() {
	i.SQLType = (i.BodyJson)["type"].(string)
}

func (i *Index) parseBodyData() {
	i.BodyData = (i.BodyJson)["data"].([]interface{})
	i.BodyFirstData = i.BodyData[0].(map[string]interface{})
}

func (i *Index) parseIsPrimaryTable() {
	i.IsPrimaryTable = isMatch(i.PrimaryTable, i.SyncTable)
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

func (i *Index) proc() {
	// 首字符大写
	name := cases.Title(language.Und).String(strings.ToLower(i.SQLType))

	// 根据type类型，执行操作(update or insert)
	v := reflect.ValueOf(i).MethodByName(name)
	if v.IsValid() && v.Kind() == reflect.Func {
		param := []reflect.Value{
			//reflect.ValueOf(111),
		}
		v.Call(param)
	} else {
		logs.Log.Info(context.Background(), "no proc")
	}
}

func (i *Index) Update() {
	// es查询主表内容,如果没有主表document，挂起goroutine,等待主表插入后执行(sleep，不能使用chan因为是分布式系统)
	for {
		// 查主表信息，找到es的doc_id,用于后续更新doc。没有查到主表信息，则loop等待其他协程同步主表信息到es
		if err := i.sync.FindPrimaryTableByPForeignKey(i); err != nil {
			//
			if v, ok := err.(errorpkg.Errx); ok && v.Code() == CodeSyncNoLoop {
				return
			}
			time.Sleep(2 * time.Second)
			continue
		} else {
			break
		}
	}

	i.sync.UpdateSync(i)
}

func (i *Index) Insert() {
	// es查询主表内容,如果没有主表document，挂起goroutine,等待主表插入后执行(sleep，不能使用chan因为是分布式系统)
	for {
		// 是主表insert，直接同步到es
		if i.IsPrimaryTable {
			break
		}

		// 非主表，查主表信息，找到es的doc_id,用于后续更新doc
		// 没有查到主表信息，则loop等待其他协程同步主表信息到es
		if err := i.sync.FindPrimaryTableByPForeignKey(i); err != nil {
			//
			if v, ok := err.(errorpkg.Errx); ok && v.Code() == CodeSyncNoLoop {
				return
			}
			time.Sleep(2 * time.Second)
			continue
		} else {
			break
		}
	}

	i.sync.InsertSync(i)
}
