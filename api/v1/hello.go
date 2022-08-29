package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}


//type CardSelectReq struct {
//	g.Meta `path:"/hello" tags:"Hello" summary:"You first hello api"`
//	CardSelectMsg floor.CardSelectMsg
//}
//
//type CardSelectRes struct {
//	g.Meta `mime:"json" example:"string"`
//}