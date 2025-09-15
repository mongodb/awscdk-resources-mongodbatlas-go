package awscdkresourcesmongodbatlas


// Flex backup configuration.
type BackupSettings struct {
	// Flag that indicates whether backups are performed for this flex cluster.
	//
	// Backup uses flex cluster backups.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

