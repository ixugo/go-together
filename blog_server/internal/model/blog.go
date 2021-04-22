package model

import pb "together/proto"

type EsBlogReq struct {
	Url   string                  `json:"url"`
	Menus []*pb.GetListReply_Data `json:"menus"`
}
