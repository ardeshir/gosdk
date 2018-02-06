// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package comprehendiface provides an interface to enable mocking the Amazon Comprehend service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package comprehendiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
)

// ComprehendAPI provides an interface to enable mocking the
// comprehend.Comprehend service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Comprehend.
//    func myFunc(svc comprehendiface.ComprehendAPI) bool {
//        // Make svc.BatchDetectDominantLanguage request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := comprehend.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockComprehendClient struct {
//        comprehendiface.ComprehendAPI
//    }
//    func (m *mockComprehendClient) BatchDetectDominantLanguage(input *comprehend.BatchDetectDominantLanguageInput) (*comprehend.BatchDetectDominantLanguageOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockComprehendClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ComprehendAPI interface {
	BatchDetectDominantLanguageRequest(*comprehend.BatchDetectDominantLanguageInput) comprehend.BatchDetectDominantLanguageRequest

	BatchDetectEntitiesRequest(*comprehend.BatchDetectEntitiesInput) comprehend.BatchDetectEntitiesRequest

	BatchDetectKeyPhrasesRequest(*comprehend.BatchDetectKeyPhrasesInput) comprehend.BatchDetectKeyPhrasesRequest

	BatchDetectSentimentRequest(*comprehend.BatchDetectSentimentInput) comprehend.BatchDetectSentimentRequest

	DescribeTopicsDetectionJobRequest(*comprehend.DescribeTopicsDetectionJobInput) comprehend.DescribeTopicsDetectionJobRequest

	DetectDominantLanguageRequest(*comprehend.DetectDominantLanguageInput) comprehend.DetectDominantLanguageRequest

	DetectEntitiesRequest(*comprehend.DetectEntitiesInput) comprehend.DetectEntitiesRequest

	DetectKeyPhrasesRequest(*comprehend.DetectKeyPhrasesInput) comprehend.DetectKeyPhrasesRequest

	DetectSentimentRequest(*comprehend.DetectSentimentInput) comprehend.DetectSentimentRequest

	ListTopicsDetectionJobsRequest(*comprehend.ListTopicsDetectionJobsInput) comprehend.ListTopicsDetectionJobsRequest

	ListTopicsDetectionJobsPages(*comprehend.ListTopicsDetectionJobsInput, func(*comprehend.ListTopicsDetectionJobsOutput, bool) bool) error
	ListTopicsDetectionJobsPagesWithContext(aws.Context, *comprehend.ListTopicsDetectionJobsInput, func(*comprehend.ListTopicsDetectionJobsOutput, bool) bool, ...aws.Option) error

	StartTopicsDetectionJobRequest(*comprehend.StartTopicsDetectionJobInput) comprehend.StartTopicsDetectionJobRequest
}

var _ ComprehendAPI = (*comprehend.Comprehend)(nil)
