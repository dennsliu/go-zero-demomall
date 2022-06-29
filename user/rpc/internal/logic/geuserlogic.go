package logic

import (
	"context"

	"go-zero-demomall/mall/user/rpc/internal/svc"
	"go-zero-demomall/mall/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GeUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGeUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GeUserLogic {
	return &GeUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GeUserLogic) GeUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	return &user.UserResponse{
		Id:   "1",
		Name: "Denny testing",
	}, nil
}
