package main

import (
	"os"
	"runtime"

	"github.com/wunderlist/snowblower/common"
	"github.com/wunderlist/snowblower/collector"
	"github.com/wunderlist/snowblower/enricher"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/sns"
	"github.com/spf13/cobra"
)

func main() {

	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	// build configuration
	config := common.Config{
		CollectorPort: 		os.Getenv("COLLECTOR_PORT"),
		CollectedSnsTopic: 	os.Getenv("COLLECTED_SNS_TOPIC"),
		CollectedSqsURL:	os.Getenv("COLLECTED_SQS_URL"),
		EnrichedSnsTopic: 	os.Getenv("ENRICHED_SNS_TOPIC"),
		EnrichedSqsURL:		os.Getenv("ENRICHED_SQS_URL"),	
	}

	// default collector port
	if config.CollectorPort == "" {
		config.CollectorPort = "8080"
	}

	// aws credentials
	if os.Getenv("AWS_ACCESS_KEY_ID") != "" {
		config.Credentials = aws.DefaultCreds()
	} else {
		config.Credentials = aws.IAMCreds()
	}

	// sns service
	config.SnsService = sns.New(&aws.Config{
		Credentials: config.Credentials,
		Region:      "eu-west-1",
	})


	// build collector command
	var collectorCmd = &cobra.Command{
		Use:   "collect",
		Short: "Run the collector",
		Run: func(cmd *cobra.Command, args []string) {
			if config.CollectedSnsTopic == "" {
				panic("COLLECTED_SNS_TOPIC required")
			}
			collector.Start(config)
		},
	}


	// build enricher command
	var enricherCmd = &cobra.Command{
		Use:   "enrich",
		Short: "Run the enricher",
		Run: func(cd *cobra.Command, args []string) {
			if config.CollectedSqsURL == "" {
				panic("COLLECTED_SQS_URL required")
			}
			if config.EnrichedSqsURL == "" {
				panic("ENRICHED_SQS_URL required")
			}
			// ensure we have database information here
			enricher.Start(config)
		},
	}


	// build main command
	var rootCmd = &cobra.Command{Use: "snowblower"}
	rootCmd.AddCommand(collectorCmd)
	rootCmd.AddCommand(enricherCmd)
	rootCmd.Execute()
}


