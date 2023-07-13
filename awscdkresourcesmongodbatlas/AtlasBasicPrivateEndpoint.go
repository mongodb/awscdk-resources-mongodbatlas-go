package awscdkresourcesmongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/internal"
)

type AtlasBasicPrivateEndpoint interface {
	constructs.Construct
	AtlasBasic() AtlasBasic
	// The tree node.
	Node() constructs.Node
	PrivateEndpoint() CfnPrivateEndpoint
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AtlasBasicPrivateEndpoint
type jsiiProxy_AtlasBasicPrivateEndpoint struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AtlasBasicPrivateEndpoint) AtlasBasic() AtlasBasic {
	var returns AtlasBasic
	_jsii_.Get(
		j,
		"atlasBasic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasBasicPrivateEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AtlasBasicPrivateEndpoint) PrivateEndpoint() CfnPrivateEndpoint {
	var returns CfnPrivateEndpoint
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}


// Creates an instance of AtlasBasicPrivateEndpoint.
func NewAtlasBasicPrivateEndpoint(scope constructs.Construct, id *string, props *AtlasBasicPrivateEndpointProps) AtlasBasicPrivateEndpoint {
	_init_.Initialize()

	if err := validateNewAtlasBasicPrivateEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AtlasBasicPrivateEndpoint{}

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.AtlasBasicPrivateEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of AtlasBasicPrivateEndpoint.
func NewAtlasBasicPrivateEndpoint_Override(a AtlasBasicPrivateEndpoint, scope constructs.Construct, id *string, props *AtlasBasicPrivateEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.AtlasBasicPrivateEndpoint",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AtlasBasicPrivateEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAtlasBasicPrivateEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscdk-resources-mongodbatlas.AtlasBasicPrivateEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AtlasBasicPrivateEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

