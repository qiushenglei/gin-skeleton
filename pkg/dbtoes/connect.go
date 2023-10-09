package dbtoes

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
)

func NewESClient(cfg elasticsearch.Config) *elasticsearch.Client {
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// ping
	_, err = client.Ping()
	if err != nil {
		panic(err)
	}

	// 本地调试代码，查看index的mapping
	i, err := client.Indices.GetMapping(
		client.Indices.GetMapping.WithIndex("student_score_idx"),
	)
	if err != nil {
		panic(err)
	}
	_, err = io.ReadAll(i.Body)
	if err != nil {
		panic(err)
	}

	return client
}

func NewESTypedClient1(cfg elasticsearch.Config) *elasticsearch.TypedClient {
	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		panic(err)
	}

	// ping
	if ok, err := client.Ping().Do(context.Background()); err != nil || !ok {
		panic(err)
	}

	// 本地调试代码，查看index的mapping
	resp, err := client.Indices.GetMapping().Index("student_score").Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	return client
}
