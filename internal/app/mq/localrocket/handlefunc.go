package localrocket

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/qiushenglei/gin-skeleton/internal/app/data"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/es/sync/student_score_idx"
	"github.com/qiushenglei/gin-skeleton/pkg/dbtoes"
	"github.com/qiushenglei/gin-skeleton/pkg/logs"
)

// OrderPayEventHandle 订单支付事件处理
func OrderPayEventHandle(ctx context.Context, msg *primitive.MessageExt) error {
	fmt.Println(string(msg.Body))
	fmt.Printf("msg: %+v", msg)
	logs.Log.Info(ctx, string(msg.Body))
	return nil
}

// OrderPayEventHandle 订单支付事件处理
func CanalSyncESHandle(ctx context.Context, msg *primitive.MessageExt) error {
	logs.Log.Info(ctx, "这里是es的handle", msg.MsgId, msg.Message.GetProperty("CONSUME_START_TIME"), string(msg.Body))
	if msg.MsgId == "7F0000010065384AD17B5797E8990001" {
		fmt.Println("debug")
	}
	index := dbtoes.NewIndex(
		dbtoes.WithForeignKey(student_score_idx.ForeignKey),
		dbtoes.WithPrimaryTable1("user"),
		dbtoes.WithTables([]string{"user", "class", "score", "subject"}),
		dbtoes.WithSynchronizer(student_score_idx.NewStudentScoreSync()),
		dbtoes.WithEsTypedClient(data.TypedESClient),
	)
	err := index.Start(msg.Body)
	return err
}

// DelayHandle 延时任务处理
func DelayHandle(ctx context.Context, msg *primitive.MessageExt) error {
	logs.Log.Info(ctx, "delay consumer msg", msg.MsgId, msg.Message.GetProperty("CONSUME_START_TIME"), string(msg.Body))

	return nil
}
