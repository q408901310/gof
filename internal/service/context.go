package service

type (
	sContext struct{}
)

var (
	insContext = sContext{}
)

func Context() *sContext {
	return &insContext
}
