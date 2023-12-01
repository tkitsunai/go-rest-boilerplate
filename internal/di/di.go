package di

import "github.com/samber/do"

func SetupProvider() *do.Injector {
	injector := do.New()
	return injector
}
