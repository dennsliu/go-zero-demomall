package logic

import (
	"context"
	"errors"

	"go-zero-demomall/order/api/internal/svc"
	"go-zero-demomall/order/api/internal/types"
	"go-zero-demomall/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdRequest{
		Id: "1",
	})
	if err != nil {
		return nil, err
	}
	if user.Name != "test" {
		return nil, errors.New("User not found")
	}
	returnString := "testing order ID: " + req.Id
	return &types.OrderReply{
		Id:   req.Id,
		Name: returnString,
	}, nil
}
