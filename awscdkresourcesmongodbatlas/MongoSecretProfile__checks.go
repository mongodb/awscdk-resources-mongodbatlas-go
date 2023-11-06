//go:build !no_runtime_type_checking

package awscdkresourcesmongodbatlas

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateMongoSecretProfile_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewMongoSecretProfileParameters(scope constructs.Construct, id *string, profileName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if profileName == nil {
		return fmt.Errorf("parameter profileName is required, but nil was provided")
	}

	return nil
}

