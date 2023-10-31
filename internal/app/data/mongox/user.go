package mongox

import (
	"context"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/qiushenglei/gin-skeleton/internal/app/data"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDoc struct {
	StudentID string   `bson:"student_id"`
	Content   string   `bson:"content,omitempty"`
	Hobby     []string `bson:"hobby,omitempty"`
	Address   []string `bson:"address,omitempty"`
}

func Insert(c context.Context, student *entity.StudentSetDataRequest) {
	user := UserDoc{
		StudentID: student.StudentId,
		Content:   student.Content,
		Hobby:     student.Hobby,
		Address:   student.Address,
	}
	res, err := data.MongoClient.Database("order").Collection("user").InsertOne(c, user)
	if err != nil {
		panic(err)
	}
	logging.Info("mongo insert id ", res.InsertedID)
	return
}

// Upsert 更新或插入 on duplicate key update
func Upsert(c context.Context, student *entity.StudentSetDataRequest) {
	upsertFilter := bson.D{{"address", bson.D{{"$all", bson.A{"家家乐"}}}}}
	//upsertFilter := bson.D{{"student_id", "1713656556052025344"}}
	upsertDoc := bson.D{
		{
			"$set", bson.D{
				{"content", "6666我的宝贝"},
				{"hobby", bson.A{"马佳佳", "猪脚饭"}},
				{"Address", bson.A{"中国", "河南", "曹县"}},
			},
		},
		//{
		//	"$push", bson.D{
		//		{"address", "麻布囤"},
		//	},
		//},
	}

	// 默认是不开启修改未找到后插入
	updateOptions := options.Update().SetUpsert(true)

	res, err := data.MongoClient.Database("order").Collection("user").UpdateOne(c, upsertFilter, upsertDoc, updateOptions)
	if err != nil {
		panic(err)
	}
	logging.Info("mongo insert id ", res.UpsertedID)
	return
}

func GetUserMongoBySearchCond(c context.Context, cond *entity.SearchCond) ([]UserDoc, error) {
	if cond.Address == "" && cond.Content == "" && cond.StudentId == "" {
		return nil, nil
	}

	filter := QueryFilter(cond)
	cur, err := data.MongoClient.Database("order").Collection("user").Find(c, filter)
	if err != nil {
		panic(err)
	}

	var user []UserDoc
	if err := cur.All(c, &user); err != nil {
		panic(err)
	}

	if len(user) <= 0 {
		return user, errorpkg.ErrChildQueryNil
	}
	return user, err
}

func QueryFilter(cond *entity.SearchCond) bson.D {
	var condSet []bson.D

	// StudentID “==” 判断
	if cond.StudentId != "" {
		condSet = append(condSet, bson.D{{"student", cond.StudentId}})
	}

	// Content 正则判断
	if cond.Content != "" {
		condSet = append(condSet, bson.D{{"content", bson.D{{"$regex", "才会赢$"}}}})
	}

	// Address数组 包含判断
	if cond.Address != "" {
		condSet = append(condSet, bson.D{{"address", bson.D{{"$all", bson.A{cond.Address}}}}})
	}

	filter := bson.D{{
		"$and",
		condSet,
	}}
	return filter
}
