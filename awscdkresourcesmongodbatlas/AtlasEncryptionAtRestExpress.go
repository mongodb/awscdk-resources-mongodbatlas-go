// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/internal"
)

type AtlasEncryptionAtRestExpress interface {
	constructs.Construct
	AccessList() CfnProjectIpAccessList
	Cluster() CfnCluster
	DatabaseUser() CfnDatabaseUser
	EncryptionAtRest() CfnEncryptionAtRest
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AtlasEncryptionAtRestExpress
type jsiiProxy_AtlasEncryptionAtRestExpress struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AtlasEncryptionAtRestExpress) AccessList() CfnProjectIpAccessList {
	var returns CfnProjectIpAccessList
	_jsii_.Get(
		j,
		"accessList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasEncryptionAtRestExpress) Cluster() CfnCluster {
	var returns CfnCluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasEncryptionAtRestExpress) DatabaseUser() CfnDatabaseUser {
	var returns CfnDatabaseUser
	_jsii_.Get(
		j,
		"databaseUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasEncryptionAtRestExpress) EncryptionAtRest() CfnEncryptionAtRest {
	var returns CfnEncryptionAtRest
	_jsii_.Get(
		j,
		"encryptionAtRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasEncryptionAtRestExpress) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewAtlasEncryptionAtRestExpress(scope constructs.Construct, id *string, props *AtlasEncryptionAtRestExpressProps) AtlasEncryptionAtRestExpress {
	_init_.Initialize()

	if err := validateNewAtlasEncryptionAtRestExpressParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AtlasEncryptionAtRestExpress{}

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRestExpress",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAtlasEncryptionAtRestExpress_Override(a AtlasEncryptionAtRestExpress, scope constructs.Construct, id *string, props *AtlasEncryptionAtRestExpressProps) {
	_init_.Initialize()

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRestExpress",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AtlasEncryptionAtRestExpress_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAtlasEncryptionAtRestExpress_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRestExpress",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AtlasEncryptionAtRestExpress) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

