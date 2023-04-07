//go:build no_runtime_type_checking

// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package mongodbatlas

// Building without runtime type checking enabled, so all the below just return nil

func validateAtlasEncryptionAtRest_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAtlasEncryptionAtRestParameters(scope constructs.Construct, id *string, props *AtlasEncryptionAtRestProps) error {
	return nil
}

