// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package machinelearningiface_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/service/machinelearning"
	"github.com/awslabs/aws-sdk-go/service/machinelearning/machinelearningiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*machinelearningiface.MachineLearningAPI)(nil), machinelearning.New(nil))
}
