创建mall服务
mkdir go-zero-demomall
cd go-zero-demomall 
go mod init go-zero-demomall

创建user rpc服务

$ mkdir -p mall/user/rpc
添加user.proto文件，增加getUser方法

$ vim mall/user/rpc/user.proto
写入ptoto代码,注意： 每一个 *.proto文件只允许有一个service error: only one service expected

生成rpc代码：
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

写业务逻辑
$ vim internal/logic/getuserlogic.go


创建 order api服务
# 回到项目根目录
$ mkdir -p order/api && cd order/api

添加api文件
$ vim order.api


生成order服务
$ goctl api go -api order.api -dir .

添加user rpc配置 到order
$ vim internal/config/config.go
package config
import (
    "github.com/zeromicro/go-zero/zrpc"
    "github.com/zeromicro/go-zero/rest"
)

type Config struct {
    rest.RestConf
    UserRpc zrpc.RpcClientConf
}
添加yaml配置
$ vim etc/order.yaml 
Name: order
Host: 0.0.0.0
Port: 8888
UserRpc:
  Etcd:
    Hosts:
    - 127.0.0.1:2379
    Key: user.rpc

完善服务依赖

$ vim internal/svc/servicecontext.go
package svc

import (
    "go-zero-demo/mall/order/api/internal/config"
    "go-zero-demo/mall/user/rpc/user"

    "github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
    Config  config.Config
    UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config:  c,
        UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
    }
}

添加order演示逻辑

给 getorderlogic 添加业务逻辑

$ vim internal/logic/getorderlogic.go
package logic

import (
    "context"
    "errors"

    "go-zero-demo/mall/order/api/internal/svc"
    "go-zero-demo/mall/order/api/internal/types"
    "go-zero-demo/mall/user/rpc/types/user"

    "github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOrderLogic {
    return GetOrderLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (*types.OrderReply, error) {
    user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
        Id: "1",
    })
    if err != nil {
        return nil, err
    }

    if user.Name != "test" {
        return nil, errors.New("用户不存在")
    }

    return &types.OrderReply{
        Id:   req.Id,
        Name: "test order",
    }, nil
}