// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas


// Comparison operator to apply when checking the current metric value against **matcher[n].value**.
type MatcherOperator string

const (
	// EQUALS.
	MatcherOperator_EQUALS MatcherOperator = "EQUALS"
	// CONTAINS.
	MatcherOperator_CONTAINS MatcherOperator = "CONTAINS"
	// STARTS_WITH.
	MatcherOperator_STARTS_WITH MatcherOperator = "STARTS_WITH"
	// ENDS_WITH.
	MatcherOperator_ENDS_WITH MatcherOperator = "ENDS_WITH"
	// NOT_EQUALS.
	MatcherOperator_NOT_EQUALS MatcherOperator = "NOT_EQUALS"
	// NOT_CONTAINS.
	MatcherOperator_NOT_CONTAINS MatcherOperator = "NOT_CONTAINS"
	// REGEX.
	MatcherOperator_REGEX MatcherOperator = "REGEX"
)

