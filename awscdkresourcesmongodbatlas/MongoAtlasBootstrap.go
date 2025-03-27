package awscdkresourcesmongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/internal"
)

// Generate the CFN extension execution role.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry-public.html
//
type MongoAtlasBootstrap interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	Role() awsiam.IRole
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for MongoAtlasBootstrap
type jsiiProxy_MongoAtlasBootstrap struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_MongoAtlasBootstrap) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoAtlasBootstrap) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}


func NewMongoAtlasBootstrap(scope constructs.Construct, id *string, props MongoAtlasBootstrapProps) MongoAtlasBootstrap {
	_init_.Initialize()

	if err := validateNewMongoAtlasBootstrapParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_MongoAtlasBootstrap{}

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.MongoAtlasBootstrap",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewMongoAtlasBootstrap_Override(m MongoAtlasBootstrap, scope constructs.Construct, id *string, props MongoAtlasBootstrapProps) {
	_init_.Initialize()

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.MongoAtlasBootstrap",
		[]interface{}{scope, id, props},
		m,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func MongoAtlasBootstrap_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMongoAtlasBootstrap_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscdk-resources-mongodbatlas.MongoAtlasBootstrap",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoAtlasBootstrap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

