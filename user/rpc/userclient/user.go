package userclient

import (
	"context"
	"go-zero-demomall/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type (
	IdRequest    = user.IdRequest
	UserResponse = user.UserResponse

	User interface {
		GetUser(ctx context.Context, in *IdRequest) (*UserResponse, error)
	}
	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdRequest) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in)
}
