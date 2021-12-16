// Code generated by goctl. DO NOT EDIT!
// Source: github.com/chabby-hao/acitivity-rpc.proto

package server

import (
	"context"

	"github.com/chabby-hao/acitivity-rpc/activity"
	"github.com/chabby-hao/acitivity-rpc/internal/logic"
	"github.com/chabby-hao/acitivity-rpc/internal/svc"
)

type ActivityServer struct {
	svcCtx *svc.ServiceContext
}

func NewActivityServer(svcCtx *svc.ServiceContext) *ActivityServer {
	return &ActivityServer{
		svcCtx: svcCtx,
	}
}

func (s *ActivityServer) ListAll(ctx context.Context, in *activity.ListAllRequest) (*activity.ListAllResponse, error) {
	l := logic.NewListAllLogic(ctx, s.svcCtx)
	return l.ListAll(in)
}

func (s *ActivityServer) ListByType(ctx context.Context, in *activity.ListByTypeRequest) (*activity.ListByTypeResponse, error) {
	l := logic.NewListByTypeLogic(ctx, s.svcCtx)
	return l.ListByType(in)
}
