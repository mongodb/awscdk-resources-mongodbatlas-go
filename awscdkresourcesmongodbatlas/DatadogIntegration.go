// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mongodb/awscdk-resources-mongodbatlas-go/awscdkresourcesmongodbatlas/internal"
)

type DatadogIntegration interface {
	constructs.Construct
	CfnThirdPartyIntegration() CfnThirdPartyIntegration
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DatadogIntegration
type jsiiProxy_DatadogIntegration struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DatadogIntegration) CfnThirdPartyIntegration() CfnThirdPartyIntegration {
	var returns CfnThirdPartyIntegration
	_jsii_.Get(
		j,
		"cfnThirdPartyIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatadogIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewDatadogIntegration(scope constructs.Construct, id *string, props *DatadogIntegrationProps) DatadogIntegration {
	_init_.Initialize()

	if err := validateNewDatadogIntegrationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatadogIntegration{}

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.DatadogIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDatadogIntegration_Override(d DatadogIntegration, scope constructs.Construct, id *string, props *DatadogIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.DatadogIntegration",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func DatadogIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatadogIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscdk-resources-mongodbatlas.DatadogIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatadogIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

