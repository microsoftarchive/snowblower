package main

import (
	"os"
	"runtime"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/cobra"
)

var config struct {
	snsTopic      string
	sqsURL        string
	collectorPort string
	awsregion     string
	awsSession    *session.Session
}

func main() {

	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	config.collectorPort = os.Getenv("PORT")
	if config.collectorPort == "" {
		config.collectorPort = "8080"
	}

	config.snsTopic = os.Getenv("SNS_TOPIC")
	config.sqsURL = os.Getenv("SQS_URL")
	config.awsregion = os.Getenv("AWS_DEFAULT_REGION")

	config.awsSession = session.Must(session.NewSession())

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
