// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package opsworksiface provides an interface for the AWS OpsWorks.
package opsworksiface

import (
	"github.com/awslabs/aws-sdk-go/service/opsworks"
)

// OpsWorksAPI is the interface type for opsworks.OpsWorks.
type OpsWorksAPI interface {
	AssignInstance(*opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error)

	AssignVolume(*opsworks.AssignVolumeInput) (*opsworks.AssignVolumeOutput, error)

	AssociateElasticIP(*opsworks.AssociateElasticIPInput) (*opsworks.AssociateElasticIPOutput, error)

	AttachElasticLoadBalancer(*opsworks.AttachElasticLoadBalancerInput) (*opsworks.AttachElasticLoadBalancerOutput, error)

	CloneStack(*opsworks.CloneStackInput) (*opsworks.CloneStackOutput, error)

	CreateApp(*opsworks.CreateAppInput) (*opsworks.CreateAppOutput, error)

	CreateDeployment(*opsworks.CreateDeploymentInput) (*opsworks.CreateDeploymentOutput, error)

	CreateInstance(*opsworks.CreateInstanceInput) (*opsworks.CreateInstanceOutput, error)

	CreateLayer(*opsworks.CreateLayerInput) (*opsworks.CreateLayerOutput, error)

	CreateStack(*opsworks.CreateStackInput) (*opsworks.CreateStackOutput, error)

	CreateUserProfile(*opsworks.CreateUserProfileInput) (*opsworks.CreateUserProfileOutput, error)

	DeleteApp(*opsworks.DeleteAppInput) (*opsworks.DeleteAppOutput, error)

	DeleteInstance(*opsworks.DeleteInstanceInput) (*opsworks.DeleteInstanceOutput, error)

	DeleteLayer(*opsworks.DeleteLayerInput) (*opsworks.DeleteLayerOutput, error)

	DeleteStack(*opsworks.DeleteStackInput) (*opsworks.DeleteStackOutput, error)

	DeleteUserProfile(*opsworks.DeleteUserProfileInput) (*opsworks.DeleteUserProfileOutput, error)

	DeregisterElasticIP(*opsworks.DeregisterElasticIPInput) (*opsworks.DeregisterElasticIPOutput, error)

	DeregisterInstance(*opsworks.DeregisterInstanceInput) (*opsworks.DeregisterInstanceOutput, error)

	DeregisterRDSDBInstance(*opsworks.DeregisterRDSDBInstanceInput) (*opsworks.DeregisterRDSDBInstanceOutput, error)

	DeregisterVolume(*opsworks.DeregisterVolumeInput) (*opsworks.DeregisterVolumeOutput, error)

	DescribeApps(*opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error)

	DescribeCommands(*opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error)

	DescribeDeployments(*opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error)

	DescribeElasticIPs(*opsworks.DescribeElasticIPsInput) (*opsworks.DescribeElasticIPsOutput, error)

	DescribeElasticLoadBalancers(*opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error)

	DescribeInstances(*opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error)

	DescribeLayers(*opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error)

	DescribeLoadBasedAutoScaling(*opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error)

	DescribeMyUserProfile(*opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error)

	DescribePermissions(*opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error)

	DescribeRAIDArrays(*opsworks.DescribeRAIDArraysInput) (*opsworks.DescribeRAIDArraysOutput, error)

	DescribeRDSDBInstances(*opsworks.DescribeRDSDBInstancesInput) (*opsworks.DescribeRDSDBInstancesOutput, error)

	DescribeServiceErrors(*opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error)

	DescribeStackProvisioningParameters(*opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error)

	DescribeStackSummary(*opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error)

	DescribeStacks(*opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error)

	DescribeTimeBasedAutoScaling(*opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error)

	DescribeUserProfiles(*opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error)

	DescribeVolumes(*opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error)

	DetachElasticLoadBalancer(*opsworks.DetachElasticLoadBalancerInput) (*opsworks.DetachElasticLoadBalancerOutput, error)

	DisassociateElasticIP(*opsworks.DisassociateElasticIPInput) (*opsworks.DisassociateElasticIPOutput, error)

	GetHostnameSuggestion(*opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error)

	GrantAccess(*opsworks.GrantAccessInput) (*opsworks.GrantAccessOutput, error)

	RebootInstance(*opsworks.RebootInstanceInput) (*opsworks.RebootInstanceOutput, error)

	RegisterElasticIP(*opsworks.RegisterElasticIPInput) (*opsworks.RegisterElasticIPOutput, error)

	RegisterInstance(*opsworks.RegisterInstanceInput) (*opsworks.RegisterInstanceOutput, error)

	RegisterRDSDBInstance(*opsworks.RegisterRDSDBInstanceInput) (*opsworks.RegisterRDSDBInstanceOutput, error)

	RegisterVolume(*opsworks.RegisterVolumeInput) (*opsworks.RegisterVolumeOutput, error)

	SetLoadBasedAutoScaling(*opsworks.SetLoadBasedAutoScalingInput) (*opsworks.SetLoadBasedAutoScalingOutput, error)

	SetPermission(*opsworks.SetPermissionInput) (*opsworks.SetPermissionOutput, error)

	SetTimeBasedAutoScaling(*opsworks.SetTimeBasedAutoScalingInput) (*opsworks.SetTimeBasedAutoScalingOutput, error)

	StartInstance(*opsworks.StartInstanceInput) (*opsworks.StartInstanceOutput, error)

	StartStack(*opsworks.StartStackInput) (*opsworks.StartStackOutput, error)

	StopInstance(*opsworks.StopInstanceInput) (*opsworks.StopInstanceOutput, error)

	StopStack(*opsworks.StopStackInput) (*opsworks.StopStackOutput, error)

	UnassignInstance(*opsworks.UnassignInstanceInput) (*opsworks.UnassignInstanceOutput, error)

	UnassignVolume(*opsworks.UnassignVolumeInput) (*opsworks.UnassignVolumeOutput, error)

	UpdateApp(*opsworks.UpdateAppInput) (*opsworks.UpdateAppOutput, error)

	UpdateElasticIP(*opsworks.UpdateElasticIPInput) (*opsworks.UpdateElasticIPOutput, error)

	UpdateInstance(*opsworks.UpdateInstanceInput) (*opsworks.UpdateInstanceOutput, error)

	UpdateLayer(*opsworks.UpdateLayerInput) (*opsworks.UpdateLayerOutput, error)

	UpdateMyUserProfile(*opsworks.UpdateMyUserProfileInput) (*opsworks.UpdateMyUserProfileOutput, error)

	UpdateRDSDBInstance(*opsworks.UpdateRDSDBInstanceInput) (*opsworks.UpdateRDSDBInstanceOutput, error)

	UpdateStack(*opsworks.UpdateStackInput) (*opsworks.UpdateStackOutput, error)

	UpdateUserProfile(*opsworks.UpdateUserProfileInput) (*opsworks.UpdateUserProfileOutput, error)

	UpdateVolume(*opsworks.UpdateVolumeInput) (*opsworks.UpdateVolumeOutput, error)
}
