// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pricing

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

const opDescribeServices = "DescribeServices"

// DescribeServicesRequest is a API request type for the DescribeServices API operation.
type DescribeServicesRequest struct {
	*aws.Request
	Input *DescribeServicesInput
}

// Send marshals and sends the DescribeServices API request.
func (r DescribeServicesRequest) Send() (*DescribeServicesOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*DescribeServicesOutput), nil
}

// DescribeServicesRequest returns a request value for making API operation for
// AWS Price List Service.
//
// Returns the metadata for one service or a list of the metadata for all services.
// Use this without a service code to get the service codes for all services.
// Use it with a service code, such as AmazonEC2, to get information specific
// to that service, such as the attribute names available for that service.
// For example, some of the attribute names available for EC2 are volumeType,
// maxIopsVolume, operation, locationType, and instanceCapacity10xlarge.
//
//    // Example sending a request using the DescribeServicesRequest method.
//    req := client.DescribeServicesRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/DescribeServices
func (c *Pricing) DescribeServicesRequest(input *DescribeServicesInput) DescribeServicesRequest {
	op := &aws.Operation{
		Name:       opDescribeServices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeServicesInput{}
	}

	output := &DescribeServicesOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return DescribeServicesRequest{Request: req, Input: input}
}

// DescribeServicesPages iterates over the pages of a DescribeServices operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See DescribeServices method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a DescribeServices operation.
//    pageNum := 0
//    err := client.DescribeServicesPages(params,
//        func(page *DescribeServicesOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *Pricing) DescribeServicesPages(input *DescribeServicesInput, fn func(*DescribeServicesOutput, bool) bool) error {
	return c.DescribeServicesPagesWithContext(aws.BackgroundContext(), input, fn)
}

// DescribeServicesPagesWithContext same as DescribeServicesPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Pricing) DescribeServicesPagesWithContext(ctx aws.Context, input *DescribeServicesInput, fn func(*DescribeServicesOutput, bool) bool, opts ...aws.Option) error {
	p := aws.Pagination{
		NewRequest: func() (*aws.Request, error) {
			var inCpy *DescribeServicesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeServicesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*DescribeServicesOutput), !p.HasNextPage())
	}
	return p.Err()
}

const opGetAttributeValues = "GetAttributeValues"

// GetAttributeValuesRequest is a API request type for the GetAttributeValues API operation.
type GetAttributeValuesRequest struct {
	*aws.Request
	Input *GetAttributeValuesInput
}

// Send marshals and sends the GetAttributeValues API request.
func (r GetAttributeValuesRequest) Send() (*GetAttributeValuesOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetAttributeValuesOutput), nil
}

// GetAttributeValuesRequest returns a request value for making API operation for
// AWS Price List Service.
//
// Returns a list of attribute values. Attibutes are similar to the details
// in a Price List API offer file. For a list of available attributes, see Offer
// File Definitions (http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/reading-an-offer.html#pps-defs)
// in the AWS Billing and Cost Management User Guide (http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-what-is.html).
//
//    // Example sending a request using the GetAttributeValuesRequest method.
//    req := client.GetAttributeValuesRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/GetAttributeValues
func (c *Pricing) GetAttributeValuesRequest(input *GetAttributeValuesInput) GetAttributeValuesRequest {
	op := &aws.Operation{
		Name:       opGetAttributeValues,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetAttributeValuesInput{}
	}

	output := &GetAttributeValuesOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetAttributeValuesRequest{Request: req, Input: input}
}

// GetAttributeValuesPages iterates over the pages of a GetAttributeValues operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See GetAttributeValues method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a GetAttributeValues operation.
//    pageNum := 0
//    err := client.GetAttributeValuesPages(params,
//        func(page *GetAttributeValuesOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *Pricing) GetAttributeValuesPages(input *GetAttributeValuesInput, fn func(*GetAttributeValuesOutput, bool) bool) error {
	return c.GetAttributeValuesPagesWithContext(aws.BackgroundContext(), input, fn)
}

// GetAttributeValuesPagesWithContext same as GetAttributeValuesPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Pricing) GetAttributeValuesPagesWithContext(ctx aws.Context, input *GetAttributeValuesInput, fn func(*GetAttributeValuesOutput, bool) bool, opts ...aws.Option) error {
	p := aws.Pagination{
		NewRequest: func() (*aws.Request, error) {
			var inCpy *GetAttributeValuesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.GetAttributeValuesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*GetAttributeValuesOutput), !p.HasNextPage())
	}
	return p.Err()
}

const opGetProducts = "GetProducts"

// GetProductsRequest is a API request type for the GetProducts API operation.
type GetProductsRequest struct {
	*aws.Request
	Input *GetProductsInput
}

// Send marshals and sends the GetProducts API request.
func (r GetProductsRequest) Send() (*GetProductsOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*GetProductsOutput), nil
}

// GetProductsRequest returns a request value for making API operation for
// AWS Price List Service.
//
// Returns a list of all products that match the filter criteria.
//
//    // Example sending a request using the GetProductsRequest method.
//    req := client.GetProductsRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/GetProducts
func (c *Pricing) GetProductsRequest(input *GetProductsInput) GetProductsRequest {
	op := &aws.Operation{
		Name:       opGetProducts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetProductsInput{}
	}

	output := &GetProductsOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return GetProductsRequest{Request: req, Input: input}
}

// GetProductsPages iterates over the pages of a GetProducts operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See GetProducts method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a GetProducts operation.
//    pageNum := 0
//    err := client.GetProductsPages(params,
//        func(page *GetProductsOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *Pricing) GetProductsPages(input *GetProductsInput, fn func(*GetProductsOutput, bool) bool) error {
	return c.GetProductsPagesWithContext(aws.BackgroundContext(), input, fn)
}

// GetProductsPagesWithContext same as GetProductsPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Pricing) GetProductsPagesWithContext(ctx aws.Context, input *GetProductsInput, fn func(*GetProductsOutput, bool) bool, opts ...aws.Option) error {
	p := aws.Pagination{
		NewRequest: func() (*aws.Request, error) {
			var inCpy *GetProductsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.GetProductsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*GetProductsOutput), !p.HasNextPage())
	}
	return p.Err()
}

// The values of a given attribute, such as Throughput Optimized HDD or Provisioned
// IOPS for the Amazon EC2volumeType attribute.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/AttributeValue
type AttributeValue struct {
	_ struct{} `type:"structure"`

	// The specific value of an attributeName.
	Value *string `type:"string"`
}

// String returns the string representation
func (s AttributeValue) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AttributeValue) GoString() string {
	return s.String()
}

// SetValue sets the Value field's value.
func (s *AttributeValue) SetValue(v string) *AttributeValue {
	s.Value = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/DescribeServicesRequest
type DescribeServicesInput struct {
	_ struct{} `type:"structure"`

	// The format version that you want the response to be in.
	//
	// Valid values are: aws_v1
	FormatVersion *string `type:"string"`

	// The maximum number of results that you want returned in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token that indicates the next set of results that you want
	// to retrieve.
	NextToken *string `type:"string"`

	// The code for the service whose information you want to retrieve, such as
	// AmazonEC2. You can use the ServiceCode to filter the results in a GetProducts
	// call. To retrieve a list of all services, leave this blank.
	ServiceCode *string `type:"string"`
}

// String returns the string representation
func (s DescribeServicesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeServicesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeServicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeServicesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFormatVersion sets the FormatVersion field's value.
func (s *DescribeServicesInput) SetFormatVersion(v string) *DescribeServicesInput {
	s.FormatVersion = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeServicesInput) SetMaxResults(v int64) *DescribeServicesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeServicesInput) SetNextToken(v string) *DescribeServicesInput {
	s.NextToken = &v
	return s
}

// SetServiceCode sets the ServiceCode field's value.
func (s *DescribeServicesInput) SetServiceCode(v string) *DescribeServicesInput {
	s.ServiceCode = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/DescribeServicesResponse
type DescribeServicesOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The format version of the response. For example, aws_v1.
	FormatVersion *string `type:"string"`

	// The pagination token for the next set of retreivable results.
	NextToken *string `type:"string"`

	// The service metadata for the service or services in the response.
	Services []Service `type:"list"`
}

// String returns the string representation
func (s DescribeServicesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeServicesOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s DescribeServicesOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// SetFormatVersion sets the FormatVersion field's value.
func (s *DescribeServicesOutput) SetFormatVersion(v string) *DescribeServicesOutput {
	s.FormatVersion = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeServicesOutput) SetNextToken(v string) *DescribeServicesOutput {
	s.NextToken = &v
	return s
}

// SetServices sets the Services field's value.
func (s *DescribeServicesOutput) SetServices(v []Service) *DescribeServicesOutput {
	s.Services = v
	return s
}

// The constraints that you want all returned products to match.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/Filter
type Filter struct {
	_ struct{} `type:"structure"`

	// The product metadata field that you want to filter on. You can filter by
	// just the service code to see all products for a specific service, filter
	// by just the attribute name to see a specific attribute for multiple services,
	// or use both a service code and an attribute name to retrieve only products
	// that match both fields.
	//
	// Valid values include: ServiceCode, and all attribute names
	//
	// For example, you can filter by the AmazonEC2 service code and the volumeType
	// attribute name to get the prices for only Amazon EC2 volumes.
	//
	// Field is a required field
	Field *string `type:"string" required:"true"`

	// The type of filter that you want to use.
	//
	// Valid values are: TERM_MATCH. TERM_MATCH returns only products that match
	// both the given filter field and the given value.
	//
	// Type is a required field
	Type FilterType `type:"string" required:"true" enum:"true"`

	// The service code or attribute value that you want to filter by. If you are
	// filtering by service code this is the actual service code, such as AmazonEC2.
	// If you are filtering by attribute name, this is the attribute value that
	// you want the returned products to match, such as a Provisioned IOPS volume.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Filter) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Filter) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Filter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Filter"}

	if s.Field == nil {
		invalidParams.Add(aws.NewErrParamRequired("Field"))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetField sets the Field field's value.
func (s *Filter) SetField(v string) *Filter {
	s.Field = &v
	return s
}

// SetType sets the Type field's value.
func (s *Filter) SetType(v FilterType) *Filter {
	s.Type = v
	return s
}

// SetValue sets the Value field's value.
func (s *Filter) SetValue(v string) *Filter {
	s.Value = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/GetAttributeValuesRequest
type GetAttributeValuesInput struct {
	_ struct{} `type:"structure"`

	// The name of the attribute that you want to retrieve the values for, such
	// as volumeType.
	//
	// AttributeName is a required field
	AttributeName *string `type:"string" required:"true"`

	// The maximum number of results to return in response.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token that indicates the next set of results that you want
	// to retrieve.
	NextToken *string `type:"string"`

	// The service code for the service whose attributes you want to retrieve. For
	// example, if you want the retrieve an EC2 attribute, use AmazonEC2.
	//
	// ServiceCode is a required field
	ServiceCode *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetAttributeValuesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAttributeValuesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAttributeValuesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAttributeValuesInput"}

	if s.AttributeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttributeName"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ServiceCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceCode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributeName sets the AttributeName field's value.
func (s *GetAttributeValuesInput) SetAttributeName(v string) *GetAttributeValuesInput {
	s.AttributeName = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *GetAttributeValuesInput) SetMaxResults(v int64) *GetAttributeValuesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetAttributeValuesInput) SetNextToken(v string) *GetAttributeValuesInput {
	s.NextToken = &v
	return s
}

// SetServiceCode sets the ServiceCode field's value.
func (s *GetAttributeValuesInput) SetServiceCode(v string) *GetAttributeValuesInput {
	s.ServiceCode = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/GetAttributeValuesResponse
type GetAttributeValuesOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The list of values for an attribute. For example, Throughput Optimized HDD
	// and Provisioned IOPS are two available values for the AmazonEC2volumeType.
	AttributeValues []AttributeValue `type:"list"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetAttributeValuesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAttributeValuesOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetAttributeValuesOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// SetAttributeValues sets the AttributeValues field's value.
func (s *GetAttributeValuesOutput) SetAttributeValues(v []AttributeValue) *GetAttributeValuesOutput {
	s.AttributeValues = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetAttributeValuesOutput) SetNextToken(v string) *GetAttributeValuesOutput {
	s.NextToken = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/GetProductsRequest
type GetProductsInput struct {
	_ struct{} `type:"structure"`

	// The list of filters that limit the returned products. only products that
	// match all filters are returned.
	Filters []Filter `type:"list"`

	// The format version that you want the response to be in.
	//
	// Valid values are: aws_v1
	FormatVersion *string `type:"string"`

	// The maximum number of results to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token that indicates the next set of results that you want
	// to retrieve.
	NextToken *string `type:"string"`

	// The code for the service whose products you want to retrieve.
	ServiceCode *string `type:"string"`
}

// String returns the string representation
func (s GetProductsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetProductsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetProductsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetProductsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilters sets the Filters field's value.
func (s *GetProductsInput) SetFilters(v []Filter) *GetProductsInput {
	s.Filters = v
	return s
}

// SetFormatVersion sets the FormatVersion field's value.
func (s *GetProductsInput) SetFormatVersion(v string) *GetProductsInput {
	s.FormatVersion = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *GetProductsInput) SetMaxResults(v int64) *GetProductsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetProductsInput) SetNextToken(v string) *GetProductsInput {
	s.NextToken = &v
	return s
}

// SetServiceCode sets the ServiceCode field's value.
func (s *GetProductsInput) SetServiceCode(v string) *GetProductsInput {
	s.ServiceCode = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/GetProductsResponse
type GetProductsOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// The format version of the response. For example, aws_v1.
	FormatVersion *string `type:"string"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `type:"string"`

	// The list of products that match your filters. The list contains both the
	// product metadata and the price information.
	PriceList []aws.JSONValue `type:"list"`
}

// String returns the string representation
func (s GetProductsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetProductsOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s GetProductsOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// SetFormatVersion sets the FormatVersion field's value.
func (s *GetProductsOutput) SetFormatVersion(v string) *GetProductsOutput {
	s.FormatVersion = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetProductsOutput) SetNextToken(v string) *GetProductsOutput {
	s.NextToken = &v
	return s
}

// SetPriceList sets the PriceList field's value.
func (s *GetProductsOutput) SetPriceList(v []aws.JSONValue) *GetProductsOutput {
	s.PriceList = v
	return s
}

// The metadata for a service, such as the service code and available attribute
// names.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15/Service
type Service struct {
	_ struct{} `type:"structure"`

	// The attributes that are available for this service.
	AttributeNames []string `type:"list"`

	// The code for the AWS service.
	ServiceCode *string `type:"string"`
}

// String returns the string representation
func (s Service) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Service) GoString() string {
	return s.String()
}

// SetAttributeNames sets the AttributeNames field's value.
func (s *Service) SetAttributeNames(v []string) *Service {
	s.AttributeNames = v
	return s
}

// SetServiceCode sets the ServiceCode field's value.
func (s *Service) SetServiceCode(v string) *Service {
	s.ServiceCode = &v
	return s
}

type FilterType string

// Enum values for FilterType
const (
	FilterTypeTermMatch FilterType = "TERM_MATCH"
)
