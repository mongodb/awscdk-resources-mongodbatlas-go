package awscdkresourcesmongodbatlas


type AtlasUser struct {
	// Two alphabet characters that identifies MongoDB Cloud user's geographic location.
	//
	// This parameter uses the ISO 3166-1a2 code format.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Email address that belongs to the MongoDB Cloud user.
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// First or given name that belongs to the MongoDB Cloud user.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Last name, family name, or surname that belongs to the MongoDB Cloud user.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both.
	//
	// RFC 5988 outlines these relationships.
	Links *[]*Link `field:"optional" json:"links" yaml:"links"`
	// Mobile phone number that belongs to the MongoDB Cloud user.
	MobileNumber *string `field:"optional" json:"mobileNumber" yaml:"mobileNumber"`
	// Password applied with the username to log in to MongoDB Cloud.
	//
	// MongoDB Cloud does not return this parameter except in response to creating a new MongoDB Cloud user. Only the MongoDB Cloud user can update their password after it has been set from the MongoDB Cloud console.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// List of objects that display the MongoDB Cloud user's roles and the corresponding organization or project to which that role applies.
	//
	// A role can apply to one organization or one project but not both.
	Roles *[]*AtlasRole `field:"optional" json:"roles" yaml:"roles"`
	// List of unique 24-hexadecimal digit strings that identifies the teams to which this MongoDB Cloud user belongs.
	TeamIds *[]*string `field:"optional" json:"teamIds" yaml:"teamIds"`
	// Email address that represents the username of the MongoDB Cloud user.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

