package main

import (
	"os"
	"runtime"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/sns"
	"github.com/spf13/cobra"
)

var config struct {
	credentials   aws.CredentialsProvider
	snsTopic      string
	snsService    *sns.SNS
	sqsURL        string
	collectorPort string
}

func main() {

	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	config.collectorPort = os.Getenv("PORT")
	if config.collectorPort == "" {
		config.collectorPort = "8080"
	}

	//var credentials aws.CredentialsProvider
	if os.Getenv("AWS_ACCESS_KEY_ID") != "" {
		config.credentials = aws.DefaultCreds()
	} else {
		config.credentials = aws.IAMCreds()
	}

	config.snsTopic = os.Getenv("SNS_TOPIC")
	config.sqsURL = os.Getenv("SQS_URL")

	config.snsService = sns.New(&aws.Config{
		Credentials: config.credentials,
		Region:      "eu-west-1",
	})

	var collectorCmd = &cobra.Command{
		Use:   "collect",
		Short: "Run the collector",
		Run: func(cmd *cobra.Command, args []string) {
			if config.snsTopic == "" {
				panic("SNS_TOPIC required")
			}
			startCollector()
		},
	}

	var etlCmd = &cobra.Command{
		Use:   "etl",
		Short: "Run the ETL processor",
		Run: func(cd *cobra.Command, args []string) {
			if config.sqsURL == "" {
				panic("SQS_URL required")
			}
			// ensure we have database information here
			startETL()
		},
	}

	var rootCmd = &cobra.Command{Use: "snowblower"}
	rootCmd.AddCommand(collectorCmd)
	rootCmd.AddCommand(etlCmd)
	rootCmd.Execute()

}
