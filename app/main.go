package main

import (
	"log"
	// "os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	// messageModels "github.com/ricardoferrari/localstack/models"
	// bootstrap "github.com/ricardoferrari/localstack/modules"
	// messageUseCase "github.com/ricardoferrari/localstack/usecases"
	// "go.uber.org/fx"
	// "go.uber.org/fx/fxevent"
)

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "localstack",
		Config: aws.Config{
			// 	Region: aws.String("us-east-1"),
			Endpoint: aws.String("http://localhost:4566"),
		},
	}))

	svc := sqs.New(sess)

	SendMessageOutput, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Title": {
				DataType:    aws.String("String"),
				StringValue: aws.String("The Whistler"),
			},
			"Author": {
				DataType:    aws.String("String"),
				StringValue: aws.String("John Grisham"),
			},
			"WeeksOn": {
				DataType:    aws.String("Number"),
				StringValue: aws.String("7"),
			},
		},
		MessageBody: aws.String("Information about current NY Times fiction bestseller for week of 23/11/2016."),
		QueueUrl:    aws.String("http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/teste"),
	})

	if err != nil {
		log.Println("Error", err)
	} else {
		log.Println("Success", SendMessageOutput.MessageId)
	}

	// fx.New(
	// 	fx.Options(bootstrap.UtilsModule, bootstrap.ModuleControllers),
	// 	fx.WithLogger(func() fxevent.Logger {
	// 		return &fxevent.ConsoleLogger{W: os.Stdout}
	// 	}),
	// 	fx.Invoke(func(l messageUseCase.MessageUseCaseInterface) {
	// 		l.AddMessage(messageModels.Message{Text: "Shoes"})
	// 		l.AddMessage(messageModels.Message{Text: "Shirt"})
	// 		l.AddMessage(messageModels.Message{Text: "Pants"})
	// 	}),
	// 	fx.Invoke(func(l messageUseCase.MessageUseCaseInterface, logger *log.Logger) {
	// 		logger.Println("Products:", l.GetMessages())
	// 	}),
	// ).Run()
}
