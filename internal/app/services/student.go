package services

import (
	"github.com/gin-gonic/gin"
	"github.com/qiushenglei/gin-skeleton/internal/app/data"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/models"
	"github.com/qiushenglei/gin-skeleton/internal/app/data/mysql/query"
	"github.com/qiushenglei/gin-skeleton/internal/app/entity"
	"gorm.io/gorm/clause"
)

func SetData(ctx *gin.Context, student *entity.StudentSetData) (interface{}, error) {
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
			return err
		}
		student.ClassId = int(class.ID)

		// set user
		user := models.NewUser(student)
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
			return err
		}

		// set score
		scores := models.NewScores(student.ScoreInfo)
		tx.Score.WithContext(ctx).Clauses().CreateInBatches()

		// set subject
		return nil
	})
	// set Class model

	// set User
	//_, err := model.SetData(ctx, student)
	//
	//if err != nil {
	//	tx.Rollback()
	//	return nil, nil
	//}

	// set score model

	// set Subject model

	return nil, nil
}
