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
		fmt.Print("msg >")
		s.Scan()
		body := s.Text()
		_, err := svc.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(0),
			MessageBody:  aws.String(body),
			QueueUrl:     queueURL,
		})

		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
