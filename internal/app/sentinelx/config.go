package sentinelx

import (
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/flow"
)

var (
	GlobalRateLimiter  = "global_limiter"
	FindApiRateLimiter = "find"
	GlobalBreaker      = "global_breaker"
)

func RegisterSentinelRule() {
	registerRateLimit()
	registerBreaker()
}

// 注册限流规则
func registerRateLimit() {
	// 配置一条限流规则
	_, err := flow.LoadRules([]*flow.Rule{
		{
			// 100/s的速率
			Resource:               GlobalRateLimiter,
			TokenCalculateStrategy: flow.Direct, //
			ControlBehavior:        flow.Reject, // 超出的请求量直接拒绝
			Threshold:              100000,      // 1单位时间窗口下的请求量
			StatIntervalInMs:       1000,        // 时间窗口1000ms
		},
		{
			Resource:               FindApiRateLimiter,
			Threshold:              10,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
		},
	})
	if err != nil {
		panic(err)
	}
}

// 注册熔断规则
func registerBreaker() {
	// 配置一条限流规则
	_, err := circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		{
			// 100/s的速率
			Resource:         GlobalBreaker,
			Strategy:         circuitbreaker.ErrorCount, // 熔断策略：当前是错误统计数
			RetryTimeoutMs:   60 * 1000,                 // 即熔断触发后持续的时间（单位为 ms）。资源进入熔断状态后，在配置的熔断时长内，请求都会快速失败。熔断结束后进入探测恢复模式（HALF-OPEN）
			MinRequestAmount: 1000,                      // 静默数量，时间窗口下不超过这个值的请求都是在静默期内
			Threshold:        50,                        // 对于错误统计计数策略，窗口时间内50次就会触发熔断
			StatIntervalMs:   1000,                      // 时间窗口
		},
	})
	if err != nil {
		panic(err)
	}
}
