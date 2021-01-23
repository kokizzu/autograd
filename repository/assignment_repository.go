package repository

import (
	"context"

	"github.com/miun173/autograd/model"
	"github.com/miun173/autograd/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// AssignmentRepository ..
type AssignmentRepository interface {
	Create(ctx context.Context, assignment *model.Assignment) error
	DeleteByID(ctx context.Context, id int64) error
	FindAll(ctx context.Context, cursor model.Cursor) (assignments []*model.Assignment, count int64, err error)
	FindByID(ctx context.Context, id int64) (*model.Assignment, error)
	Update(ctx context.Context, assignment *model.Assignment) error
}

type assignmentRepo struct {
	db *gorm.DB
}

// NewAssignmentRepository ..
func NewAssignmentRepository(db *gorm.DB) AssignmentRepository {
	return &assignmentRepo{
		db: db,
	}
}

func (a *assignmentRepo) Create(ctx context.Context, assignment *model.Assignment) error {
	err := a.db.Create(assignment).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":        utils.Dump(ctx),
			"assignment": utils.Dump(assignment),
		}).Error(err)
		return err
	}

	return nil
}

func (a *assignmentRepo) DeleteByID(ctx context.Context, id int64) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.Dump(ctx),
		"id":  id,
	})

	tx := a.db.Begin()
	err := tx.Where("assignment_id = ?", id).Delete(&model.Submission{}).Error
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}

	err = tx.Where("id = ?", id).Delete(&model.Assignment{}).Error
	if err != nil {
		tx.Rollback()
		logger.Error(err)
		return err
	}

	tx.Commit()

	return nil
}

func (a *assignmentRepo) FindAll(ctx context.Context, cursor model.Cursor) (assignments []*model.Assignment, count int64, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":    utils.Dump(ctx),
		"cursor": utils.Dump(cursor),
	})

	err = a.db.Model(model.Assignment{}).Count(&count).Error
	if err != nil {
		logger.Error(err)
		return nil, count, err
	}

	err = a.db.Limit(int(cursor.GetSize())).Offset(int(cursor.GetOffset())).
		Order("created_at " + cursor.GetSort()).Find(&assignments).Error
	if err != nil {
		logger.Error(err)
		return nil, count, err
	}

	return
}

func (a *assignmentRepo) FindByID(ctx context.Context, id int64) (*model.Assignment, error) {
	assignment := &model.Assignment{}
	err := a.db.Where("id = ? ", id).Take(assignment).Error
	switch err {
	case nil: // ignore
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		logrus.WithFields(logrus.Fields{
			"ctx": utils.Dump(ctx),
			"id":  id,
		}).Error(err)
		return nil, err
	}

	return assignment, nil
}

func (a *assignmentRepo) Update(ctx context.Context, assignment *model.Assignment) error {
	err := a.db.Model(&model.Assignment{}).Where("id = ?", assignment.ID).Updates(assignment).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":        utils.Dump(ctx),
			"assignment": utils.Dump(assignment),
		}).Error(err)
		return err
	}

	return nil
}