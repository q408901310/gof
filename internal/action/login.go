package action

import (
	"context"
	"gof/internal/pb"
)

type aLogin struct{}

var (
	Login = aLogin{}
)

func (a *aLogin) start(ctx context.Context, data *pb.DataMsg) {

}
