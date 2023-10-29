package crontabs

import (
	"context"
	"fmt"
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
	JobList1 = []*Job1{
		{
			Desc:     "example Job",
			Schedule: "* * * * *",
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
	//logs.Log.Error(context.Background(), "run Job")
}

type Job1 struct {
	Desc     string
	Schedule string
}

func (j *Job1) Run(ctx context.Context) {
	isRunning := true
	go func() {
		select {
		case <-ctx.Done():
			isRunning = false
		}
	}()

	// 如果业务可评估单次运行时长，且在kill时长内，比如现在是5秒kill时间，可以直接for
	// 其他有io的操作，如sql查询、redis，传ctx进去就好，这些包内部会读ctx，如果取消了，它们会直接返回的
	for isRunning {
		fmt.Println("业务正常运行", isRunning)
		// 可以对比下区别
		//time.Sleep(10 * time.Second) // 这里模拟一次执行要10秒
		time.Sleep(3 * time.Second) // 这里模拟一次执行要3秒
	}
	fmt.Println("业务安全关闭")
}
