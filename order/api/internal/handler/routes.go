// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-demomall/order/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/order/get/:id",
				Handler: getOrderHandler(serverCtx),
			},
		},
	)
}
