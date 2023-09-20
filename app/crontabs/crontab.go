package crontabs

import (
	"fmt"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
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
	logs.Log.Error("run Job")
}
