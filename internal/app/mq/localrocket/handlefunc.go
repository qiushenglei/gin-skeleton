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
	logs.Log.Info(ctx, string(msg.Body))

	index := dbtoes.NewIndex(
		dbtoes.WithForeignKey(student_score_idx.ForeignKey),
		dbtoes.WithPrimaryTable1("user"),
		dbtoes.WithTables([]string{"user", "class", "score", "subject"}),
		dbtoes.WithSynchronizer(student_score_idx.NewStudentScoreSync()),
		dbtoes.WithEsTypedClient(data.TypedESClient),
	)
	index.Start(msg.Body)
	return nil
}
