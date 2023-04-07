// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas


type Event struct {
	Awseventbridge *EventAwseventbridge `field:"optional" json:"awseventbridge" yaml:"awseventbridge"`
	Function *EventFunction `field:"optional" json:"function" yaml:"function"`
}

