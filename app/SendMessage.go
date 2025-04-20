// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
// snippet-start:[sqs.go.receive_messages]
package main

// snippet-start:[sqs.go.receive_messages.imports]
import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// GetQueueURL gets the URL of an Amazon SQS queue
// Inputs:
//
//	sess is the current session, which provides configuration for the SDK's service clients
//	queueName is the name of the queue
//
// Output:
//
//	If success, the URL of the queue and nil
//	Otherwise, an empty string and an error from the call to
func getQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
	// snippet-start:[sqs.go.receive_messages.queue_url]
	svc := sqs.New(sess)

	urlResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	// snippet-end:[sqs.go.receive_messages.queue_url]
	if err != nil {
		return nil, err
	}

	return urlResult, nil
}

// Send message to an Amazon SQS queue
// Inputs:
//
//	sess is the current session, which provides configuration for the SDK's service clients
//	queueURL is the URL of the queue
//	handle is the message handle
//
// Output:
//
//	If success, the URL of the queue and nil
//	Otherwise, an empty string and an error from the call to DeleteMessage
func SendMessage(sess *session.Session, queueURL *string, handle *string) (string, error) {
	svc := sqs.New(sess)
	_, err := svc.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    queueURL,
		MessageBody: handle,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return *queueURL, nil
}

func main() {
	// snippet-start:[sqs.go.receive_messages.args]
	queue := flag.String("q", "", "The name of the queue")
	timeout := flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	message := flag.String("s", "Empty message", "The message you want to send")
	flag.Parse()

	if *message == "" {
		fmt.Println("You must supply a message")
		return
	}

	if *queue == "" {
		fmt.Println("You must supply the name of a queue (-q QUEUE)")
		return
	}

	if *timeout < 0 {
		*timeout = 0
	}

	if *timeout > 12*60*60 {
		*timeout = 12 * 60 * 60
	}
	// snippet-end:[sqs.go.receive_messages.args]

	// Create a session that gets credential values from ~/.aws/credentials
	// and the default region from ~/.aws/config
	// snippet-start:[sqs.go.receive_messages.sess]
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "localstack",
		Config: aws.Config{
			// Region:   aws.String("us-east-1"),
			Endpoint: aws.String("http://localhost:4566"),
		},
	}))
	// snippet-end:[sqs.go.receive_messages.sess]

	// Get URL of queue
	urlResult, err := getQueueURL(sess, queue)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	// snippet-start:[sqs.go.receive_message.url]
	queueURL := urlResult.QueueUrl
	// snippet-end:[sqs.go.receive_message.url]

	// Send the message
	_, err = SendMessage(sess, queueURL, message)
	if err != nil {
		fmt.Println("Got an error sending the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully sent the message to the queue:", *queueURL)

}
