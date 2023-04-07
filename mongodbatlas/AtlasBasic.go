// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mongodb/awscdk-resources-mongodbatlas-go/mongodbatlas/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mongodb/awscdk-resources-mongodbatlas-go/mongodbatlas/internal"
)

type AtlasBasic interface {
	constructs.Construct
	IpAccessList() CfnProjectIpAccessList
	MCluster() CfnCluster
	MDBUser() CfnDatabaseUser
	MProject() CfnProject
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AtlasBasic
type jsiiProxy_AtlasBasic struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AtlasBasic) IpAccessList() CfnProjectIpAccessList {
	var returns CfnProjectIpAccessList
	_jsii_.Get(
		j,
		"ipAccessList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasBasic) MCluster() CfnCluster {
	var returns CfnCluster
	_jsii_.Get(
		j,
		"mCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasBasic) MDBUser() CfnDatabaseUser {
	var returns CfnDatabaseUser
	_jsii_.Get(
		j,
		"mDBUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasBasic) MProject() CfnProject {
	var returns CfnProject
	_jsii_.Get(
		j,
		"mProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasBasic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of AtlasBasic.
func NewAtlasBasic(scope constructs.Construct, id *string, props *AtlasBasicProps) AtlasBasic {
	_init_.Initialize()

	if err := validateNewAtlasBasicParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AtlasBasic{}

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.AtlasBasic",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of AtlasBasic.
func NewAtlasBasic_Override(a AtlasBasic, scope constructs.Construct, id *string, props *AtlasBasicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.AtlasBasic",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AtlasBasic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAtlasBasic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscdk-resources-mongodbatlas.AtlasBasic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AtlasBasic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

