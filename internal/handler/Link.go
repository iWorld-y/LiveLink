package handler

import (
	"context"
	linkApi "github.com/iWorld-y/LiveLink/idl/pb/link"
	"log"
)

type LinkHandler struct {
	linkApi.UnimplementedLinkServer
}

func (l *LinkHandler) GetMessage(ctx context.Context, req *linkApi.GetLinkReq) (*linkApi.GetLinkResp, error) {
	log.Println(req)
	return &linkApi.GetLinkResp{
		Errno:  0,
		Errmsg: "",
		Data: &linkApi.GetLinkData{
			GroupId: 0,
			List: []*linkApi.GetLinkList{{
				UserId: 0,
				Msg:    "hello, world",
			}},
		},
	}, nil
}
