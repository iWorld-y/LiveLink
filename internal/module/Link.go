package module

import (
	"context"
	linkApi "github.com/iWorld-y/TradeHub/idl/pb/link"
)

type LinkModule struct {
}

func (l *LinkModule) GetMessage(ctx context.Context, req *linkApi.GetLinkReq) (*linkApi.GetLinkData, error) {
	data := &linkApi.GetLinkData{
		GroupId: req.GroupId,
		List: []*linkApi.GetLinkList{
			{
				UserId: 1,
				Msg:    "hello",
			},
		},
	}
	return data, nil
}
