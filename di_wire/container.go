//go:build wireinject
// +build wireinject

package di_wire

import "github.com/google/wire"

func CreateConcatService() *ConcatService {
	panic(wire.Build(
		wire.Struct(new(Logger), "*"),
		NewHttpClient,
		NewConcatService,
	))
}
