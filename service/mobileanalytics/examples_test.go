// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package mobileanalytics_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/mobileanalytics"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleMobileAnalytics_PutEvents() {
	svc := mobileanalytics.New(nil)

	params := &mobileanalytics.PutEventsInput{
		ClientContext: aws.String("String"), // Required
		Events: []*mobileanalytics.Event{ // Required
			&mobileanalytics.Event{ // Required
				EventType: aws.String("String50Chars"),    // Required
				Timestamp: aws.String("ISO8601Timestamp"), // Required
				Attributes: &map[string]*string{
					"Key": aws.String("String0to1000Chars"), // Required
					// More values...
				},
				Metrics: &map[string]*float64{
					"Key": aws.Double(1.0), // Required
					// More values...
				},
				Session: &mobileanalytics.Session{
					Duration:       aws.Long(1),
					ID:             aws.String("String50Chars"),
					StartTimestamp: aws.String("ISO8601Timestamp"),
					StopTimestamp:  aws.String("ISO8601Timestamp"),
				},
				Version: aws.String("String10Chars"),
			},
			// More values...
		},
		ClientContextEncoding: aws.String("String"),
	}
	resp, err := svc.PutEvents(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}
