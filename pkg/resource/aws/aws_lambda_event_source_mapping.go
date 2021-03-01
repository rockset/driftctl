// GENERATED, DO NOT EDIT THIS FILE
package aws

const AwsLambdaEventSourceMappingResourceType = "aws_lambda_event_source_mapping"

type AwsLambdaEventSourceMapping struct {
	BatchSize                      *int    `cty:"batch_size"`
	BisectBatchOnFunctionError     *bool   `cty:"bisect_batch_on_function_error"`
	Enabled                        *bool   `cty:"enabled"`
	EventSourceArn                 *string `cty:"event_source_arn"`
	FunctionArn                    *string `cty:"function_arn" computed:"true"`
	FunctionName                   *string `cty:"function_name"`
	Id                             string  `cty:"id" computed:"true"`
	LastModified                   *string `cty:"last_modified" computed:"true" diff:"-"`
	LastProcessingResult           *string `cty:"last_processing_result" computed:"true" diff:"-"`
	MaximumBatchingWindowInSeconds *int    `cty:"maximum_batching_window_in_seconds"`
	MaximumRecordAgeInSeconds      *int    `cty:"maximum_record_age_in_seconds" computed:"true"`
	MaximumRetryAttempts           *int    `cty:"maximum_retry_attempts" computed:"true"`
	ParallelizationFactor          *int    `cty:"parallelization_factor" computed:"true"`
	StartingPosition               *string `cty:"starting_position" diff:"-"`
	StartingPositionTimestamp      *string `cty:"starting_position_timestamp" diff:"-"`
	State                          *string `cty:"state" computed:"true" diff:"-"`
	StateTransitionReason          *string `cty:"state_transition_reason" computed:"true" diff:"-"`
	Uuid                           *string `cty:"uuid" computed:"true"`
	DestinationConfig              *[]struct {
		OnFailure *[]struct {
			DestinationArn *string `cty:"destination_arn"`
		} `cty:"on_failure"`
	} `cty:"destination_config"`
}

func (r *AwsLambdaEventSourceMapping) TerraformId() string {
	return r.Id
}

func (r *AwsLambdaEventSourceMapping) TerraformType() string {
	return AwsLambdaEventSourceMappingResourceType
}
