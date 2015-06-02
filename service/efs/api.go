// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package efs provides a client for Amazon Elastic File System.
package efs

import (
	"sync"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

var oprw sync.Mutex

// CreateFileSystemRequest generates a request for the CreateFileSystem operation.
func (c *EFS) CreateFileSystemRequest(input *CreateFileSystemInput) (req *aws.Request, output *FileSystemDescription) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateFileSystem == nil {
		opCreateFileSystem = &aws.Operation{
			Name:       "CreateFileSystem",
			HTTPMethod: "POST",
			HTTPPath:   "/2015-02-01/file-systems",
		}
	}

	if input == nil {
		input = &CreateFileSystemInput{}
	}

	req = c.newRequest(opCreateFileSystem, input, output)
	output = &FileSystemDescription{}
	req.Data = output
	return
}

// Creates a new, empty file system. The operation requires a creation token
// in the request that Amazon EFS uses to ensure idempotent creation (calling
// the operation with same creation token has no effect). If a file system does
// not currently exist that is owned by the caller's AWS account with the specified
// creation token, this operation does the following:
//
//  Creates a new, empty file system. The file system will have an Amazon EFS
// assigned ID, and an initial lifecycle state "creating".   Returns with the
// description of the created file system.   Otherwise, this operation returns
// a FileSystemAlreadyExists error with the ID of the existing file system.
//
// For basic use cases, you can use a randomly generated UUID for the creation
// token.  The idempotent operation allows you to retry a CreateFileSystem call
// without risk of creating an extra file system. This can happen when an initial
// call fails in a way that leaves it uncertain whether or not a file system
// was actually created. An example might be that a transport level timeout
// occurred or your connection was reset. As long as you use the same creation
// token, if the initial call had succeeded in creating a file system, the client
// can learn of its existence from the FileSystemAlreadyExists error.
//
// The CreateFileSystem call returns while the file system's lifecycle state
// is still "creating". You can check the file system creation status by calling
// the DescribeFileSystems API, which among other things returns the file system
// state.  After the file system is fully created, Amazon EFS sets its lifecycle
// state to "available", at which point you can create one or more mount targets
// for the file system (CreateMountTarget) in your VPC. You mount your Amazon
// EFS file system on an EC2 instances in your VPC via the mount target. For
// more information, see Amazon EFS: How it Works (http://docs.aws.amazon.com/efs/latest/ug/how-it-works.html)
//
//  This operation requires permission for the elasticfilesystem:CreateFileSystem
// action.
func (c *EFS) CreateFileSystem(input *CreateFileSystemInput) (*FileSystemDescription, error) {
	req, out := c.CreateFileSystemRequest(input)
	err := req.Send()
	return out, err
}

var opCreateFileSystem *aws.Operation

// CreateMountTargetRequest generates a request for the CreateMountTarget operation.
func (c *EFS) CreateMountTargetRequest(input *CreateMountTargetInput) (req *aws.Request, output *MountTargetDescription) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateMountTarget == nil {
		opCreateMountTarget = &aws.Operation{
			Name:       "CreateMountTarget",
			HTTPMethod: "POST",
			HTTPPath:   "/2015-02-01/mount-targets",
		}
	}

	if input == nil {
		input = &CreateMountTargetInput{}
	}

	req = c.newRequest(opCreateMountTarget, input, output)
	output = &MountTargetDescription{}
	req.Data = output
	return
}

// Creates a mount target for a file system. You can then mount the file system
// on EC2 instances via the mount target.
//
// You can create one mount target in each Availability Zone in your VPC. All
// EC2 instances in a VPC within a given Availability Zone share a single mount
// target for a given file system. If you have multiple subnets in an Availability
// Zone, you create a mount target in one of the subnets. EC2 instances do not
// need to be in the same subnet as the mount target in order to access their
// file system. For more information, see Amazon EFS: How it Works (http://docs.aws.amazon.com/efs/latest/ug/how-it-works.html).
//
// In the request, you also specify a file system ID for which you are creating
// the mount target and the file system's lifecycle state must be "available"
// (see DescribeFileSystems).
//
//  In the request, you also provide a subnet ID, which serves several purposes:
//
//  It determines the VPC in which Amazon EFS creates the mount target. It
// determines the Availability Zone in which Amazon EFS creates the mount target.
//  It determines the IP address range from which Amazon EFS selects the IP
// address of the mount target if you don't specify an IP address in the request.
//   After creating the mount target, Amazon EFS returns a response that includes,
// a MountTargetId and an IpAddress. You use this IP address when mounting the
// file system in an EC2 instance. You can also use the mount target's DNS name
// when mounting the file system. The EC2 instance on which you mount the file
// system via the mount target can resolve the mount target's DNS name to its
// IP address. For more information, see How it Works: Implementation Overview
// (http://docs.aws.amazon.com/efs/latest/ug/how-it-works.html#how-it-works-implementation)
//
//  Note that you can create mount targets for a file system in only one VPC,
// and there can be only one mount target per Availability Zone. That is, if
// the file system already has one or more mount targets created for it, the
// request to add another mount target must meet the following requirements:
//
//   The subnet specified in the request must belong to the same VPC as the
// subnets of the existing mount targets.
//
//  The subnet specified in the request must not be in the same Availability
// Zone as any of the subnets of the existing mount targets.  If the request
// satisfies the requirements, Amazon EFS does the following:
//
//  Creates a new mount target in the specified subnet.  Also creates a new
// network interface in the subnet as follows:  If the request provides an IpAddress,
// Amazon EFS assigns that IP address to the network interface. Otherwise, Amazon
// EFS assigns a free address in the subnet (in the same way that the Amazon
// EC2 CreateNetworkInterface call does when a request does not specify a primary
// private IP address). If the request provides SecurityGroups, this network
// interface is associated with those security groups. Otherwise, it belongs
// to the default security group for the subnet's VPC. Assigns the description
// "Mount target fsmt-id for file system fs-id" where fsmt-id is the mount target
// ID, and fs-id is the FileSystemId. Sets the requesterManaged property of
// the network interface to "true", and the requesterId value to "EFS".  Each
// Amazon EFS mount target has one corresponding requestor-managed EC2 network
// interface. After the network interface is created, Amazon EFS sets the NetworkInterfaceId
// field in the mount target's description to the network interface ID, and
// the IpAddress field to its address. If network interface creation fails,
// the entire CreateMountTarget operation fails.
//
//   The CreateMountTarget call returns only after creating the network interface,
// but while the mount target state is still "creating". You can check the mount
// target creation status by calling the DescribeFileSystems API, which among
// other things returns the mount target state. We recommend you create a mount
// target in each of the Availability Zones. There are cost considerations for
// using a file system in an Availability Zone through a mount target created
// in another Availability Zone. For more information, go to Amazon EFS (http://aws.amazon.com/efs/)
// product detail page. In addition, by always using a mount target local to
// the instance's Availability Zone, you eliminate a partial failure scenario;
// if the Availablity Zone in which your mount target is created goes down,
// then you won't be able to access your file system through that mount target.
//
// This operation requires permission for the following action on the file
// system:
//
//  elasticfilesystem:CreateMountTarget  This operation also requires permission
// for the following Amazon EC2 actions:
//
//  ec2:DescribeSubnets ec2:DescribeNetworkInterfaces ec2:CreateNetworkInterface
func (c *EFS) CreateMountTarget(input *CreateMountTargetInput) (*MountTargetDescription, error) {
	req, out := c.CreateMountTargetRequest(input)
	err := req.Send()
	return out, err
}

var opCreateMountTarget *aws.Operation

// CreateTagsRequest generates a request for the CreateTags operation.
func (c *EFS) CreateTagsRequest(input *CreateTagsInput) (req *aws.Request, output *CreateTagsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opCreateTags == nil {
		opCreateTags = &aws.Operation{
			Name:       "CreateTags",
			HTTPMethod: "POST",
			HTTPPath:   "/2015-02-01/create-tags/{FileSystemId}",
		}
	}

	if input == nil {
		input = &CreateTagsInput{}
	}

	req = c.newRequest(opCreateTags, input, output)
	output = &CreateTagsOutput{}
	req.Data = output
	return
}

// Creates or overwrites tags associated with a file system. Each tag is a key-value
// pair. If a tag key specified in the request already exists on the file system,
// this operation overwrites its value with the value provided in the request.
// If you add the "Name" tag to your file system, Amazon EFS returns it in the
// response to the DescribeFileSystems API.
//
// This operation requires permission for the elasticfilesystem:CreateTags
// action.
func (c *EFS) CreateTags(input *CreateTagsInput) (*CreateTagsOutput, error) {
	req, out := c.CreateTagsRequest(input)
	err := req.Send()
	return out, err
}

var opCreateTags *aws.Operation

// DeleteFileSystemRequest generates a request for the DeleteFileSystem operation.
func (c *EFS) DeleteFileSystemRequest(input *DeleteFileSystemInput) (req *aws.Request, output *DeleteFileSystemOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteFileSystem == nil {
		opDeleteFileSystem = &aws.Operation{
			Name:       "DeleteFileSystem",
			HTTPMethod: "DELETE",
			HTTPPath:   "/2015-02-01/file-systems/{FileSystemId}",
		}
	}

	if input == nil {
		input = &DeleteFileSystemInput{}
	}

	req = c.newRequest(opDeleteFileSystem, input, output)
	output = &DeleteFileSystemOutput{}
	req.Data = output
	return
}

// Deletes a file system, permanently severing access to its contents. Upon
// return, the file system no longer exists and you will not be able to access
// any contents of the deleted file system.
//
//  You cannot delete a file system that is in use. That is, if the file system
// has any mount targets, you must first delete them. For more information,
// see DescribeMountTargets and DeleteMountTarget.
//
// The DeleteFileSystem call returns while the file system state is still "deleting".
// You can check the file system deletion status by calling the DescribeFileSystems
// API, which returns a list of file systems in your account. If you pass file
// system ID or creation token for the deleted file system, the DescribeFileSystems
// will return a 404 "FileSystemNotFound" error. This operation requires permission
// for the elasticfilesystem:DeleteFileSystem action.
func (c *EFS) DeleteFileSystem(input *DeleteFileSystemInput) (*DeleteFileSystemOutput, error) {
	req, out := c.DeleteFileSystemRequest(input)
	err := req.Send()
	return out, err
}

var opDeleteFileSystem *aws.Operation

// DeleteMountTargetRequest generates a request for the DeleteMountTarget operation.
func (c *EFS) DeleteMountTargetRequest(input *DeleteMountTargetInput) (req *aws.Request, output *DeleteMountTargetOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteMountTarget == nil {
		opDeleteMountTarget = &aws.Operation{
			Name:       "DeleteMountTarget",
			HTTPMethod: "DELETE",
			HTTPPath:   "/2015-02-01/mount-targets/{MountTargetId}",
		}
	}

	if input == nil {
		input = &DeleteMountTargetInput{}
	}

	req = c.newRequest(opDeleteMountTarget, input, output)
	output = &DeleteMountTargetOutput{}
	req.Data = output
	return
}

// Deletes the specified mount target.
//
//  This operation forcibly breaks any mounts of the file system via the mount
// target being deleted, which might disrupt instances or applications using
// those mounts. To avoid applications getting cut off abruptly, you might consider
// unmounting any mounts of the mount target, if feasible. The operation also
// deletes the associated network interface. Uncommitted writes may be lost,
// but breaking a mount target using this operation does not corrupt the file
// system itself. The file system you created remains. You can mount an EC2
// instance in your VPC using another mount target.
//
//  This operation requires permission for the following action on the file
// system:
//
//  elasticfilesystem:DeleteMountTarget  The DeleteMountTarget call returns
// while the mount target state is still "deleting". You can check the mount
// target deletion by calling the DescribeMountTargets API, which returns a
// list of mount target descriptions for the given file system.  The operation
// also requires permission for the following Amazon EC2 action on the mount
// target's network interface:
//
//  ec2:DeleteNetworkInterface
func (c *EFS) DeleteMountTarget(input *DeleteMountTargetInput) (*DeleteMountTargetOutput, error) {
	req, out := c.DeleteMountTargetRequest(input)
	err := req.Send()
	return out, err
}

var opDeleteMountTarget *aws.Operation

// DeleteTagsRequest generates a request for the DeleteTags operation.
func (c *EFS) DeleteTagsRequest(input *DeleteTagsInput) (req *aws.Request, output *DeleteTagsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDeleteTags == nil {
		opDeleteTags = &aws.Operation{
			Name:       "DeleteTags",
			HTTPMethod: "POST",
			HTTPPath:   "/2015-02-01/delete-tags/{FileSystemId}",
		}
	}

	if input == nil {
		input = &DeleteTagsInput{}
	}

	req = c.newRequest(opDeleteTags, input, output)
	output = &DeleteTagsOutput{}
	req.Data = output
	return
}

// Deletes the specified tags from a file system. If the DeleteTags request
// includes a tag key that does not exist, Amazon EFS ignores it; it is not
// an error. For more information about tags and related restrictions, go to
// Tag Restrictions (http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the AWS Billing and Cost Management User Guide.
//
// This operation requires permission for the elasticfilesystem:DeleteTags
// action.
func (c *EFS) DeleteTags(input *DeleteTagsInput) (*DeleteTagsOutput, error) {
	req, out := c.DeleteTagsRequest(input)
	err := req.Send()
	return out, err
}

var opDeleteTags *aws.Operation

// DescribeFileSystemsRequest generates a request for the DescribeFileSystems operation.
func (c *EFS) DescribeFileSystemsRequest(input *DescribeFileSystemsInput) (req *aws.Request, output *DescribeFileSystemsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeFileSystems == nil {
		opDescribeFileSystems = &aws.Operation{
			Name:       "DescribeFileSystems",
			HTTPMethod: "GET",
			HTTPPath:   "/2015-02-01/file-systems",
		}
	}

	if input == nil {
		input = &DescribeFileSystemsInput{}
	}

	req = c.newRequest(opDescribeFileSystems, input, output)
	output = &DescribeFileSystemsOutput{}
	req.Data = output
	return
}

// Returns the description of a specific Amazon EFS file system if either the
// file system CreationToken or the FileSystemId is provided; otherwise, returns
// descriptions of all file systems owned by the caller's AWS account in the
// AWS region of the endpoint that you're calling.
//
//  When retrieving all file system descriptions, you can optionally specify
// the MaxItems parameter to limit the number of descriptions in a response.
// If more file system descriptions remain, Amazon EFS returns a NextMarker,
// an opaque token, in the response. In this case, you should send a subsequent
// request with the Marker request parameter set to the value of NextMarker.
//
//  So to retrieve a list of your file system descriptions, the expected usage
// of this API is an iterative process of first calling DescribeFileSystems
// without the Marker and then continuing to call it with the Marker parameter
// set to the value of the NextMarker from the previous response until the response
// has no NextMarker.
//
//  Note that the implementation may return fewer than MaxItems file system
// descriptions while still including a NextMarker value.
//
//  The order of file systems returned in the response of one DescribeFileSystems
// call, and the order of file systems returned across the responses of a multi-call
// iteration, is unspecified.
//
//  This operation requires permission for the elasticfilesystem:DescribeFileSystems
// action.
func (c *EFS) DescribeFileSystems(input *DescribeFileSystemsInput) (*DescribeFileSystemsOutput, error) {
	req, out := c.DescribeFileSystemsRequest(input)
	err := req.Send()
	return out, err
}

var opDescribeFileSystems *aws.Operation

// DescribeMountTargetSecurityGroupsRequest generates a request for the DescribeMountTargetSecurityGroups operation.
func (c *EFS) DescribeMountTargetSecurityGroupsRequest(input *DescribeMountTargetSecurityGroupsInput) (req *aws.Request, output *DescribeMountTargetSecurityGroupsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeMountTargetSecurityGroups == nil {
		opDescribeMountTargetSecurityGroups = &aws.Operation{
			Name:       "DescribeMountTargetSecurityGroups",
			HTTPMethod: "GET",
			HTTPPath:   "/2015-02-01/mount-targets/{MountTargetId}/security-groups",
		}
	}

	if input == nil {
		input = &DescribeMountTargetSecurityGroupsInput{}
	}

	req = c.newRequest(opDescribeMountTargetSecurityGroups, input, output)
	output = &DescribeMountTargetSecurityGroupsOutput{}
	req.Data = output
	return
}

// Returns the security groups currently in effect for a mount target. This
// operation requires that the network interface of the mount target has been
// created and the life cycle state of the mount target is not "deleted".
//
// This operation requires permissions for the following actions:
//
//   elasticfilesystem:DescribeMountTargetSecurityGroups action on the mount
// target's file system.   ec2:DescribeNetworkInterfaceAttribute action on the
// mount target's network interface.
func (c *EFS) DescribeMountTargetSecurityGroups(input *DescribeMountTargetSecurityGroupsInput) (*DescribeMountTargetSecurityGroupsOutput, error) {
	req, out := c.DescribeMountTargetSecurityGroupsRequest(input)
	err := req.Send()
	return out, err
}

var opDescribeMountTargetSecurityGroups *aws.Operation

// DescribeMountTargetsRequest generates a request for the DescribeMountTargets operation.
func (c *EFS) DescribeMountTargetsRequest(input *DescribeMountTargetsInput) (req *aws.Request, output *DescribeMountTargetsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeMountTargets == nil {
		opDescribeMountTargets = &aws.Operation{
			Name:       "DescribeMountTargets",
			HTTPMethod: "GET",
			HTTPPath:   "/2015-02-01/mount-targets",
		}
	}

	if input == nil {
		input = &DescribeMountTargetsInput{}
	}

	req = c.newRequest(opDescribeMountTargets, input, output)
	output = &DescribeMountTargetsOutput{}
	req.Data = output
	return
}

// Returns the descriptions of the current mount targets for a file system.
// The order of mount targets returned in the response is unspecified.
//
//  This operation requires permission for the elasticfilesystem:DescribeMountTargets
// action on the file system FileSystemId.
func (c *EFS) DescribeMountTargets(input *DescribeMountTargetsInput) (*DescribeMountTargetsOutput, error) {
	req, out := c.DescribeMountTargetsRequest(input)
	err := req.Send()
	return out, err
}

var opDescribeMountTargets *aws.Operation

// DescribeTagsRequest generates a request for the DescribeTags operation.
func (c *EFS) DescribeTagsRequest(input *DescribeTagsInput) (req *aws.Request, output *DescribeTagsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opDescribeTags == nil {
		opDescribeTags = &aws.Operation{
			Name:       "DescribeTags",
			HTTPMethod: "GET",
			HTTPPath:   "/2015-02-01/tags/{FileSystemId}/",
		}
	}

	if input == nil {
		input = &DescribeTagsInput{}
	}

	req = c.newRequest(opDescribeTags, input, output)
	output = &DescribeTagsOutput{}
	req.Data = output
	return
}

// Returns the tags associated with a file system. The order of tags returned
// in the response of one DescribeTags call, and the order of tags returned
// across the responses of a multi-call iteration (when using pagination), is
// unspecified.
//
//  This operation requires permission for the elasticfilesystem:DescribeTags
// action.
func (c *EFS) DescribeTags(input *DescribeTagsInput) (*DescribeTagsOutput, error) {
	req, out := c.DescribeTagsRequest(input)
	err := req.Send()
	return out, err
}

var opDescribeTags *aws.Operation

// ModifyMountTargetSecurityGroupsRequest generates a request for the ModifyMountTargetSecurityGroups operation.
func (c *EFS) ModifyMountTargetSecurityGroupsRequest(input *ModifyMountTargetSecurityGroupsInput) (req *aws.Request, output *ModifyMountTargetSecurityGroupsOutput) {
	oprw.Lock()
	defer oprw.Unlock()

	if opModifyMountTargetSecurityGroups == nil {
		opModifyMountTargetSecurityGroups = &aws.Operation{
			Name:       "ModifyMountTargetSecurityGroups",
			HTTPMethod: "PUT",
			HTTPPath:   "/2015-02-01/mount-targets/{MountTargetId}/security-groups",
		}
	}

	if input == nil {
		input = &ModifyMountTargetSecurityGroupsInput{}
	}

	req = c.newRequest(opModifyMountTargetSecurityGroups, input, output)
	output = &ModifyMountTargetSecurityGroupsOutput{}
	req.Data = output
	return
}

// Modifies the set of security groups in effect for a mount target.
//
// When you create a mount target, Amazon EFS also creates a new network interface
// (see CreateMountTarget). This operation replaces the security groups in effect
// for the network interface associated with a mount target, with the SecurityGroups
// provided in the request. This operation requires that the network interface
// of the mount target has been created and the life cycle state of the mount
// target is not "deleted".
//
// The operation requires permissions for the following actions:
//
//   elasticfilesystem:ModifyMountTargetSecurityGroups action on the mount
// target's file system.   ec2:ModifyNetworkInterfaceAttribute action on the
// mount target's network interface.
func (c *EFS) ModifyMountTargetSecurityGroups(input *ModifyMountTargetSecurityGroupsInput) (*ModifyMountTargetSecurityGroupsOutput, error) {
	req, out := c.ModifyMountTargetSecurityGroupsRequest(input)
	err := req.Send()
	return out, err
}

var opModifyMountTargetSecurityGroups *aws.Operation

type CreateFileSystemInput struct {
	// String of up to 64 ASCII characters. Amazon EFS uses this to ensure idempotent
	// creation.
	CreationToken *string `type:"string" required:"true"`

	metadataCreateFileSystemInput `json:"-" xml:"-"`
}

type metadataCreateFileSystemInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateMountTargetInput struct {
	// The ID of the file system for which to create the mount target.
	FileSystemID *string `locationName:"FileSystemId" type:"string" required:"true"`

	// A valid IPv4 address within the address range of the specified subnet.
	IPAddress *string `locationName:"IpAddress" type:"string"`

	// Up to 5 VPC security group IDs, of the form "sg-xxxxxxxx". These must be
	// for the same VPC as subnet specified.
	SecurityGroups []*string `type:"list"`

	// The ID of the subnet to add the mount target in.
	SubnetID *string `locationName:"SubnetId" type:"string" required:"true"`

	metadataCreateMountTargetInput `json:"-" xml:"-"`
}

type metadataCreateMountTargetInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateTagsInput struct {
	// String. The ID of the file system whose tags you want to modify. This operation
	// modifies only the tags and not the file system.
	FileSystemID *string `location:"uri" locationName:"FileSystemId" type:"string" required:"true"`

	// An array of Tag objects to add. Each Tag object is a key-value pair.
	Tags []*Tag `type:"list" required:"true"`

	metadataCreateTagsInput `json:"-" xml:"-"`
}

type metadataCreateTagsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateTagsOutput struct {
	metadataCreateTagsOutput `json:"-" xml:"-"`
}

type metadataCreateTagsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteFileSystemInput struct {
	// The ID of the file system you want to delete.
	FileSystemID *string `location:"uri" locationName:"FileSystemId" type:"string" required:"true"`

	metadataDeleteFileSystemInput `json:"-" xml:"-"`
}

type metadataDeleteFileSystemInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteFileSystemOutput struct {
	metadataDeleteFileSystemOutput `json:"-" xml:"-"`
}

type metadataDeleteFileSystemOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteMountTargetInput struct {
	// String. The ID of the mount target to delete.
	MountTargetID *string `location:"uri" locationName:"MountTargetId" type:"string" required:"true"`

	metadataDeleteMountTargetInput `json:"-" xml:"-"`
}

type metadataDeleteMountTargetInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteMountTargetOutput struct {
	metadataDeleteMountTargetOutput `json:"-" xml:"-"`
}

type metadataDeleteMountTargetOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteTagsInput struct {
	// String. The ID of the file system whose tags you want to delete.
	FileSystemID *string `location:"uri" locationName:"FileSystemId" type:"string" required:"true"`

	// A list of tag keys to delete.
	TagKeys []*string `type:"list" required:"true"`

	metadataDeleteTagsInput `json:"-" xml:"-"`
}

type metadataDeleteTagsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteTagsOutput struct {
	metadataDeleteTagsOutput `json:"-" xml:"-"`
}

type metadataDeleteTagsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeFileSystemsInput struct {
	// Optional string. Restricts the list to the file system with this creation
	// token (you specify a creation token at the time of creating an Amazon EFS
	// file system).
	CreationToken *string `location:"querystring" locationName:"CreationToken" type:"string"`

	// Optional string. File system ID whose description you want to retrieve.
	FileSystemID *string `location:"querystring" locationName:"FileSystemId" type:"string"`

	// Optional string. Opaque pagination token returned from a previous DescribeFileSystems
	// operation. If present, specifies to continue the list from where the returning
	// call had left off.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// Optional integer. Specifies the maximum number of file systems to return
	// in the response. This parameter value must be greater than 0. The number
	// of items Amazon EFS returns will be the minimum of the MaxItems parameter
	// specified in the request and the service's internal maximum number of items
	// per page.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`

	metadataDescribeFileSystemsInput `json:"-" xml:"-"`
}

type metadataDescribeFileSystemsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeFileSystemsOutput struct {
	// An array of file system descriptions.
	FileSystems []*FileSystemDescription `type:"list"`

	// A string, present if provided by caller in the request.
	Marker *string `type:"string"`

	// A string, present if there are more file systems than returned in the response.
	// You can use the NextMarker in the subsequent request to fetch the descriptions.
	NextMarker *string `type:"string"`

	metadataDescribeFileSystemsOutput `json:"-" xml:"-"`
}

type metadataDescribeFileSystemsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeMountTargetSecurityGroupsInput struct {
	// The ID of the mount target whose security groups you want to retrieve.
	MountTargetID *string `location:"uri" locationName:"MountTargetId" type:"string" required:"true"`

	metadataDescribeMountTargetSecurityGroupsInput `json:"-" xml:"-"`
}

type metadataDescribeMountTargetSecurityGroupsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeMountTargetSecurityGroupsOutput struct {
	// An array of security groups.
	SecurityGroups []*string `type:"list" required:"true"`

	metadataDescribeMountTargetSecurityGroupsOutput `json:"-" xml:"-"`
}

type metadataDescribeMountTargetSecurityGroupsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeMountTargetsInput struct {
	// String. The ID of the file system whose mount targets you want to list.
	FileSystemID *string `location:"querystring" locationName:"FileSystemId" type:"string" required:"true"`

	// Optional. String. Opaque pagination token returned from a previous DescribeMountTargets
	// operation. If present, it specifies to continue the list from where the previous
	// returning call left off.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// Optional. Maximum number of mount targets to return in the response. It must
	// be an integer with a value greater than zero.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`

	metadataDescribeMountTargetsInput `json:"-" xml:"-"`
}

type metadataDescribeMountTargetsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeMountTargetsOutput struct {
	// If the request included the Marker, the response returns that value in this
	// field.
	Marker *string `type:"string"`

	// Returns the file system's mount targets as an array of MountTargetDescription
	// objects.
	MountTargets []*MountTargetDescription `type:"list"`

	// If a value is present, there are more mount targets to return. In a subsequent
	// request, you can provide Marker in your request with this value to retrieve
	// the next set of mount targets.
	NextMarker *string `type:"string"`

	metadataDescribeMountTargetsOutput `json:"-" xml:"-"`
}

type metadataDescribeMountTargetsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeTagsInput struct {
	// The ID of the file system whose tag set you want to retrieve.
	FileSystemID *string `location:"uri" locationName:"FileSystemId" type:"string" required:"true"`

	// Optional. String. Opaque pagination token returned from a previous DescribeTags
	// operation. If present, it specifies to continue the list from where the previous
	// call left off.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// Optional. Maximum number of file system tags to return in the response. It
	// must be an integer with a value greater than zero.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" type:"integer"`

	metadataDescribeTagsInput `json:"-" xml:"-"`
}

type metadataDescribeTagsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeTagsOutput struct {
	// If the request included a Marker, the response returns that value in this
	// field.
	Marker *string `type:"string"`

	// If a value is present, there are more tags to return. In a subsequent request,
	// you can provide the value of NextMarker as the value of the Marker parameter
	// in your next request to retrieve the next set of tags.
	NextMarker *string `type:"string"`

	// Returns tags associated with the file system as an array of Tag objects.
	Tags []*Tag `type:"list" required:"true"`

	metadataDescribeTagsOutput `json:"-" xml:"-"`
}

type metadataDescribeTagsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// This object provides description of a file system.
type FileSystemDescription struct {
	// The time at which the file system was created, in seconds, since 1970-01-01T00:00:00Z.
	CreationTime *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// Opaque string specified in the request.
	CreationToken *string `type:"string" required:"true"`

	// The file system ID assigned by Amazon EFS.
	FileSystemID *string `locationName:"FileSystemId" type:"string" required:"true"`

	// A predefined string value that indicates the lifecycle phase of the file
	// system.
	LifeCycleState *string `type:"string" required:"true"`

	// You can add tags to a file system (see CreateTags) including a "Name" tag.
	// If the file system has a "Name" tag, Amazon EFS returns the value in this
	// field.
	Name *string `type:"string"`

	// The current number of mount targets (see CreateMountTarget) the file system
	// has.
	NumberOfMountTargets *int64 `type:"integer" required:"true"`

	// The AWS account that created the file system. If the file system was created
	// by an IAM user, the parent account to which the user belongs is the owner.
	OwnerID *string `locationName:"OwnerId" type:"string" required:"true"`

	// This object provides the latest known metered size of data stored in the
	// file system, in bytes, in its Value field, and the time at which that size
	// was determined in its Timestamp field. The Timestamp value is the integer
	// number of seconds since 1970-01-01T00:00:00Z. Note that the value does not
	// represent the size of a consistent snapshot of the file system, but it is
	// eventually consistent when there are no writes to the file system. That is,
	// the value will represent actual size only if the file system is not modified
	// for a period longer than a couple of hours. Otherwise, the value is not the
	// exact size the file system was at any instant in time.
	SizeInBytes *FileSystemSize `type:"structure" required:"true"`

	metadataFileSystemDescription `json:"-" xml:"-"`
}

type metadataFileSystemDescription struct {
	SDKShapeTraits bool `type:"structure"`
}

// This object provides the latest known metered size, in bytes, of data stored
// in the file system, in its Value field, and the time at which that size was
// determined in its Timestamp field. Note that the value does not represent
// the size of a consistent snapshot of the file system, but it is eventually
// consistent when there are no writes to the file system. That is, the value
// will represent the actual size only if the file system is not modified for
// a period longer than a couple of hours. Otherwise, the value is not necessarily
// the exact size the file system was at any instant in time.
type FileSystemSize struct {
	// The time at which the size of data, returned in the Value field, was determined.
	// The value is the integer number of seconds since 1970-01-01T00:00:00Z.
	Timestamp *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The latest known metered size, in bytes, of data stored in the file system.
	Value *int64 `type:"long" required:"true"`

	metadataFileSystemSize `json:"-" xml:"-"`
}

type metadataFileSystemSize struct {
	SDKShapeTraits bool `type:"structure"`
}

type ModifyMountTargetSecurityGroupsInput struct {
	// The ID of the mount target whose security groups you want to modify.
	MountTargetID *string `location:"uri" locationName:"MountTargetId" type:"string" required:"true"`

	// An array of up to five VPC security group IDs.
	SecurityGroups []*string `type:"list"`

	metadataModifyMountTargetSecurityGroupsInput `json:"-" xml:"-"`
}

type metadataModifyMountTargetSecurityGroupsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ModifyMountTargetSecurityGroupsOutput struct {
	metadataModifyMountTargetSecurityGroupsOutput `json:"-" xml:"-"`
}

type metadataModifyMountTargetSecurityGroupsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// This object provides description of a mount target.
type MountTargetDescription struct {
	// The ID of the file system for which the mount target is intended.
	FileSystemID *string `locationName:"FileSystemId" type:"string" required:"true"`

	// The address at which the file system may be mounted via the mount target.
	IPAddress *string `locationName:"IpAddress" type:"string"`

	// The lifecycle state the mount target is in.
	LifeCycleState *string `type:"string" required:"true"`

	// The system-assigned mount target ID.
	MountTargetID *string `locationName:"MountTargetId" type:"string" required:"true"`

	// The ID of the network interface that Amazon EFS created when it created the
	// mount target.
	NetworkInterfaceID *string `locationName:"NetworkInterfaceId" type:"string"`

	// The AWS account ID that owns the resource.
	OwnerID *string `locationName:"OwnerId" type:"string"`

	// The ID of the subnet that the mount target is in.
	SubnetID *string `locationName:"SubnetId" type:"string" required:"true"`

	metadataMountTargetDescription `json:"-" xml:"-"`
}

type metadataMountTargetDescription struct {
	SDKShapeTraits bool `type:"structure"`
}

// A tag is a pair of key and value. The allowed characters in keys and values
// are letters, whitespace, and numbers, representable in UTF-8, and the characters
// '+', '-', '=', '.', '_', ':', and '/'.
type Tag struct {
	// Tag key, a string. The key must not start with "aws:".
	Key *string `type:"string" required:"true"`

	// Value of the tag key.
	Value *string `type:"string" required:"true"`

	metadataTag `json:"-" xml:"-"`
}

type metadataTag struct {
	SDKShapeTraits bool `type:"structure"`
}
