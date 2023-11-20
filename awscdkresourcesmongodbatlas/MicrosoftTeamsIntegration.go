package awscdkresourcesmongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/v3/internal"
)

type MicrosoftTeamsIntegration interface {
	constructs.Construct
	CfnThirdPartyIntegration() CfnThirdPartyIntegration
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for MicrosoftTeamsIntegration
type jsiiProxy_MicrosoftTeamsIntegration struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_MicrosoftTeamsIntegration) CfnThirdPartyIntegration() CfnThirdPartyIntegration {
	var returns CfnThirdPartyIntegration
	_jsii_.Get(
		j,
		"cfnThirdPartyIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MicrosoftTeamsIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewMicrosoftTeamsIntegration(scope constructs.Construct, id *string, props *MicrosoftTeamsIntegrationProps) MicrosoftTeamsIntegration {
	_init_.Initialize()

	if err := validateNewMicrosoftTeamsIntegrationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MicrosoftTeamsIntegration{}

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.MicrosoftTeamsIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewMicrosoftTeamsIntegration_Override(m MicrosoftTeamsIntegration, scope constructs.Construct, id *string, props *MicrosoftTeamsIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.MicrosoftTeamsIntegration",
		[]interface{}{scope, id, props},
		m,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func MicrosoftTeamsIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMicrosoftTeamsIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscdk-resources-mongodbatlas.MicrosoftTeamsIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MicrosoftTeamsIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

