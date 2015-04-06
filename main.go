package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/sns"
)

//var snsService *sns.SNS
//var snsTopic string

func main() {

	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	var credentials aws.CredentialsProvider
	if os.Getenv("AWS_ACCESS_KEY_ID") != "" {
		credentials = aws.DefaultCreds()
	} else {
		credentials = aws.IAMCreds()
	}

	snsTopic := os.Getenv("SNS_TOPIC")
	if snsTopic == "" {
		panic("SNS_TOPIC required")
	}

	snsService := sns.New(&aws.Config{
		Credentials: credentials,
		Region:      "eu-west-1",
	})

	println(snsService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	healthHandler := &health{}
	http.Handle("/api/health", healthHandler)
	http.Handle("/health", healthHandler)

	collectorHandler := &collector{
		publisher: &snsPublisher{
			service: snsService,
			topic:   snsTopic,
		},
	}

	http.Handle("/com.snowplowanalytics.snowplow/tp2", collectorHandler)
	http.Handle("/i", collectorHandler)

	portString := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on %s", portString)
	http.ListenAndServe(portString, nil)

}
