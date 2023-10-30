package test

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/qiushenglei/gin-skeleton/internal/app/sentinelx"
	"testing"
)

func TestMain(m *testing.M) {

	//cmd.RegistAll("test")

	m.Run()
}

func BenchmarkRateLimit(b *testing.B) {

	//b.ReportAllocs()
	//sentinel.InitDefault()
	//b.ReportMetric(b.N,  "ops/s")
	sentinelx.RegisterSentinelRule()
	for i := 0; i < b.N; i++ {
		e, err := sentinel.Entry(sentinelx.GlobalRateLimit)
		if err != nil {
			b.Error(err)
		} else {
			//b.Run()
			fmt.Println("ok ", i)
			e.Exit()
		}
	}

}
