// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// WAF provides the API operation methods for making requests to
// AWS WAF. See this package's package overview docs
// for details on the service.
//
// WAF methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type WAF struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*WAF)

// Used for custom request initialization logic
var initRequest func(*WAF, *aws.Request)

// Service information constants
const (
	ServiceName = "waf"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the WAF client with a config.
// If additional configuration is needed for the client instance use the
// optional aws.Config parameter to add your extra config.
//
// Example:
//     // Create a WAF client from just a config.
//     svc := waf.New(myConfig)
//
//     // Create a WAF client with additional configuration
//     svc := waf.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *WAF {
	var signingName string
	signingRegion := config.Region

	svc := &WAF{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2015-08-24",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSWAF_20150824",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a WAF operation and runs any
// custom request initialization.
func (c *WAF) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
