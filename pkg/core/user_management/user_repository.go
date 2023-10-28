package user_management

import (
	"context"
	"fmt"

	"github.com/fahmifan/autograd/pkg/core"
	"github.com/fahmifan/autograd/pkg/core/auth"
	"github.com/fahmifan/autograd/pkg/dbmodel"
	"gorm.io/gorm"
)

type ManagedUserWriter struct{}

func (ManagedUserWriter) SaveUserWithPassword(ctx context.Context, tx *gorm.DB, user ManagedUser, password auth.CipherPassword) error {
	model := dbmodel.User{
		Base: dbmodel.Base{
			ID:       user.ID,
			Metadata: core.NewModelMetadata(user.TimestampMetadata),
		},
		Name:     user.Name,
		Email:    user.Email,
		Password: string(password),
		Role:     string(user.Role),
	}

	return tx.Save(&model).Error
}

type ManagedUserReader struct{}

func (ManagedUserReader) FindUserByID(ctx context.Context, tx *gorm.DB, id string) (ManagedUser, error) {
	var model dbmodel.User
	if err := tx.First(&model, id).Error; err != nil {
		return ManagedUser{}, err
	}

	return ManagedUser{
		ID:                model.ID,
		Name:              model.Name,
		Email:             model.Email,
		Role:              auth.Role(model.Role),
		TimestampMetadata: core.TimestampMetaFromModel(model.Metadata),
	}, nil
}

type FindAllManagedUsersRequest struct {
	Page  int32
	Limit int32
}

type FindAllManagedUsersResponse struct {
	Users []ManagedUser
	Count int32
}

func (ManagedUserReader) FindAll(ctx context.Context, tx *gorm.DB, req FindAllManagedUsersRequest) (res FindAllManagedUsersResponse, err error) {
	var models []dbmodel.User
	if err := tx.Limit(int(req.Limit)).Offset(int((req.Page - 1) * req.Limit)).Find(&models).Error; err != nil {
		return res, fmt.Errorf("find all: %w", err)
	}

	var count int64
	if err := tx.Model(&dbmodel.User{}).Count(&count).Error; err != nil {
		return res, fmt.Errorf("count: %w", err)
	}

	res.Users = make([]ManagedUser, len(models))
	for i, model := range models {
		res.Users[i] = ManagedUser{
			ID:                model.ID,
			Name:              model.Name,
			Email:             model.Email,
			Role:              auth.Role(model.Role),
			TimestampMetadata: core.TimestampMetaFromModel(model.Metadata),
		}
	}

	res.Count = int32(count)

	return res, nil
}
