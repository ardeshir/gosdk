// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// APIGateway provides the API operation methods for making requests to
// Amazon API Gateway. See this package's package overview docs
// for details on the service.
//
// APIGateway methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type APIGateway struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*APIGateway)

// Used for custom request initialization logic
var initRequest func(*APIGateway, *aws.Request)

// Service information constants
const (
	ServiceName = "apigateway" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName  // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the APIGateway client with a config.
// If additional configuration is needed for the client instance use the
// optional aws.Config parameter to add your extra config.
//
// Example:
//     // Create a APIGateway client from just a config.
//     svc := apigateway.New(myConfig)
//
//     // Create a APIGateway client with additional configuration
//     svc := apigateway.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *APIGateway {
	var signingName string
	signingRegion := config.Region

	svc := &APIGateway{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2015-07-09",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a APIGateway operation and runs any
// custom request initialization.
func (c *APIGateway) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
