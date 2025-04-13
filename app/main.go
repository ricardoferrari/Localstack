package main

import (
	"log"
	"os"

	productModels "github.com/ricardoferrari/gorest/models"
	bootstrap "github.com/ricardoferrari/gorest/modules"
	productUseCase "github.com/ricardoferrari/gorest/usecases"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	fx.New(
		fx.Options(bootstrap.UtilsModule, bootstrap.ModuleControllers),
		fx.WithLogger(func() fxevent.Logger {
			return &fxevent.ConsoleLogger{W: os.Stdout}
		}),
		fx.Invoke(func(l productUseCase.ProductListInterface) {
			l.AddProduct(productModels.Product{Id: 1, Title: "Shoes", Price: 19.99})
			l.AddProduct(productModels.Product{Id: 2, Title: "Shirt", Price: 9.99})
			l.AddProduct(productModels.Product{Id: 3, Title: "Pants", Price: 14.99})
		}),
		fx.Invoke(func(l productUseCase.ProductListInterface, logger *log.Logger) {
			logger.Println("Products:", l.GetProducts())
		}),
	).Run()
}
