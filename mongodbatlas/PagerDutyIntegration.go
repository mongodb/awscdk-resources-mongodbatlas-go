// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/mongodb/awscdk-resources-mongodbatlas-go/mongodbatlas/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/mongodb/awscdk-resources-mongodbatlas-go/mongodbatlas/internal"
)

type PagerDutyIntegration interface {
	constructs.Construct
	CfnThirdPartyIntegration() CfnThirdPartyIntegration
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for PagerDutyIntegration
type jsiiProxy_PagerDutyIntegration struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PagerDutyIntegration) CfnThirdPartyIntegration() CfnThirdPartyIntegration {
	var returns CfnThirdPartyIntegration
	_jsii_.Get(
		j,
		"cfnThirdPartyIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PagerDutyIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewPagerDutyIntegration(scope constructs.Construct, id *string, props *PagerDutyIntegrationProps) PagerDutyIntegration {
	_init_.Initialize()

	if err := validateNewPagerDutyIntegrationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PagerDutyIntegration{}

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.PagerDutyIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPagerDutyIntegration_Override(p PagerDutyIntegration, scope constructs.Construct, id *string, props *PagerDutyIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"awscdk-resources-mongodbatlas.PagerDutyIntegration",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func PagerDutyIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePagerDutyIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscdk-resources-mongodbatlas.PagerDutyIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PagerDutyIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

