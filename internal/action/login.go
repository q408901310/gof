package action

import "context"

type aLogin struct{}

var (
	Login = aLogin{}
)

func (a *aLogin) zone(ctx context.Context) {

}
