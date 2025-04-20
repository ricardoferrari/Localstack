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

// snippet-end:[sqs.go.receive_messages.imports]

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
func GetQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
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

// GetMessages gets the messages from an Amazon SQS queue
// Inputs:
//
//	sess is the current session, which provides configuration for the SDK's service clients
//	queueURL is the URL of the queue
//	timeout is how long, in seconds, the message is unavailable to other consumers
//
// Output:
//
//	If success, the latest message and nil
//	Otherwise, nil and an error from the call to ReceiveMessage
func GetMessages(sess *session.Session, queueURL *string, timeout *int64) (*sqs.ReceiveMessageOutput, error) {
	// Create an SQS service client
	svc := sqs.New(sess)

	// snippet-start:[sqs.go.receive_messages.call]
	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   timeout,
	})
	// snippet-end:[sqs.go.receive_messages.call]
	if err != nil {
		return nil, err
	}

	return msgResult, nil
}

func DeleteMessage(sess *session.Session, queueURL *string, handle *string) (string, error) {
	svc := sqs.New(sess)
	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: handle,
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
	flag.Parse()

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
	urlResult, err := GetQueueURL(sess, queue)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	// snippet-start:[sqs.go.receive_message.url]
	queueURL := urlResult.QueueUrl
	// snippet-end:[sqs.go.receive_message.url]

	msgResult, err := GetMessages(sess, queueURL, timeout)
	if err != nil {
		fmt.Println("Got an error receiving messages:")
		fmt.Println(err)
		return
	}
	if len(msgResult.Messages) == 0 {
		fmt.Println("No messages received")
		return
	}
	fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
	fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
	fmt.Println("Message Body:   " + *msgResult.Messages[0].Body)

	fmt.Println("Acknowledging the received message...")
	// Acknowledge the message by deleting it from the queue
	_, err = DeleteMessage(sess, queueURL, msgResult.Messages[0].ReceiptHandle)
	if err != nil {
		fmt.Println("Got an error deleting the message:")
		fmt.Println(err)
		return
	}
	fmt.Println("Deleted the message")

}

// snippet-end:[sqs.go.receive_messages]
