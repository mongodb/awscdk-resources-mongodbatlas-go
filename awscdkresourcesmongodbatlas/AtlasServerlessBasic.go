package awscdkresourcesmongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/internal"
)

type AtlasServerlessBasic interface {
	constructs.Construct
	IpAccessList() CfnProjectIpAccessList
	MDBUser() CfnDatabaseUser
	MProject() CfnProject
	Mserverless() CfnServerlessInstance
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AtlasServerlessBasic
type jsiiProxy_AtlasServerlessBasic struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AtlasServerlessBasic) IpAccessList() CfnProjectIpAccessList {
	var returns CfnProjectIpAccessList
	_jsii_.Get(
		j,
		"ipAccessList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasServerlessBasic) MDBUser() CfnDatabaseUser {
	var returns CfnDatabaseUser
	_jsii_.Get(
		j,
		"mDBUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasServerlessBasic) MProject() CfnProject {
	var returns CfnProject
	_jsii_.Get(
		j,
		"mProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasServerlessBasic) Mserverless() CfnServerlessInstance {
	var returns CfnServerlessInstance
	_jsii_.Get(
		j,
		"mserverless",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasServerlessBasic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of AtlasServerlessBasic.
func NewAtlasServerlessBasic(scope constructs.Construct, id *string, props *AtlasServerlessBasicProps) AtlasServerlessBasic {
	_init_.Initialize()

	if err := validateNewAtlasServerlessBasicParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AtlasServerlessBasic{}

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.AtlasServerlessBasic",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of AtlasServerlessBasic.
func NewAtlasServerlessBasic_Override(a AtlasServerlessBasic, scope constructs.Construct, id *string, props *AtlasServerlessBasicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.AtlasServerlessBasic",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AtlasServerlessBasic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAtlasServerlessBasic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscdk-resources-mongodbatlas.AtlasServerlessBasic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AtlasServerlessBasic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

