package service

import (
	"context"

	"connectrpc.com/connect"
	"github.com/fahmifan/autograd/pkg/core"
	"github.com/fahmifan/autograd/pkg/core/assignments/assignments_cmd"
	"github.com/fahmifan/autograd/pkg/core/assignments/assignments_query"
	"github.com/fahmifan/autograd/pkg/core/auth/auth_cmd"
	"github.com/fahmifan/autograd/pkg/core/mediastore/mediastore_cmd"
	"github.com/fahmifan/autograd/pkg/core/student_assignment/student_assignment_cmd"
	"github.com/fahmifan/autograd/pkg/core/student_assignment/student_assignment_query"
	"github.com/fahmifan/autograd/pkg/core/user_management/user_management_cmd"
	"github.com/fahmifan/autograd/pkg/core/user_management/user_management_query"
	autogradv1 "github.com/fahmifan/autograd/pkg/pb/autograd/v1"
	"github.com/fahmifan/autograd/pkg/pb/autograd/v1/autogradv1connect"
	"gorm.io/gorm"
)

type Service struct {
	*auth_cmd.AuthCmd
	*user_management_query.UserManagementQuery
	*user_management_cmd.UserManagementCmd
	*assignments_cmd.AssignmentCmd
	*assignments_query.AssignmentsQuery
	*mediastore_cmd.MediaStoreCmd
	*student_assignment_query.StudentAssignmentQuery
	*student_assignment_cmd.StudentAssignmentCmd
}

var _ autogradv1connect.AutogradServiceHandler = &Service{}

func NewService(gormDB *gorm.DB, jwtKey string, mediaCfg core.MediaConfig) *Service {
	coreCtx := &core.Ctx{
		GormDB:      gormDB,
		JWTKey:      jwtKey,
		MediaConfig: mediaCfg,
	}
	return &Service{
		AuthCmd:                &auth_cmd.AuthCmd{Ctx: coreCtx},
		UserManagementCmd:      &user_management_cmd.UserManagementCmd{Ctx: coreCtx},
		UserManagementQuery:    &user_management_query.UserManagementQuery{Ctx: coreCtx},
		AssignmentCmd:          &assignments_cmd.AssignmentCmd{Ctx: coreCtx},
		AssignmentsQuery:       &assignments_query.AssignmentsQuery{Ctx: coreCtx},
		MediaStoreCmd:          &mediastore_cmd.MediaStoreCmd{Ctx: coreCtx},
		StudentAssignmentQuery: &student_assignment_query.StudentAssignmentQuery{Ctx: coreCtx},
		StudentAssignmentCmd:   &student_assignment_cmd.StudentAssignmentCmd{Ctx: coreCtx},
	}
}

func (service *Service) Ping(ctx context.Context, req *connect.Request[autogradv1.Empty]) (*connect.Response[autogradv1.PingResponse], error) {
	return &connect.Response[autogradv1.PingResponse]{
		Msg: &autogradv1.PingResponse{
			Message: "pong",
		},
	}, nil
}
