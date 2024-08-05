package handler

import (
	"context"
	"log"

	linkApi "github.com/iWorld-y/TradeHub/idl/pb/link"
	"github.com/iWorld-y/TradeHub/internal/module"
)

type LinkHandler struct {
	linkApi.UnimplementedLinkServer
	linkM module.LinkModule
}

func (l *LinkHandler) GetMessage(ctx context.Context, req *linkApi.GetLinkReq) (*linkApi.GetLinkResp, error) {
	data, err := l.linkM.GetMessage(ctx, req)
	if err != nil {
		log.Println(err)
	}
	return &linkApi.GetLinkResp{
		Errno:  0,
		Errmsg: "",
		Data:   data,
	}, nil
}
