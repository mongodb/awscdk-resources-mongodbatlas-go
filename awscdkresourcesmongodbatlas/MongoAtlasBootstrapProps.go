package awscdkresourcesmongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/v2/jsii"
)

type MongoAtlasBootstrapProps interface {
	// The IAM role name for CloudFormation Extension Execution.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public.html
	//
	// Default: auto generat the name.
	//
	RoleName() *string
	// The secret profile name for MongoDB Atlas.
	// See: https://github.com/mongodb/mongodbatlas-cloudformation-resources/tree/master#2-configure-your-profile
	//
	// Default: generate a dummy secret.
	//
	SecretProfile() *string
	// List of strings representing mongoDB atlas types to activate.
	//
	// the exported AtlasBasicResources can be used for AtlasBasic resources
	// if not provided non resources will be activated.
	//
	// Example:
	//   typesToActivate=["Cluster","Project"] this will actiate MongoDB::Atlas::Project && MongoDB::Atlas::Cluster
	//
	TypesToActivate() *[]*string
}

// The jsii proxy struct for MongoAtlasBootstrapProps
type jsiiProxy_MongoAtlasBootstrapProps struct {
	_ byte // padding
}

func (j *jsiiProxy_MongoAtlasBootstrapProps) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoAtlasBootstrapProps) SecretProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoAtlasBootstrapProps) TypesToActivate() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"typesToActivate",
		&returns,
	)
	return returns
}


func NewMongoAtlasBootstrapProps() MongoAtlasBootstrapProps {
	_init_.Initialize()

	j := jsiiProxy_MongoAtlasBootstrapProps{}

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.MongoAtlasBootstrapProps",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewMongoAtlasBootstrapProps_Override(m MongoAtlasBootstrapProps) {
	_init_.Initialize()

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.MongoAtlasBootstrapProps",
		nil, // no parameters
		m,
	)
}

