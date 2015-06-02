// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudwatchlogsiface provides an interface for the Amazon CloudWatch Logs.
package cloudwatchlogsiface

import (
	"github.com/awslabs/aws-sdk-go/service/cloudwatchlogs"
)

// CloudWatchLogsAPI is the interface type for cloudwatchlogs.CloudWatchLogs.
type CloudWatchLogsAPI interface {
	CreateLogGroup(*cloudwatchlogs.CreateLogGroupInput) (*cloudwatchlogs.CreateLogGroupOutput, error)

	CreateLogStream(*cloudwatchlogs.CreateLogStreamInput) (*cloudwatchlogs.CreateLogStreamOutput, error)

	DeleteLogGroup(*cloudwatchlogs.DeleteLogGroupInput) (*cloudwatchlogs.DeleteLogGroupOutput, error)

	DeleteLogStream(*cloudwatchlogs.DeleteLogStreamInput) (*cloudwatchlogs.DeleteLogStreamOutput, error)

	DeleteMetricFilter(*cloudwatchlogs.DeleteMetricFilterInput) (*cloudwatchlogs.DeleteMetricFilterOutput, error)

	DeleteRetentionPolicy(*cloudwatchlogs.DeleteRetentionPolicyInput) (*cloudwatchlogs.DeleteRetentionPolicyOutput, error)

	DescribeLogGroups(*cloudwatchlogs.DescribeLogGroupsInput) (*cloudwatchlogs.DescribeLogGroupsOutput, error)

	DescribeLogStreams(*cloudwatchlogs.DescribeLogStreamsInput) (*cloudwatchlogs.DescribeLogStreamsOutput, error)

	DescribeMetricFilters(*cloudwatchlogs.DescribeMetricFiltersInput) (*cloudwatchlogs.DescribeMetricFiltersOutput, error)

	FilterLogEvents(*cloudwatchlogs.FilterLogEventsInput) (*cloudwatchlogs.FilterLogEventsOutput, error)

	GetLogEvents(*cloudwatchlogs.GetLogEventsInput) (*cloudwatchlogs.GetLogEventsOutput, error)

	PutLogEvents(*cloudwatchlogs.PutLogEventsInput) (*cloudwatchlogs.PutLogEventsOutput, error)

	PutMetricFilter(*cloudwatchlogs.PutMetricFilterInput) (*cloudwatchlogs.PutMetricFilterOutput, error)

	PutRetentionPolicy(*cloudwatchlogs.PutRetentionPolicyInput) (*cloudwatchlogs.PutRetentionPolicyOutput, error)

	TestMetricFilter(*cloudwatchlogs.TestMetricFilterInput) (*cloudwatchlogs.TestMetricFilterOutput, error)
}
