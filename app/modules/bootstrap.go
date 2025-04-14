package bootstrap

import (
	betControllers "github.com/ricardoferrari/localstack/controllers"
	messageRepository "github.com/ricardoferrari/localstack/repositories"
	betUseCase "github.com/ricardoferrari/localstack/usecases"
	"go.uber.org/fx"
)

var UtilsModule = fx.Module("bootstrap",
	fx.Provide(messageRepository.NewMessageRepository),
	fx.Provide(messageRepository.NewBetRepository),
	fx.Provide(betUseCase.NewMessageUseCase),
	fx.Provide(NewLogger),
	fx.Provide(NewHTTPServer),
)

var ModuleControllers = fx.Module("controllers",
	fx.Invoke(betControllers.NewBetController),
)
