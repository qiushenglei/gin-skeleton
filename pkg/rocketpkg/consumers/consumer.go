package consumers

import (
	"errors"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	consumer2 "github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/qiushenglei/gin-skeleton/internal/app/configs"
	"github.com/qiushenglei/gin-skeleton/internal/app/global/utils"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
	"github.com/qiushenglei/gin-skeleton/pkg/rocketpkg"
	"golang.org/x/net/context"
	"strconv"
	"strings"
	"sync"
)

type RocketMQConsumers struct {
	ef rocketpkg.EventConf // 私有rocket的事件配置map
}

// NewConsumers 创建消费结构体
func NewConsumers(ef rocketpkg.EventConf) *RocketMQConsumers {
	return &RocketMQConsumers{
		ef: ef,
	}
}

// Start 开始订阅消费
func (c *RocketMQConsumers) Start() error {
	if len(c.ef) <= 0 {
		return nil
	}

	for en, ef := range c.ef {
		err := c.newConsumerByEvent(ef)
		if err != nil {
			fmt.Printf("%s消费者创建失败: %s\n", en, err.Error())
			return err
		} else {
			fmt.Printf("%s消费者创建成功\n", en)
		}
	}
	return nil
}

// newConsumerByEvent 内部根据事件配置创建rocketmq消费者
func (c *RocketMQConsumers) newConsumerByEvent(e rocketpkg.Event) error {
	if e.ConsumerNum <= 0 {
		return errors.New("注册消费者个数为0个")
	}
	ctx := context.Background()

	wg := sync.WaitGroup{}

	// 启动多个消费者，加入同一个消费者组
	for i := 0; i < e.ConsumerNum; i++ {
		wg.Add(1)
		index := i

		handleFunc := func(ctx context.Context) {
			//go func(ctx context.Context) {
			defer wg.Done()

			// 此循环中，业务逻辑要求e是一样的，所以e逃逸也没关系
			// i变量会逃逸到堆上,所以传入下一个函数栈帧的i是一个地址变量，指向了堆上的i，i会一直变，所以不能直接使用i变量。需要for循环内加一个局部变量index:=i
			err := c.newPushConsumers(e, index)
			if err != nil {
				panic(err)
			}
		}
		utils.Go(ctx, handleFunc, nil)
	}
	wg.Wait()

	fmt.Printf("%s启动%d完毕\n", e.EventName, e.ConsumerNum)

	// 创建pushConsumer就不阻塞了。pushConsumer.Start会创建自己的协程， 给业务逻辑的mainGoroutine来处理生命周期

	return nil
}

// newPushConsumers 创建pushConsumer
func (c *RocketMQConsumers) newPushConsumers(e rocketpkg.Event, serialNumber int) error {
	// 创建pushConsumer
	// 这个host可以加到配置项里，可以定义多个rocketmq-server的连接
	host := fmt.Sprintf("%s:%s", configs.EnvConfig.GetString("ROCKETMQ_TEST_HOST"), configs.EnvConfig.GetString("ROCKETMQ_TEST_PORT"))
	consumer, err := rocketmq.NewPushConsumer(
		consumer2.WithGroupName(e.ConsumerGroupName), //指定消费者组
		consumer2.WithNamespace(e.NameSpace),
		consumer2.WithInstance(e.ConsumerGroupName+strconv.Itoa(serialNumber)),     //同一消费者组的不同消费者标识，不加上会有问题
		consumer2.WithNsResolver(primitive.NewPassthroughResolver([]string{host})), //服务ip
		consumer2.WithRetry(5),                                                     // 重试次数
	)

	if err != nil {
		return err
	}

	// 消息选择器，使用tag方式，指定消费某些tag的消息
	selector := consumer2.MessageSelector{
		Type:       consumer2.TAG,
		Expression: c.spliceTagString(e),
	}

	// 订阅
	err = consumer.Subscribe(e.Topic, selector, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer2.ConsumeResult, error) {
		// 接住panic，做重试或者其他任务
		defer func() {
			if r := recover(); r != nil {
				// TODO::重试或者结束
			}
		}()

		// 目前只处理了只有一个消息的情况
		err := e.ConsumerHandleFunc(ctx, ext[0])
		if err != nil {
			// TODO 给log加上各种服务文件类型，目前只有web的
			logs.Log.Error(ctx)
			return consumer2.ConsumeRetryLater, err
		}

		return consumer2.ConsumeSuccess, nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Start the PushConsumer for consuming message
	err = consumer.Start()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Printf("启动序号为：%d成功\n", serialNumber)

	return nil
}

// spliceTagString 拼接tags
func (c *RocketMQConsumers) spliceTagString(e rocketpkg.Event) string {
	return strings.Join(e.Tags, " || ")
}
