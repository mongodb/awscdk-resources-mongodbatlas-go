package awscdkresourcesmongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/v3/internal"
)

type AtlasEncryptionAtRest interface {
	constructs.Construct
	CfnEncryptionAtRest() CfnEncryptionAtRest
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AtlasEncryptionAtRest
type jsiiProxy_AtlasEncryptionAtRest struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AtlasEncryptionAtRest) CfnEncryptionAtRest() CfnEncryptionAtRest {
	var returns CfnEncryptionAtRest
	_jsii_.Get(
		j,
		"cfnEncryptionAtRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasEncryptionAtRest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewAtlasEncryptionAtRest(scope constructs.Construct, id *string, props *AtlasEncryptionAtRestProps) AtlasEncryptionAtRest {
	_init_.Initialize()

	if err := validateNewAtlasEncryptionAtRestParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AtlasEncryptionAtRest{}

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRest",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAtlasEncryptionAtRest_Override(a AtlasEncryptionAtRest, scope constructs.Construct, id *string, props *AtlasEncryptionAtRestProps) {
	_init_.Initialize()

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRest",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AtlasEncryptionAtRest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAtlasEncryptionAtRest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AtlasEncryptionAtRest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

