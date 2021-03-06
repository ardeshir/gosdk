// +build integration

package performance

import (
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/elb"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics"
	"github.com/aws/aws-sdk-go-v2/service/mobileanalytics"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/simpledb"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/swf"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

var clients = []reflect.Value{
	reflect.ValueOf(acm.New),
	reflect.ValueOf(apigateway.New),
	reflect.ValueOf(autoscaling.New),
	reflect.ValueOf(cloudformation.New),
	reflect.ValueOf(cloudfront.New),
	reflect.ValueOf(cloudhsm.New),
	reflect.ValueOf(cloudsearch.New),
	reflect.ValueOf(cloudsearchdomain.New),
	reflect.ValueOf(cloudtrail.New),
	reflect.ValueOf(cloudwatch.New),
	reflect.ValueOf(cloudwatchevents.New),
	reflect.ValueOf(cloudwatchlogs.New),
	reflect.ValueOf(codecommit.New),
	reflect.ValueOf(codedeploy.New),
	reflect.ValueOf(codepipeline.New),
	reflect.ValueOf(cognitoidentity.New),
	reflect.ValueOf(cognitosync.New),
	reflect.ValueOf(configservice.New),
	reflect.ValueOf(datapipeline.New),
	reflect.ValueOf(devicefarm.New),
	reflect.ValueOf(directconnect.New),
	reflect.ValueOf(directoryservice.New),
	reflect.ValueOf(dynamodb.New),
	reflect.ValueOf(dynamodbstreams.New),
	reflect.ValueOf(ec2.New),
	reflect.ValueOf(ecr.New),
	reflect.ValueOf(ecs.New),
	reflect.ValueOf(efs.New),
	reflect.ValueOf(elasticache.New),
	reflect.ValueOf(elasticbeanstalk.New),
	reflect.ValueOf(elasticsearchservice.New),
	reflect.ValueOf(elastictranscoder.New),
	reflect.ValueOf(elb.New),
	reflect.ValueOf(emr.New),
	reflect.ValueOf(firehose.New),
	reflect.ValueOf(glacier.New),
	reflect.ValueOf(iam.New),
	reflect.ValueOf(inspector.New),
	reflect.ValueOf(iot.New),
	reflect.ValueOf(iotdataplane.New),
	reflect.ValueOf(kinesis.New),
	reflect.ValueOf(kms.New),
	reflect.ValueOf(lambda.New),
	reflect.ValueOf(machinelearning.New),
	reflect.ValueOf(marketplacecommerceanalytics.New),
	reflect.ValueOf(mobileanalytics.New),
	reflect.ValueOf(opsworks.New),
	reflect.ValueOf(rds.New),
	reflect.ValueOf(redshift.New),
	reflect.ValueOf(route53.New),
	reflect.ValueOf(route53domains.New),
	reflect.ValueOf(s3.New),
	reflect.ValueOf(ses.New),
	reflect.ValueOf(simpledb.New),
	reflect.ValueOf(sns.New),
	reflect.ValueOf(sqs.New),
	reflect.ValueOf(ssm.New),
	reflect.ValueOf(storagegateway.New),
	reflect.ValueOf(sts.New),
	reflect.ValueOf(support.New),
	reflect.ValueOf(swf.New),
	reflect.ValueOf(waf.New),
	reflect.ValueOf(workspaces.New),
}
