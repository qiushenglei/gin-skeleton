package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/data"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mongox"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/canal_test/model"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/canal_test/query"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/models"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"github.com/qiushenglei/gin-skeleton/pkg/errorpkg"
	"gorm.io/gorm/clause"
)

func SetData(ctx *gin.Context, student *entity.StudentSetDataRequest) uint32 {
	var user *model.User
	q := query.Use(data.MySQLCanalTestClient)

	q.Transaction(func(tx *query.Query) error {

		// set Class
		class := models.NewClass(student.ClassInfo)
		if err := tx.Class.WithContext(ctx).Clauses(
			clause.OnConflict{
				Columns: []clause.Column{
					{Name: "id"},
				},
				DoUpdates: clause.AssignmentColumns([]string{"class_name", "grade", "update_time"}),
			},
		).Create(class); err != nil {
			panic(err)
		}
		student.ClassId = int(class.ID)

		// set user
		user = models.NewUser(student)
		err := tx.User.Table("user").WithContext(ctx).Clauses(clause.OnConflict{
			// 如果id冲突就执行 on duplicate key update 后面的字段
			Columns: []clause.Column{
				{Name: "id"},
				{Name: "student_id"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"class_id",
				"student_id",
				"label",
				"username",
				"update_time",
				"is_deleted",
			}),
		}).Create(user) //这里用save方法都不行，必须是create
		if err != nil {
			panic(err)
		}

		// set Subject
		subject := models.NewSubject(student)
		err = tx.Subject.WithContext(ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{
				clause.Column{
					Name: "id",
				},
			},
			DoUpdates: clause.AssignmentColumns([]string{"subject_name"}),
		}).CreateInBatches(subject, len(subject))
		if err != nil {
			panic(err)
		}

		// set score
		scores := models.NewScores(student.ScoreInfo, subject, student.StudentId)
		if err := tx.Score.WithContext(ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.AssignmentColumns([]string{"student_id", "subject_id", "score", "update_time"}),
		}).CreateInBatches(scores, len(scores)); err != nil {
			panic(err)
		}
		return nil
	})

	// set user mongodb
	//mongox.Insert(ctx, student)
	mongox.Upsert(ctx, student)

	return user.ID
}

func GetData(c *gin.Context, cond *entity.SearchCond) {
	// get user
	GetUserData(c, cond)

	// get other
}

func GetUserData(c *gin.Context, cond *entity.SearchCond) ([]*model.User, error) {
	q := query.Q.User.Table("user").WithContext(c)

	// class cond
	class, err := models.GetClassBySearchCond(c, cond)
	if errors.Is(err, errorpkg.ErrChildQueryNil) {
		// 没有满足条件的数据
		return nil, nil
	} else if class != nil {
		q = q.Where(query.Q.User.ClassID.Eq(class.ID))
	}

	// score cond
	scores, err := models.GetScoreBySearchCond(c, cond)
	if errors.Is(err, errorpkg.ErrChildQueryNil) {
		// 没有满足条件的数据
		return nil, nil
	} else if len(scores) > 0 {
		var studentIds []string
		for _, v := range scores {
			studentIds = append(studentIds, v.StudentID)
		}
		q = q.Where(query.Q.User.StudentID.In(studentIds...))
	}

	// mongo cond
	mongoRes, err := mongox.GetUserMongoBySearchCond(c, cond)
	if errors.Is(err, errorpkg.ErrChildQueryNil) {
		// 没有满足条件的数据
		return nil, nil
	} else {
		var studentIds []string
		for _, v := range mongoRes {
			studentIds = append(studentIds, v.StudentID)
		}
		q.Where(query.Q.User.StudentID.In(studentIds...))
	}

	var users []*model.User
	if err := q.Scan(&users); err != nil {
		panic(err)
	}
	return users, nil
}
