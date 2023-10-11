package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gof/internal/model"
)

type ServerListReq struct {
	g.Meta `path:"/server/list" tags:"区服" method:"get" summary:"区服列表"`
}

type ServerListRes struct {
	List []*model.ServerListItem `json:"list"`
}

type ServerChooseReq struct {
	g.Meta `path:"/user/login" tags:"用户" method:"post" summary:"选择区服"`
	Server uint `json:"server"`
}

type ServerChooseRes struct {
}
