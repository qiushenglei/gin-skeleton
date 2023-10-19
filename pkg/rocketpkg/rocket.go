package rocketpkg

import (
	"context"
	"errors"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
)

const Expire = "Expire"

type Event struct {
	EventName          EventName
	Topic              string   // 消息主题
	Tags               []string // 消息tags
	Delay              *Delay
	ProducerGroup      string
	ConsumerNameSpace  string                                             // 消费者group
	ConsumerGroupName  string                                             // 所属消费者组
	ConsumerNum        int                                                // 消费者个数
	ConsumerHandleFunc func(context.Context, *primitive.MessageExt) error // 处理消息方法
}

type Delay struct {
	Expire int
}

type EventConf map[EventName]Event

type EventName string

var TimeLevelMap = map[int]int{
	1:    1,
	5:    2,
	10:   3,
	30:   4,
	60:   5,
	120:  6,
	180:  7,
	240:  8,
	300:  9,
	360:  10,
	420:  11,
	480:  12,
	540:  13,
	600:  14,
	1200: 15,
	1800: 16,
	3600: 17,
	7200: 18,
}

// GetEvent 根据事件名称，从事件配置map中获取配置信息(topic 和 tags等)
func GetEvent(ef EventConf, en EventName) (event Event, err error) {
	var ok bool
	if event, ok = ef[en]; ok != true {
		return Event{}, errors.New("no register event")
	}
	return event, nil
}

// HandleDelayExpire 处理延迟时间
func HandleDelayExpire(delay *Delay) (left, level int) {
	//GOClientLongest := 2 * 60 * 60
	GOClientLongest := 10

	// 判断delay是否合规  delay.Expire % GOClientLongest 得到最后2小时的值，进到CalculateLevel判断是否合规
	CalculateLevel(delay.Expire % GOClientLongest)

	// 小于2h的直接设置level
	if delay.Expire <= GOClientLongest {
		return 0, CalculateLevel(delay.Expire)
	}

	// Go client的最长延时等级
	return delay.Expire - GOClientLongest, CalculateLevel(GOClientLongest)
}

func CalculateLevel(expire int) (level int) {
	var ok bool

	if expire <= 0 {
		return level
	}

	if level, ok = TimeLevelMap[expire]; !ok {
		panic(errorpkg.NewBizErrx(errorpkg.CodeFalse, "delay time false"))
	}
	return level
}
