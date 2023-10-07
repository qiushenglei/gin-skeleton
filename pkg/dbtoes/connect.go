package dbtoes

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"io"
)

func NewESClient(cfg elasticsearch.Config) *elasticsearch.Client {
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// ping
	resp, err := client.Ping()
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	logs.Log.Info(context.Background(), string(body))

	// 本地调试代码，查看index的mapping
	i, err := client.Indices.GetMapping(
		client.Indices.GetMapping.WithIndex("student_score_idx"),
	)
	if err != nil {
		panic(err)
	}
	body, err = io.ReadAll(i.Body)
	if err != nil {
		panic(err)
	}

	return client
}
