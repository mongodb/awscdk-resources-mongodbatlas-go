package awscdkresourcesmongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/internal"
)

type MongoSecretProfile interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for MongoSecretProfile
type jsiiProxy_MongoSecretProfile struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_MongoSecretProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewMongoSecretProfile(scope constructs.Construct, id *string, profileName *string) MongoSecretProfile {
	_init_.Initialize()

	if err := validateNewMongoSecretProfileParameters(scope, id, profileName); err != nil {
		panic(err)
	}
	j := jsiiProxy_MongoSecretProfile{}

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.MongoSecretProfile",
		[]interface{}{scope, id, profileName},
		&j,
	)

	return &j
}

func NewMongoSecretProfile_Override(m MongoSecretProfile, scope constructs.Construct, id *string, profileName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.MongoSecretProfile",
		[]interface{}{scope, id, profileName},
		m,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func MongoSecretProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMongoSecretProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscdk-resources-mongodbatlas.MongoSecretProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MongoSecretProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

