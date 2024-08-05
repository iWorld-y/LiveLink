package handler

import (
	"context"
	linkApi "github.com/iWorld-y/LiveLink/idl/pb/link"
)

type Link struct {
	linkApi.UnimplementedLinkServer
}

func (l *Link) GetMessage(ctx context.Context, link *linkApi.LinkClient) (*proto.GetLinkResp, error) {

}
