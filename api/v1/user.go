package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gof/internal/model"
)

type UserLoginReq struct {
	g.Meta   `path:"/user/login" tags:"用户" method:"post" summary:"登录"`
	Passport string `json:"passport"`
	Password string `json:"password"`
}

type UserLoginRes struct {
	SessionId string                  `json:"sessionId"`
	List      []*model.ServerListItem `json:"list"`
}

type UserRegisterReq struct {
	g.Meta   `path:"/user/register" tags:"用户" method:"post" summary:"注册"`
	Channel  int    `json:"channel"`
	Passport string `json:"passport"`
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}

type UserRegisterRes struct {
}

type UserGuestReq struct {
	g.Meta `path:"/user/guest" tags:"用户" method:"get" summary:"游客"`
}

type UserGuestRes struct {
}
