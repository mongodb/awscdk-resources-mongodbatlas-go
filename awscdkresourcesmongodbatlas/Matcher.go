// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type Matcher struct {
	// Name of the parameter in the target object that MongoDB Cloud checks.
	//
	// The parameter must match all rules for MongoDB Cloud to check for alert configurations.
	FieldName MatcherFieldName `field:"optional" json:"fieldName" yaml:"fieldName"`
	// Comparison operator to apply when checking the current metric value against **matcher[n].value**.
	Operator MatcherOperator `field:"optional" json:"operator" yaml:"operator"`
	// Value to match or exceed using the specified **matchers.operator**.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

