package crontabs

import (
	"context"
	"fmt"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"time"
)

var (
	JobList []Job = []Job{
		Job{
			Desc:     "example Job",
			Schedule: "* * * * *",
			Fn:       ExampleJob,
		},
	}
)

type Job struct {
	Desc     string
	Schedule string
	Fn       func()
}

func ExampleJob() {
	fmt.Println("run Job")
	logs.Log.Error(context.Background(), "run Job")

	time.Sleep(5 * time.Second)

	fmt.Println("安全退出")
}
