package providers

import "go.uber.org/fx"

func Provide(opt fx.Option) *fx.App {
	return fx.New(
		opt,
	)
}
