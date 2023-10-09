package example

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/qiushenglei/gin-skeleton/pkg/dbtoes"
	"net/http"
	"time"
)

func AATest() {

	addr := "http://127.0.0.1:9200"
	cfg := elasticsearch.Config{
		Addresses: []string{addr},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 100 * time.Millisecond,
			//DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
		},
	}

	ESClient := dbtoes.NewESClient(cfg)
	ESTypedClient := dbtoes.NewESTypedClient1(cfg)

	msg := []byte(`{
	 "data": [
	   {
	     "id": "2",
	     "username": "ld",
	     "label": "2,3",
	     "class_id": "1",
	     "student_id": "3",
	     "add_time": "2023-10-05 13:21:09",
	     "update_time": "2023-10-05 13:21:11",
	     "is_deleted": "0"
	   }
	 ],
	 "database": "canal_test",
	 "es": 1696483273000,
	 "id": 4,
	 "isDdl": false,
	 "mysqlType": {
	   "id": "int unsigned",
	   "username": "varchar(255)",
	   "label": "varchar(255)",
	   "class_id": "int(10) unsigned",
	   "student_id": "int(10) unsigned",
	   "add_time": "datetime(0)",
	   "update_time": "datetime(0)",
	   "is_deleted": "tinyint(4)"
	 },
	 "old": null,
	 "pkNames": [
	   "id"
	 ],
	 "sql": "",
	 "sqlType": {
	   "id": 4,
	   "username": 12,
	   "label": 12,
	   "class_id": 4,
	   "student_id": 4,
	   "add_time": 93,
	   "update_time": 93,
	   "is_deleted": -6
	 },
	 "table": "user",
	 "ts": 1696483273479,
	 "type": "INSERT"
	}`)

	index := dbtoes.NewIndex(
		dbtoes.WithForeignKey1("user_id"),
		dbtoes.WithPrimaryTable1("user"),
		dbtoes.WithTables([]string{"user", "class", "score", "subject"}),
		dbtoes.WithSynchronizer(NewStudentScoreSync()),
		dbtoes.WithEsClient(ESClient),
		dbtoes.WithEsTypedClient(ESTypedClient),
	)

	index.Start(msg)
}
