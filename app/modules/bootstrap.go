package bootstrap

import (
	productControllers "github.com/ricardoferrari/gorest/controllers"
	productRepository "github.com/ricardoferrari/gorest/repositories"
	productUseCase "github.com/ricardoferrari/gorest/usecases"
	"go.uber.org/fx"
)

var UtilsModule = fx.Module("bootstrap",
	fx.Provide(productRepository.NewProductRepository),
	fx.Provide(productUseCase.NewProductUseCase),
	fx.Provide(NewLogger),
	fx.Provide(NewHTTPServer),
)

var ModuleControllers = fx.Module("controllers",
	fx.Invoke(productControllers.NewProductController),
)
