// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package glueiface provides an interface to enable mocking the AWS Glue service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package glueiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
)

// GlueAPI provides an interface to enable mocking the
// glue.Glue service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Glue.
//    func myFunc(svc glueiface.GlueAPI) bool {
//        // Make svc.BatchCreatePartition request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := glue.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockGlueClient struct {
//        glueiface.GlueAPI
//    }
//    func (m *mockGlueClient) BatchCreatePartition(input *glue.BatchCreatePartitionInput) (*glue.BatchCreatePartitionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockGlueClient{}
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
type GlueAPI interface {
	BatchCreatePartitionRequest(*glue.BatchCreatePartitionInput) glue.BatchCreatePartitionRequest

	BatchDeleteConnectionRequest(*glue.BatchDeleteConnectionInput) glue.BatchDeleteConnectionRequest

	BatchDeletePartitionRequest(*glue.BatchDeletePartitionInput) glue.BatchDeletePartitionRequest

	BatchDeleteTableRequest(*glue.BatchDeleteTableInput) glue.BatchDeleteTableRequest

	BatchGetPartitionRequest(*glue.BatchGetPartitionInput) glue.BatchGetPartitionRequest

	BatchStopJobRunRequest(*glue.BatchStopJobRunInput) glue.BatchStopJobRunRequest

	CreateClassifierRequest(*glue.CreateClassifierInput) glue.CreateClassifierRequest

	CreateConnectionRequest(*glue.CreateConnectionInput) glue.CreateConnectionRequest

	CreateCrawlerRequest(*glue.CreateCrawlerInput) glue.CreateCrawlerRequest

	CreateDatabaseRequest(*glue.CreateDatabaseInput) glue.CreateDatabaseRequest

	CreateDevEndpointRequest(*glue.CreateDevEndpointInput) glue.CreateDevEndpointRequest

	CreateJobRequest(*glue.CreateJobInput) glue.CreateJobRequest

	CreatePartitionRequest(*glue.CreatePartitionInput) glue.CreatePartitionRequest

	CreateScriptRequest(*glue.CreateScriptInput) glue.CreateScriptRequest

	CreateTableRequest(*glue.CreateTableInput) glue.CreateTableRequest

	CreateTriggerRequest(*glue.CreateTriggerInput) glue.CreateTriggerRequest

	CreateUserDefinedFunctionRequest(*glue.CreateUserDefinedFunctionInput) glue.CreateUserDefinedFunctionRequest

	DeleteClassifierRequest(*glue.DeleteClassifierInput) glue.DeleteClassifierRequest

	DeleteConnectionRequest(*glue.DeleteConnectionInput) glue.DeleteConnectionRequest

	DeleteCrawlerRequest(*glue.DeleteCrawlerInput) glue.DeleteCrawlerRequest

	DeleteDatabaseRequest(*glue.DeleteDatabaseInput) glue.DeleteDatabaseRequest

	DeleteDevEndpointRequest(*glue.DeleteDevEndpointInput) glue.DeleteDevEndpointRequest

	DeleteJobRequest(*glue.DeleteJobInput) glue.DeleteJobRequest

	DeletePartitionRequest(*glue.DeletePartitionInput) glue.DeletePartitionRequest

	DeleteTableRequest(*glue.DeleteTableInput) glue.DeleteTableRequest

	DeleteTriggerRequest(*glue.DeleteTriggerInput) glue.DeleteTriggerRequest

	DeleteUserDefinedFunctionRequest(*glue.DeleteUserDefinedFunctionInput) glue.DeleteUserDefinedFunctionRequest

	GetCatalogImportStatusRequest(*glue.GetCatalogImportStatusInput) glue.GetCatalogImportStatusRequest

	GetClassifierRequest(*glue.GetClassifierInput) glue.GetClassifierRequest

	GetClassifiersRequest(*glue.GetClassifiersInput) glue.GetClassifiersRequest

	GetClassifiersPages(*glue.GetClassifiersInput, func(*glue.GetClassifiersOutput, bool) bool) error
	GetClassifiersPagesWithContext(aws.Context, *glue.GetClassifiersInput, func(*glue.GetClassifiersOutput, bool) bool, ...aws.Option) error

	GetConnectionRequest(*glue.GetConnectionInput) glue.GetConnectionRequest

	GetConnectionsRequest(*glue.GetConnectionsInput) glue.GetConnectionsRequest

	GetConnectionsPages(*glue.GetConnectionsInput, func(*glue.GetConnectionsOutput, bool) bool) error
	GetConnectionsPagesWithContext(aws.Context, *glue.GetConnectionsInput, func(*glue.GetConnectionsOutput, bool) bool, ...aws.Option) error

	GetCrawlerRequest(*glue.GetCrawlerInput) glue.GetCrawlerRequest

	GetCrawlerMetricsRequest(*glue.GetCrawlerMetricsInput) glue.GetCrawlerMetricsRequest

	GetCrawlerMetricsPages(*glue.GetCrawlerMetricsInput, func(*glue.GetCrawlerMetricsOutput, bool) bool) error
	GetCrawlerMetricsPagesWithContext(aws.Context, *glue.GetCrawlerMetricsInput, func(*glue.GetCrawlerMetricsOutput, bool) bool, ...aws.Option) error

	GetCrawlersRequest(*glue.GetCrawlersInput) glue.GetCrawlersRequest

	GetCrawlersPages(*glue.GetCrawlersInput, func(*glue.GetCrawlersOutput, bool) bool) error
	GetCrawlersPagesWithContext(aws.Context, *glue.GetCrawlersInput, func(*glue.GetCrawlersOutput, bool) bool, ...aws.Option) error

	GetDatabaseRequest(*glue.GetDatabaseInput) glue.GetDatabaseRequest

	GetDatabasesRequest(*glue.GetDatabasesInput) glue.GetDatabasesRequest

	GetDatabasesPages(*glue.GetDatabasesInput, func(*glue.GetDatabasesOutput, bool) bool) error
	GetDatabasesPagesWithContext(aws.Context, *glue.GetDatabasesInput, func(*glue.GetDatabasesOutput, bool) bool, ...aws.Option) error

	GetDataflowGraphRequest(*glue.GetDataflowGraphInput) glue.GetDataflowGraphRequest

	GetDevEndpointRequest(*glue.GetDevEndpointInput) glue.GetDevEndpointRequest

	GetDevEndpointsRequest(*glue.GetDevEndpointsInput) glue.GetDevEndpointsRequest

	GetDevEndpointsPages(*glue.GetDevEndpointsInput, func(*glue.GetDevEndpointsOutput, bool) bool) error
	GetDevEndpointsPagesWithContext(aws.Context, *glue.GetDevEndpointsInput, func(*glue.GetDevEndpointsOutput, bool) bool, ...aws.Option) error

	GetJobRequest(*glue.GetJobInput) glue.GetJobRequest

	GetJobRunRequest(*glue.GetJobRunInput) glue.GetJobRunRequest

	GetJobRunsRequest(*glue.GetJobRunsInput) glue.GetJobRunsRequest

	GetJobRunsPages(*glue.GetJobRunsInput, func(*glue.GetJobRunsOutput, bool) bool) error
	GetJobRunsPagesWithContext(aws.Context, *glue.GetJobRunsInput, func(*glue.GetJobRunsOutput, bool) bool, ...aws.Option) error

	GetJobsRequest(*glue.GetJobsInput) glue.GetJobsRequest

	GetJobsPages(*glue.GetJobsInput, func(*glue.GetJobsOutput, bool) bool) error
	GetJobsPagesWithContext(aws.Context, *glue.GetJobsInput, func(*glue.GetJobsOutput, bool) bool, ...aws.Option) error

	GetMappingRequest(*glue.GetMappingInput) glue.GetMappingRequest

	GetPartitionRequest(*glue.GetPartitionInput) glue.GetPartitionRequest

	GetPartitionsRequest(*glue.GetPartitionsInput) glue.GetPartitionsRequest

	GetPartitionsPages(*glue.GetPartitionsInput, func(*glue.GetPartitionsOutput, bool) bool) error
	GetPartitionsPagesWithContext(aws.Context, *glue.GetPartitionsInput, func(*glue.GetPartitionsOutput, bool) bool, ...aws.Option) error

	GetPlanRequest(*glue.GetPlanInput) glue.GetPlanRequest

	GetTableRequest(*glue.GetTableInput) glue.GetTableRequest

	GetTableVersionsRequest(*glue.GetTableVersionsInput) glue.GetTableVersionsRequest

	GetTableVersionsPages(*glue.GetTableVersionsInput, func(*glue.GetTableVersionsOutput, bool) bool) error
	GetTableVersionsPagesWithContext(aws.Context, *glue.GetTableVersionsInput, func(*glue.GetTableVersionsOutput, bool) bool, ...aws.Option) error

	GetTablesRequest(*glue.GetTablesInput) glue.GetTablesRequest

	GetTablesPages(*glue.GetTablesInput, func(*glue.GetTablesOutput, bool) bool) error
	GetTablesPagesWithContext(aws.Context, *glue.GetTablesInput, func(*glue.GetTablesOutput, bool) bool, ...aws.Option) error

	GetTriggerRequest(*glue.GetTriggerInput) glue.GetTriggerRequest

	GetTriggersRequest(*glue.GetTriggersInput) glue.GetTriggersRequest

	GetTriggersPages(*glue.GetTriggersInput, func(*glue.GetTriggersOutput, bool) bool) error
	GetTriggersPagesWithContext(aws.Context, *glue.GetTriggersInput, func(*glue.GetTriggersOutput, bool) bool, ...aws.Option) error

	GetUserDefinedFunctionRequest(*glue.GetUserDefinedFunctionInput) glue.GetUserDefinedFunctionRequest

	GetUserDefinedFunctionsRequest(*glue.GetUserDefinedFunctionsInput) glue.GetUserDefinedFunctionsRequest

	GetUserDefinedFunctionsPages(*glue.GetUserDefinedFunctionsInput, func(*glue.GetUserDefinedFunctionsOutput, bool) bool) error
	GetUserDefinedFunctionsPagesWithContext(aws.Context, *glue.GetUserDefinedFunctionsInput, func(*glue.GetUserDefinedFunctionsOutput, bool) bool, ...aws.Option) error

	ImportCatalogToGlueRequest(*glue.ImportCatalogToGlueInput) glue.ImportCatalogToGlueRequest

	ResetJobBookmarkRequest(*glue.ResetJobBookmarkInput) glue.ResetJobBookmarkRequest

	StartCrawlerRequest(*glue.StartCrawlerInput) glue.StartCrawlerRequest

	StartCrawlerScheduleRequest(*glue.StartCrawlerScheduleInput) glue.StartCrawlerScheduleRequest

	StartJobRunRequest(*glue.StartJobRunInput) glue.StartJobRunRequest

	StartTriggerRequest(*glue.StartTriggerInput) glue.StartTriggerRequest

	StopCrawlerRequest(*glue.StopCrawlerInput) glue.StopCrawlerRequest

	StopCrawlerScheduleRequest(*glue.StopCrawlerScheduleInput) glue.StopCrawlerScheduleRequest

	StopTriggerRequest(*glue.StopTriggerInput) glue.StopTriggerRequest

	UpdateClassifierRequest(*glue.UpdateClassifierInput) glue.UpdateClassifierRequest

	UpdateConnectionRequest(*glue.UpdateConnectionInput) glue.UpdateConnectionRequest

	UpdateCrawlerRequest(*glue.UpdateCrawlerInput) glue.UpdateCrawlerRequest

	UpdateCrawlerScheduleRequest(*glue.UpdateCrawlerScheduleInput) glue.UpdateCrawlerScheduleRequest

	UpdateDatabaseRequest(*glue.UpdateDatabaseInput) glue.UpdateDatabaseRequest

	UpdateDevEndpointRequest(*glue.UpdateDevEndpointInput) glue.UpdateDevEndpointRequest

	UpdateJobRequest(*glue.UpdateJobInput) glue.UpdateJobRequest

	UpdatePartitionRequest(*glue.UpdatePartitionInput) glue.UpdatePartitionRequest

	UpdateTableRequest(*glue.UpdateTableInput) glue.UpdateTableRequest

	UpdateTriggerRequest(*glue.UpdateTriggerInput) glue.UpdateTriggerRequest

	UpdateUserDefinedFunctionRequest(*glue.UpdateUserDefinedFunctionInput) glue.UpdateUserDefinedFunctionRequest
}

var _ GlueAPI = (*glue.Glue)(nil)
