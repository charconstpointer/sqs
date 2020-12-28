package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	queueURL = flag.String("q", "", "The name of the queue")
)

func main() {
	flag.Parse()
	if *queueURL == "" {
		fmt.Println("You must supply the name of a queue (-q QUEUE)")
		return
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
	s := bufio.NewScanner(os.Stdin)
	for {
		log.Println("type anything to fetch messages")
		s.Scan()
		msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			AttributeNames: []*string{
				aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
			},
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            queueURL,
			MaxNumberOfMessages: aws.Int64(1),
			VisibilityTimeout:   aws.Int64(123),
		})
		msgCount := len(msgResult.Messages)
		if msgCount == 0 {
			log.Println("no messages to fetch")
			continue
		}
		for i, msg := range msgResult.Messages {
			log.Printf("%d %s", i, *msg.Body)
		}
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
