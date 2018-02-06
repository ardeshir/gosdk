// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// AlexaForBusiness provides the API operation methods for making requests to
// Alexa For Business. See this package's package overview docs
// for details on the service.
//
// AlexaForBusiness methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type AlexaForBusiness struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*AlexaForBusiness)

// Used for custom request initialization logic
var initRequest func(*AlexaForBusiness, *aws.Request)

// Service information constants
const (
	ServiceName = "a4b"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the AlexaForBusiness client with a config.
// If additional configuration is needed for the client instance use the
// optional aws.Config parameter to add your extra config.
//
// Example:
//     // Create a AlexaForBusiness client from just a config.
//     svc := alexaforbusiness.New(myConfig)
//
//     // Create a AlexaForBusiness client with additional configuration
//     svc := alexaforbusiness.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *AlexaForBusiness {
	var signingName string
	signingRegion := config.Region

	svc := &AlexaForBusiness{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2017-11-09",
				JSONVersion:   "1.1",
				TargetPrefix:  "AlexaForBusiness",
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

// newRequest creates a new request for a AlexaForBusiness operation and runs any
// custom request initialization.
func (c *AlexaForBusiness) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
