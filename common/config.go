package common

import(
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/sns"
)

type Config struct {
	Credentials   		aws.CredentialsProvider
	SnsService    		*sns.SNS
	CollectorPort 		string
	CollectedSnsTopic   string
	CollectedSqsURL     string
	EnrichedSnsTopic	string
	EnrichedSqsURL 		string
}