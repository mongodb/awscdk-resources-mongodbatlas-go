// MongoDB Atlas CDK Construct Library for AWS CloudFormation Resources
package awscdkresourcesmongodbatlas

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AccessListDefinition",
		reflect.TypeOf((*AccessListDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Action",
		reflect.TypeOf((*Action)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AdvancedAutoScaling",
		reflect.TypeOf((*AdvancedAutoScaling)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AdvancedRegionConfig",
		reflect.TypeOf((*AdvancedRegionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.AdvancedRegionConfigProviderName",
		reflect.TypeOf((*AdvancedRegionConfigProviderName)(nil)).Elem(),
		map[string]interface{}{
			"AWS": AdvancedRegionConfigProviderName_AWS,
			"GCP": AdvancedRegionConfigProviderName_GCP,
			"AZURE": AdvancedRegionConfigProviderName_AZURE,
			"TENANT": AdvancedRegionConfigProviderName_TENANT,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AdvancedReplicationSpec",
		reflect.TypeOf((*AdvancedReplicationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupCopySettingView",
		reflect.TypeOf((*ApiAtlasDiskBackupCopySettingView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotMemberView",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotMemberView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider)(nil)).Elem(),
		map[string]interface{}{
			"AWS": ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider_AWS,
			"AZURE": ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider_AZURE,
			"GCP": ApiAtlasDiskBackupShardedClusterSnapshotMemberViewCloudProvider_GCP,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotView",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType)(nil)).Elem(),
		map[string]interface{}{
			"HOURLY": ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_HOURLY,
			"DAILY": ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_DAILY,
			"WEEKLY": ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_WEEKLY,
			"MONTHLY": ApiAtlasDiskBackupShardedClusterSnapshotViewFrequencyType_MONTHLY,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType)(nil)).Elem(),
		map[string]interface{}{
			"ON_DEMAND": ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType_ON_DEMAND,
			"SCHEDULED": ApiAtlasDiskBackupShardedClusterSnapshotViewSnapshotType_SCHEDULED,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotViewStatus",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotViewStatus)(nil)).Elem(),
		map[string]interface{}{
			"QUEUED": ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_QUEUED,
			"IN_PROGRESS": ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_IN_PROGRESS,
			"COMPLETED": ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_COMPLETED,
			"FAILED": ApiAtlasDiskBackupShardedClusterSnapshotViewStatus_FAILED,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ApiAtlasDiskBackupShardedClusterSnapshotViewType",
		reflect.TypeOf((*ApiAtlasDiskBackupShardedClusterSnapshotViewType)(nil)).Elem(),
		map[string]interface{}{
			"REPLICA_SET": ApiAtlasDiskBackupShardedClusterSnapshotViewType_REPLICA_SET,
			"SHARDED_CLUSTER": ApiAtlasDiskBackupShardedClusterSnapshotViewType_SHARDED_CLUSTER,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasFtsAnalyzersViewManual",
		reflect.TypeOf((*ApiAtlasFtsAnalyzersViewManual)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasFtsMappingsViewManual",
		reflect.TypeOf((*ApiAtlasFtsMappingsViewManual)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasFtsSynonymMappingDefinitionView",
		reflect.TypeOf((*ApiAtlasFtsSynonymMappingDefinitionView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiAtlasNdsUserToDnMappingView",
		reflect.TypeOf((*ApiAtlasNdsUserToDnMappingView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiDeleteCopiedBackupsView",
		reflect.TypeOf((*ApiDeleteCopiedBackupsView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiPolicyItemView",
		reflect.TypeOf((*ApiPolicyItemView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ApiPolicyView",
		reflect.TypeOf((*ApiPolicyView)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.AtlasBasic",
		reflect.TypeOf((*AtlasBasic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "ipAccessList", GoGetter: "IpAccessList"},
			_jsii_.MemberProperty{JsiiProperty: "mCluster", GoGetter: "MCluster"},
			_jsii_.MemberProperty{JsiiProperty: "mDBUser", GoGetter: "MDBUser"},
			_jsii_.MemberProperty{JsiiProperty: "mProject", GoGetter: "MProject"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AtlasBasic{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.AtlasBasicPrivateEndpoint",
		reflect.TypeOf((*AtlasBasicPrivateEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "atlasBasic", GoGetter: "AtlasBasic"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "privateEndpoint", GoGetter: "PrivateEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AtlasBasicPrivateEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AtlasBasicPrivateEndpointProps",
		reflect.TypeOf((*AtlasBasicPrivateEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AtlasBasicProps",
		reflect.TypeOf((*AtlasBasicProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRest",
		reflect.TypeOf((*AtlasEncryptionAtRest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cfnEncryptionAtRest", GoGetter: "CfnEncryptionAtRest"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AtlasEncryptionAtRest{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRestExpress",
		reflect.TypeOf((*AtlasEncryptionAtRestExpress)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessList", GoGetter: "AccessList"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "databaseUser", GoGetter: "DatabaseUser"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionAtRest", GoGetter: "EncryptionAtRest"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AtlasEncryptionAtRestExpress{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRestExpressProps",
		reflect.TypeOf((*AtlasEncryptionAtRestExpressProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AtlasEncryptionAtRestProps",
		reflect.TypeOf((*AtlasEncryptionAtRestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AtlasRole",
		reflect.TypeOf((*AtlasRole)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.AtlasRoleRoleName",
		reflect.TypeOf((*AtlasRoleRoleName)(nil)).Elem(),
		map[string]interface{}{
			"ORG_OWNER": AtlasRoleRoleName_ORG_OWNER,
			"ORG_MEMBER": AtlasRoleRoleName_ORG_MEMBER,
			"ORG_GROUP_CREATOR": AtlasRoleRoleName_ORG_GROUP_CREATOR,
			"ORG_BILLING_ADMIN": AtlasRoleRoleName_ORG_BILLING_ADMIN,
			"ORG_READ_ONLY": AtlasRoleRoleName_ORG_READ_ONLY,
			"GROUP_CLUSTER_MANAGER": AtlasRoleRoleName_GROUP_CLUSTER_MANAGER,
			"GROUP_DATA_ACCESS_ADMIN": AtlasRoleRoleName_GROUP_DATA_ACCESS_ADMIN,
			"GROUP_DATA_ACCESS_READ_ONLY": AtlasRoleRoleName_GROUP_DATA_ACCESS_READ_ONLY,
			"GROUP_DATA_ACCESS_READ_WRITE": AtlasRoleRoleName_GROUP_DATA_ACCESS_READ_WRITE,
			"GROUP_OWNER": AtlasRoleRoleName_GROUP_OWNER,
			"GROUP_READ_ONLY": AtlasRoleRoleName_GROUP_READ_ONLY,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AtlasUser",
		reflect.TypeOf((*AtlasUser)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AuthConfig",
		reflect.TypeOf((*AuthConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.AuthConfigOperationType",
		reflect.TypeOf((*AuthConfigOperationType)(nil)).Elem(),
		map[string]interface{}{
			"LOGIN": AuthConfigOperationType_LOGIN,
			"CREATE": AuthConfigOperationType_CREATE,
			"DELETE": AuthConfigOperationType_DELETE,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.AuthConfigProviders",
		reflect.TypeOf((*AuthConfigProviders)(nil)).Elem(),
		map[string]interface{}{
			"ANON_USER": AuthConfigProviders_ANON_USER,
			"API_KEY": AuthConfigProviders_API_KEY,
			"CUSTOM_TOKEN": AuthConfigProviders_CUSTOM_TOKEN,
			"CUSTOM_FUNCTION": AuthConfigProviders_CUSTOM_FUNCTION,
			"LOCAL_USERPASS": AuthConfigProviders_LOCAL_USERPASS,
			"OAUTH2_APPLE": AuthConfigProviders_OAUTH2_APPLE,
			"OAUTH2_FACEBOOK": AuthConfigProviders_OAUTH2_FACEBOOK,
			"OAUTH2_GOOGLE": AuthConfigProviders_OAUTH2_GOOGLE,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.AwsKmsConfiguration",
		reflect.TypeOf((*AwsKmsConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnAlertConfiguration",
		reflect.TypeOf((*CfnAlertConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrEnabled", GoGetter: "AttrEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrUpdated", GoGetter: "AttrUpdated"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlertConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnAlertConfigurationProps",
		reflect.TypeOf((*CfnAlertConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnAuditing",
		reflect.TypeOf((*CfnAuditing)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrAuditAuthorizationSuccess", GoGetter: "AttrAuditAuthorizationSuccess"},
			_jsii_.MemberProperty{JsiiProperty: "attrAuditFilter", GoGetter: "AttrAuditFilter"},
			_jsii_.MemberProperty{JsiiProperty: "attrConfigurationType", GoGetter: "AttrConfigurationType"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAuditing{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnAuditingProps",
		reflect.TypeOf((*CfnAuditingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCloudBackUpRestoreJobs",
		reflect.TypeOf((*CfnCloudBackUpRestoreJobs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrDeliveryUrl", GoGetter: "AttrDeliveryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "attrExpiresAt", GoGetter: "AttrExpiresAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrFinishedAt", GoGetter: "AttrFinishedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrTimestamp", GoGetter: "AttrTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudBackUpRestoreJobs{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCloudBackUpRestoreJobsProps",
		reflect.TypeOf((*CfnCloudBackUpRestoreJobsProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnCloudBackUpRestoreJobsPropsDeliveryType",
		reflect.TypeOf((*CfnCloudBackUpRestoreJobsPropsDeliveryType)(nil)).Elem(),
		map[string]interface{}{
			"DOWNLOAD": CfnCloudBackUpRestoreJobsPropsDeliveryType_DOWNLOAD,
			"AUTOMATED": CfnCloudBackUpRestoreJobsPropsDeliveryType_AUTOMATED,
			"POINT_IN_TIME": CfnCloudBackUpRestoreJobsPropsDeliveryType_POINT_IN_TIME,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSchedule",
		reflect.TypeOf((*CfnCloudBackupSchedule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrClusterId", GoGetter: "AttrClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "attrNextSnapshot", GoGetter: "AttrNextSnapshot"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudBackupSchedule{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCloudBackupScheduleProps",
		reflect.TypeOf((*CfnCloudBackupScheduleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshot",
		reflect.TypeOf((*CfnCloudBackupSnapshot)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCloudProvider", GoGetter: "AttrCloudProvider"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrExpiresAt", GoGetter: "AttrExpiresAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrMasterKeyUUID", GoGetter: "AttrMasterKeyUUID"},
			_jsii_.MemberProperty{JsiiProperty: "attrMongodVersion", GoGetter: "AttrMongodVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrReplicaSetName", GoGetter: "AttrReplicaSetName"},
			_jsii_.MemberProperty{JsiiProperty: "attrSnapshotId", GoGetter: "AttrSnapshotId"},
			_jsii_.MemberProperty{JsiiProperty: "attrSnapshotIds", GoGetter: "AttrSnapshotIds"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "attrStorageSizeBytes", GoGetter: "AttrStorageSizeBytes"},
			_jsii_.MemberProperty{JsiiProperty: "attrType", GoGetter: "AttrType"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudBackupSnapshot{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshotExportBucket",
		reflect.TypeOf((*CfnCloudBackupSnapshotExportBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudBackupSnapshotExportBucket{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshotExportBucketProps",
		reflect.TypeOf((*CfnCloudBackupSnapshotExportBucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshotProps",
		reflect.TypeOf((*CfnCloudBackupSnapshotProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshotPropsFrequencyType",
		reflect.TypeOf((*CfnCloudBackupSnapshotPropsFrequencyType)(nil)).Elem(),
		map[string]interface{}{
			"HOURLY": CfnCloudBackupSnapshotPropsFrequencyType_HOURLY,
			"DAILY": CfnCloudBackupSnapshotPropsFrequencyType_DAILY,
			"WEEKLY": CfnCloudBackupSnapshotPropsFrequencyType_WEEKLY,
			"MONTHLY": CfnCloudBackupSnapshotPropsFrequencyType_MONTHLY,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnCloudBackupSnapshotPropsSnapshotType",
		reflect.TypeOf((*CfnCloudBackupSnapshotPropsSnapshotType)(nil)).Elem(),
		map[string]interface{}{
			"ON_DEMAND": CfnCloudBackupSnapshotPropsSnapshotType_ON_DEMAND,
			"SCHEDULED": CfnCloudBackupSnapshotPropsSnapshotType_SCHEDULED,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCluster",
		reflect.TypeOf((*CfnCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedDate", GoGetter: "AttrCreatedDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrMongoDBVersion", GoGetter: "AttrMongoDBVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrStateName", GoGetter: "AttrStateName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "connectionStrings", GoGetter: "ConnectionStrings"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnClusterProps",
		reflect.TypeOf((*CfnClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnClusterPropsBiConnector",
		reflect.TypeOf((*CfnClusterPropsBiConnector)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnClusterPropsEncryptionAtRestProvider",
		reflect.TypeOf((*CfnClusterPropsEncryptionAtRestProvider)(nil)).Elem(),
		map[string]interface{}{
			"AWS": CfnClusterPropsEncryptionAtRestProvider_AWS,
			"GCP": CfnClusterPropsEncryptionAtRestProvider_GCP,
			"AZURE": CfnClusterPropsEncryptionAtRestProvider_AZURE,
			"NONE": CfnClusterPropsEncryptionAtRestProvider_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnClusterPropsLabels",
		reflect.TypeOf((*CfnClusterPropsLabels)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCustomDbRole",
		reflect.TypeOf((*CfnCustomDbRole)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomDbRole{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCustomDbRoleProps",
		reflect.TypeOf((*CfnCustomDbRoleProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnCustomDnsConfigurationClusterAws",
		reflect.TypeOf((*CfnCustomDnsConfigurationClusterAws)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomDnsConfigurationClusterAws{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnCustomDnsConfigurationClusterAwsProps",
		reflect.TypeOf((*CfnCustomDnsConfigurationClusterAwsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnDataLakes",
		reflect.TypeOf((*CfnDataLakes)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrHostnames", GoGetter: "AttrHostnames"},
			_jsii_.MemberProperty{JsiiProperty: "attrStartDate", GoGetter: "AttrStartDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataLakes{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnDataLakesProps",
		reflect.TypeOf((*CfnDataLakesProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnDatabaseUser",
		reflect.TypeOf((*CfnDatabaseUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrUserCFNIdentifier", GoGetter: "AttrUserCFNIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatabaseUser{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnDatabaseUserProps",
		reflect.TypeOf((*CfnDatabaseUserProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnDatabaseUserPropsAwsiamType",
		reflect.TypeOf((*CfnDatabaseUserPropsAwsiamType)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CfnDatabaseUserPropsAwsiamType_NONE,
			"USER": CfnDatabaseUserPropsAwsiamType_USER,
			"ROLE": CfnDatabaseUserPropsAwsiamType_ROLE,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnDatabaseUserPropsLdapAuthType",
		reflect.TypeOf((*CfnDatabaseUserPropsLdapAuthType)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CfnDatabaseUserPropsLdapAuthType_NONE,
			"USER": CfnDatabaseUserPropsLdapAuthType_USER,
			"GROUP": CfnDatabaseUserPropsLdapAuthType_GROUP,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnDatabaseUserPropsX509Type",
		reflect.TypeOf((*CfnDatabaseUserPropsX509Type)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CfnDatabaseUserPropsX509Type_NONE,
			"MANAGED": CfnDatabaseUserPropsX509Type_MANAGED,
			"CUSTOMER": CfnDatabaseUserPropsX509Type_CUSTOMER,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnEncryptionAtRest",
		reflect.TypeOf((*CfnEncryptionAtRest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEncryptionAtRest{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnEncryptionAtRestProps",
		reflect.TypeOf((*CfnEncryptionAtRestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnFederatedSettingsOrgRoleMapping",
		reflect.TypeOf((*CfnFederatedSettingsOrgRoleMapping)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFederatedSettingsOrgRoleMapping{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnFederatedSettingsOrgRoleMappingProps",
		reflect.TypeOf((*CfnFederatedSettingsOrgRoleMappingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnGlobalClusterConfig",
		reflect.TypeOf((*CfnGlobalClusterConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGlobalClusterConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnGlobalClusterConfigProps",
		reflect.TypeOf((*CfnGlobalClusterConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnLdapConfiguration",
		reflect.TypeOf((*CfnLdapConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLdapConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnLdapConfigurationProps",
		reflect.TypeOf((*CfnLdapConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnLdapVerify",
		reflect.TypeOf((*CfnLdapVerify)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrRequestId", GoGetter: "AttrRequestId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLdapVerify{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnLdapVerifyProps",
		reflect.TypeOf((*CfnLdapVerifyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnMaintenanceWindow",
		reflect.TypeOf((*CfnMaintenanceWindow)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMaintenanceWindow{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnMaintenanceWindowProps",
		reflect.TypeOf((*CfnMaintenanceWindowProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnNetworkContainer",
		reflect.TypeOf((*CfnNetworkContainer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkContainer{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnNetworkContainerProps",
		reflect.TypeOf((*CfnNetworkContainerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnNetworkPeering",
		reflect.TypeOf((*CfnNetworkPeering)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrConnectionId", GoGetter: "AttrConnectionId"},
			_jsii_.MemberProperty{JsiiProperty: "attrErrorStateName", GoGetter: "AttrErrorStateName"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatusName", GoGetter: "AttrStatusName"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetworkPeering{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnNetworkPeeringProps",
		reflect.TypeOf((*CfnNetworkPeeringProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnOnlineArchive",
		reflect.TypeOf((*CfnOnlineArchive)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArchiveId", GoGetter: "AttrArchiveId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "attrTotalCount", GoGetter: "AttrTotalCount"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOnlineArchive{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnOnlineArchiveProps",
		reflect.TypeOf((*CfnOnlineArchiveProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnOnlineArchivePropsCollectionType",
		reflect.TypeOf((*CfnOnlineArchivePropsCollectionType)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": CfnOnlineArchivePropsCollectionType_STANDARD,
			"TIMESERIES": CfnOnlineArchivePropsCollectionType_TIMESERIES,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnOrgInvitation",
		reflect.TypeOf((*CfnOrgInvitation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrExpiresAt", GoGetter: "AttrExpiresAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrInviterUsername", GoGetter: "AttrInviterUsername"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOrgInvitation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnOrgInvitationProps",
		reflect.TypeOf((*CfnOrgInvitationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnOrgInvitationPropsRoles",
		reflect.TypeOf((*CfnOrgInvitationPropsRoles)(nil)).Elem(),
		map[string]interface{}{
			"ORG_OWNER": CfnOrgInvitationPropsRoles_ORG_OWNER,
			"ORG_MEMBER": CfnOrgInvitationPropsRoles_ORG_MEMBER,
			"ORG_GROUP_CREATOR": CfnOrgInvitationPropsRoles_ORG_GROUP_CREATOR,
			"ORG_BILLING_ADMIN": CfnOrgInvitationPropsRoles_ORG_BILLING_ADMIN,
			"ORG_READ_ONLY": CfnOrgInvitationPropsRoles_ORG_READ_ONLY,
			"GROUP_CLUSTER_MANAGER": CfnOrgInvitationPropsRoles_GROUP_CLUSTER_MANAGER,
			"GROUP_DATA_ACCESS_ADMIN": CfnOrgInvitationPropsRoles_GROUP_DATA_ACCESS_ADMIN,
			"GROUP_DATA_ACCESS_READ_ONLY": CfnOrgInvitationPropsRoles_GROUP_DATA_ACCESS_READ_ONLY,
			"GROUP_DATA_ACCESS_READ_WRITE": CfnOrgInvitationPropsRoles_GROUP_DATA_ACCESS_READ_WRITE,
			"GROUP_OWNER": CfnOrgInvitationPropsRoles_GROUP_OWNER,
			"GROUP_READ_ONLY": CfnOrgInvitationPropsRoles_GROUP_READ_ONLY,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnPrivateEndPointRegionalMode",
		reflect.TypeOf((*CfnPrivateEndPointRegionalMode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrivateEndPointRegionalMode{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnPrivateEndPointRegionalModeProps",
		reflect.TypeOf((*CfnPrivateEndPointRegionalModeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnPrivateEndpoint",
		reflect.TypeOf((*CfnPrivateEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrInterfaceEndpoints", GoGetter: "AttrInterfaceEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrivateEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnPrivateEndpointAdl",
		reflect.TypeOf((*CfnPrivateEndpointAdl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPrivateEndpointAdl{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnPrivateEndpointAdlProps",
		reflect.TypeOf((*CfnPrivateEndpointAdlProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnPrivateEndpointProps",
		reflect.TypeOf((*CfnPrivateEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnProject",
		reflect.TypeOf((*CfnProject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrClusterCount", GoGetter: "AttrClusterCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreated", GoGetter: "AttrCreated"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrProjectOwnerId", GoGetter: "AttrProjectOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProject{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnProjectInvitation",
		reflect.TypeOf((*CfnProjectInvitation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrExpiresAt", GoGetter: "AttrExpiresAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrInviterUsername", GoGetter: "AttrInviterUsername"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectInvitation{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnProjectInvitationProps",
		reflect.TypeOf((*CfnProjectInvitationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnProjectInvitationPropsRoles",
		reflect.TypeOf((*CfnProjectInvitationPropsRoles)(nil)).Elem(),
		map[string]interface{}{
			"GROUP_CLUSTER_MANAGER": CfnProjectInvitationPropsRoles_GROUP_CLUSTER_MANAGER,
			"GROUP_DATA_ACCESS_ADMIN": CfnProjectInvitationPropsRoles_GROUP_DATA_ACCESS_ADMIN,
			"GROUP_DATA_ACCESS_READ_ONLY": CfnProjectInvitationPropsRoles_GROUP_DATA_ACCESS_READ_ONLY,
			"GROUP_DATA_ACCESS_READ_WRITE": CfnProjectInvitationPropsRoles_GROUP_DATA_ACCESS_READ_WRITE,
			"GROUP_OWNER": CfnProjectInvitationPropsRoles_GROUP_OWNER,
			"GROUP_READ_ONLY": CfnProjectInvitationPropsRoles_GROUP_READ_ONLY,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnProjectIpAccessList",
		reflect.TypeOf((*CfnProjectIpAccessList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectIpAccessList{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnProjectIpAccessListProps",
		reflect.TypeOf((*CfnProjectIpAccessListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnProjectProps",
		reflect.TypeOf((*CfnProjectProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnSearchIndex",
		reflect.TypeOf((*CfnSearchIndex)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrIndexId", GoGetter: "AttrIndexId"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSearchIndex{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnSearchIndexProps",
		reflect.TypeOf((*CfnSearchIndexProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnServerlessInstance",
		reflect.TypeOf((*CfnServerlessInstance)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreateDate", GoGetter: "AttrCreateDate"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrMongoDBVersion", GoGetter: "AttrMongoDBVersion"},
			_jsii_.MemberProperty{JsiiProperty: "attrStateName", GoGetter: "AttrStateName"},
			_jsii_.MemberProperty{JsiiProperty: "attrTotalCount", GoGetter: "AttrTotalCount"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServerlessInstance{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnServerlessInstanceProps",
		reflect.TypeOf((*CfnServerlessInstanceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnTeams",
		reflect.TypeOf((*CfnTeams)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrTeamId", GoGetter: "AttrTeamId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTeams{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnTeamsProps",
		reflect.TypeOf((*CfnTeamsProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnTeamsPropsRoleNames",
		reflect.TypeOf((*CfnTeamsPropsRoleNames)(nil)).Elem(),
		map[string]interface{}{
			"GROUP_CLUSTER_MANAGER": CfnTeamsPropsRoleNames_GROUP_CLUSTER_MANAGER,
			"GROUP_DATA_ACCESS_ADMIN": CfnTeamsPropsRoleNames_GROUP_DATA_ACCESS_ADMIN,
			"GROUP_DATA_ACCESS_READ_ONLY": CfnTeamsPropsRoleNames_GROUP_DATA_ACCESS_READ_ONLY,
			"GROUP_DATA_ACCESS_READ_WRITE": CfnTeamsPropsRoleNames_GROUP_DATA_ACCESS_READ_WRITE,
			"GROUP_OWNER": CfnTeamsPropsRoleNames_GROUP_OWNER,
			"GROUP_READ_ONLY": CfnTeamsPropsRoleNames_GROUP_READ_ONLY,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnThirdPartyIntegration",
		reflect.TypeOf((*CfnThirdPartyIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnThirdPartyIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnThirdPartyIntegrationProps",
		reflect.TypeOf((*CfnThirdPartyIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnThirdPartyIntegrationPropsScheme",
		reflect.TypeOf((*CfnThirdPartyIntegrationPropsScheme)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": CfnThirdPartyIntegrationPropsScheme_HTTP,
			"HTTPS": CfnThirdPartyIntegrationPropsScheme_HTTPS,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnThirdPartyIntegrationPropsServiceDiscovery",
		reflect.TypeOf((*CfnThirdPartyIntegrationPropsServiceDiscovery)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": CfnThirdPartyIntegrationPropsServiceDiscovery_HTTP,
			"FILE": CfnThirdPartyIntegrationPropsServiceDiscovery_FILE,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CfnThirdPartyIntegrationPropsType",
		reflect.TypeOf((*CfnThirdPartyIntegrationPropsType)(nil)).Elem(),
		map[string]interface{}{
			"PAGER_DUTY": CfnThirdPartyIntegrationPropsType_PAGER_DUTY,
			"MICROSOFT_TEAMS": CfnThirdPartyIntegrationPropsType_MICROSOFT_TEAMS,
			"SLACK": CfnThirdPartyIntegrationPropsType_SLACK,
			"DATADOG": CfnThirdPartyIntegrationPropsType_DATADOG,
			"OPS_GENIE": CfnThirdPartyIntegrationPropsType_OPS_GENIE,
			"VICTOR_OPS": CfnThirdPartyIntegrationPropsType_VICTOR_OPS,
			"WEBHOOK": CfnThirdPartyIntegrationPropsType_WEBHOOK,
			"PROMETHEUS": CfnThirdPartyIntegrationPropsType_PROMETHEUS,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnTrigger",
		reflect.TypeOf((*CfnTrigger)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrigger{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnTriggerProps",
		reflect.TypeOf((*CfnTriggerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.CfnX509AuthenticationDatabaseUser",
		reflect.TypeOf((*CfnX509AuthenticationDatabaseUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnX509AuthenticationDatabaseUser{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CfnX509AuthenticationDatabaseUserProps",
		reflect.TypeOf((*CfnX509AuthenticationDatabaseUserProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ClusterProps",
		reflect.TypeOf((*ClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Compute",
		reflect.TypeOf((*Compute)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ConnectionStrings",
		reflect.TypeOf((*ConnectionStrings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CriteriaView",
		reflect.TypeOf((*CriteriaView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CriteriaViewDateFormat",
		reflect.TypeOf((*CriteriaViewDateFormat)(nil)).Elem(),
		map[string]interface{}{
			"ISODATE": CriteriaViewDateFormat_ISODATE,
			"EPOCH_SECONDS": CriteriaViewDateFormat_EPOCH_SECONDS,
			"EPOCH_MILLIS": CriteriaViewDateFormat_EPOCH_MILLIS,
			"EPOCH_NANOSECONDS": CriteriaViewDateFormat_EPOCH_NANOSECONDS,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.CriteriaViewType",
		reflect.TypeOf((*CriteriaViewType)(nil)).Elem(),
		map[string]interface{}{
			"DATE": CriteriaViewType_DATE,
			"CUSTOM": CriteriaViewType_CUSTOM,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.CustomerX509",
		reflect.TypeOf((*CustomerX509)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeAwsCloudProviderConfigView",
		reflect.TypeOf((*DataLakeAwsCloudProviderConfigView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeCloudProviderConfigView",
		reflect.TypeOf((*DataLakeCloudProviderConfigView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeDataProcessRegionView",
		reflect.TypeOf((*DataLakeDataProcessRegionView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.DataLakeDataProcessRegionViewCloudProvider",
		reflect.TypeOf((*DataLakeDataProcessRegionViewCloudProvider)(nil)).Elem(),
		map[string]interface{}{
			"AWS": DataLakeDataProcessRegionViewCloudProvider_AWS,
			"GCP": DataLakeDataProcessRegionViewCloudProvider_GCP,
			"AZURE": DataLakeDataProcessRegionViewCloudProvider_AZURE,
			"TENANT": DataLakeDataProcessRegionViewCloudProvider_TENANT,
			"SERVERLESS": DataLakeDataProcessRegionViewCloudProvider_SERVERLESS,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.DataLakeDataProcessRegionViewRegion",
		reflect.TypeOf((*DataLakeDataProcessRegionViewRegion)(nil)).Elem(),
		map[string]interface{}{
			"DUBLIN_IRL": DataLakeDataProcessRegionViewRegion_DUBLIN_IRL,
			"FRANKFURT_DEU": DataLakeDataProcessRegionViewRegion_FRANKFURT_DEU,
			"LONDON_GBR": DataLakeDataProcessRegionViewRegion_LONDON_GBR,
			"MUMBAI_IND": DataLakeDataProcessRegionViewRegion_MUMBAI_IND,
			"OREGON_USA": DataLakeDataProcessRegionViewRegion_OREGON_USA,
			"SYDNEY_AUS": DataLakeDataProcessRegionViewRegion_SYDNEY_AUS,
			"VIRGINIA_USA": DataLakeDataProcessRegionViewRegion_VIRGINIA_USA,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeDatabaseCollectionView",
		reflect.TypeOf((*DataLakeDatabaseCollectionView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeDatabaseDataSourceView",
		reflect.TypeOf((*DataLakeDatabaseDataSourceView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.DataLakeDatabaseDataSourceViewDefaultFormat",
		reflect.TypeOf((*DataLakeDatabaseDataSourceViewDefaultFormat)(nil)).Elem(),
		map[string]interface{}{
			"AVRO": DataLakeDatabaseDataSourceViewDefaultFormat_AVRO,
			"AVRO_GZ": DataLakeDatabaseDataSourceViewDefaultFormat_AVRO_GZ,
			"BSON": DataLakeDatabaseDataSourceViewDefaultFormat_BSON,
			"BSON_GZ": DataLakeDatabaseDataSourceViewDefaultFormat_BSON_GZ,
			"CSV": DataLakeDatabaseDataSourceViewDefaultFormat_CSV,
			"JSON": DataLakeDatabaseDataSourceViewDefaultFormat_JSON,
			"JSON_GZ": DataLakeDatabaseDataSourceViewDefaultFormat_JSON_GZ,
			"ORC": DataLakeDatabaseDataSourceViewDefaultFormat_ORC,
			"TSV": DataLakeDatabaseDataSourceViewDefaultFormat_TSV,
			"TSV_GZ": DataLakeDatabaseDataSourceViewDefaultFormat_TSV_GZ,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeDatabaseView",
		reflect.TypeOf((*DataLakeDatabaseView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeStorageView",
		reflect.TypeOf((*DataLakeStorageView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DataLakeViewView",
		reflect.TypeOf((*DataLakeViewView)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DatabaseConfig",
		reflect.TypeOf((*DatabaseConfig)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.DatabaseConfigOperationTypes",
		reflect.TypeOf((*DatabaseConfigOperationTypes)(nil)).Elem(),
		map[string]interface{}{
			"INSERT": DatabaseConfigOperationTypes_INSERT,
			"UPDATE": DatabaseConfigOperationTypes_UPDATE,
			"REPLACE": DatabaseConfigOperationTypes_REPLACE,
			"DELETE": DatabaseConfigOperationTypes_DELETE,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DatabaseUserProps",
		reflect.TypeOf((*DatabaseUserProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.DatadogIntegration",
		reflect.TypeOf((*DatadogIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cfnThirdPartyIntegration", GoGetter: "CfnThirdPartyIntegration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DatadogIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DatadogIntegrationProps",
		reflect.TypeOf((*DatadogIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.DatadogRegion",
		reflect.TypeOf((*DatadogRegion)(nil)).Elem(),
		map[string]interface{}{
			"US": DatadogRegion_US,
			"EU": DatadogRegion_EU,
			"US3": DatadogRegion_US3,
			"US5": DatadogRegion_US5,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.DiskGb",
		reflect.TypeOf((*DiskGb)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.EncryptionAtRestProps",
		reflect.TypeOf((*EncryptionAtRestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Endpoint",
		reflect.TypeOf((*Endpoint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Event",
		reflect.TypeOf((*Event)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.EventAwseventbridge",
		reflect.TypeOf((*EventAwseventbridge)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.EventAwseventbridgeAwsConfig",
		reflect.TypeOf((*EventAwseventbridgeAwsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.EventFunction",
		reflect.TypeOf((*EventFunction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.EventFunctionFuncConfig",
		reflect.TypeOf((*EventFunctionFuncConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Export",
		reflect.TypeOf((*Export)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.InheritedRole",
		reflect.TypeOf((*InheritedRole)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.IntegerThresholdView",
		reflect.TypeOf((*IntegerThresholdView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.IntegerThresholdViewOperator",
		reflect.TypeOf((*IntegerThresholdViewOperator)(nil)).Elem(),
		map[string]interface{}{
			"GREATER_THAN": IntegerThresholdViewOperator_GREATER_THAN,
			"LESS_THAN": IntegerThresholdViewOperator_LESS_THAN,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.IpAccessListProps",
		reflect.TypeOf((*IpAccessListProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.LabelDefinition",
		reflect.TypeOf((*LabelDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Link",
		reflect.TypeOf((*Link)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ListOptions",
		reflect.TypeOf((*ListOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ManagedNamespace",
		reflect.TypeOf((*ManagedNamespace)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Matcher",
		reflect.TypeOf((*Matcher)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.MatcherFieldName",
		reflect.TypeOf((*MatcherFieldName)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER_NAME": MatcherFieldName_CLUSTER_NAME,
			"HOSTNAME": MatcherFieldName_HOSTNAME,
			"HOSTNAME_AND_PORT": MatcherFieldName_HOSTNAME_AND_PORT,
			"PORT": MatcherFieldName_PORT,
			"REPLICA_SET_NAME": MatcherFieldName_REPLICA_SET_NAME,
			"SHARD_NAME": MatcherFieldName_SHARD_NAME,
			"TYPE_NAME": MatcherFieldName_TYPE_NAME,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.MatcherOperator",
		reflect.TypeOf((*MatcherOperator)(nil)).Elem(),
		map[string]interface{}{
			"EQUALS": MatcherOperator_EQUALS,
			"CONTAINS": MatcherOperator_CONTAINS,
			"STARTS_WITH": MatcherOperator_STARTS_WITH,
			"ENDS_WITH": MatcherOperator_ENDS_WITH,
			"NOT_EQUALS": MatcherOperator_NOT_EQUALS,
			"NOT_CONTAINS": MatcherOperator_NOT_CONTAINS,
			"REGEX": MatcherOperator_REGEX,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.MetricThresholdView",
		reflect.TypeOf((*MetricThresholdView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.MetricThresholdViewMode",
		reflect.TypeOf((*MetricThresholdViewMode)(nil)).Elem(),
		map[string]interface{}{
			"AVERAGE": MetricThresholdViewMode_AVERAGE,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.MetricThresholdViewOperator",
		reflect.TypeOf((*MetricThresholdViewOperator)(nil)).Elem(),
		map[string]interface{}{
			"GREATER_THAN": MetricThresholdViewOperator_GREATER_THAN,
			"LESS_THAN": MetricThresholdViewOperator_LESS_THAN,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.MicrosoftTeamsIntegration",
		reflect.TypeOf((*MicrosoftTeamsIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cfnThirdPartyIntegration", GoGetter: "CfnThirdPartyIntegration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MicrosoftTeamsIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.MicrosoftTeamsIntegrationProps",
		reflect.TypeOf((*MicrosoftTeamsIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.NotificationView",
		reflect.TypeOf((*NotificationView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.NotificationViewDatadogRegion",
		reflect.TypeOf((*NotificationViewDatadogRegion)(nil)).Elem(),
		map[string]interface{}{
			"EU": NotificationViewDatadogRegion_EU,
			"US": NotificationViewDatadogRegion_US,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.NotificationViewOpsGenieRegion",
		reflect.TypeOf((*NotificationViewOpsGenieRegion)(nil)).Elem(),
		map[string]interface{}{
			"EU": NotificationViewOpsGenieRegion_EU,
			"US": NotificationViewOpsGenieRegion_US,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.NotificationViewRoles",
		reflect.TypeOf((*NotificationViewRoles)(nil)).Elem(),
		map[string]interface{}{
			"GROUP_CLUSTER_MANAGER": NotificationViewRoles_GROUP_CLUSTER_MANAGER,
			"GROUP_DATA_ACCESS_ADMIN": NotificationViewRoles_GROUP_DATA_ACCESS_ADMIN,
			"GROUP_DATA_ACCESS_READ_ONLY": NotificationViewRoles_GROUP_DATA_ACCESS_READ_ONLY,
			"GROUP_DATA_ACCESS_READ_WRITE": NotificationViewRoles_GROUP_DATA_ACCESS_READ_WRITE,
			"GROUP_OWNER": NotificationViewRoles_GROUP_OWNER,
			"GROUP_READ_WRITE": NotificationViewRoles_GROUP_READ_WRITE,
			"ORG_OWNER": NotificationViewRoles_ORG_OWNER,
			"ORG_MEMBER": NotificationViewRoles_ORG_MEMBER,
			"ORG_GROUP_CREATOR": NotificationViewRoles_ORG_GROUP_CREATOR,
			"ORG_BILLING_ADMIN": NotificationViewRoles_ORG_BILLING_ADMIN,
			"ORG_READ_ONLY": NotificationViewRoles_ORG_READ_ONLY,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.NotificationViewSeverity",
		reflect.TypeOf((*NotificationViewSeverity)(nil)).Elem(),
		map[string]interface{}{
			"CRITICAL": NotificationViewSeverity_CRITICAL,
			"ERROR": NotificationViewSeverity_ERROR,
			"WARNING": NotificationViewSeverity_WARNING,
		},
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.NotificationViewTypeName",
		reflect.TypeOf((*NotificationViewTypeName)(nil)).Elem(),
		map[string]interface{}{
			"DATADOG": NotificationViewTypeName_DATADOG,
			"EMAIL": NotificationViewTypeName_EMAIL,
			"FLOWDOCK": NotificationViewTypeName_FLOWDOCK,
			"GROUP": NotificationViewTypeName_GROUP,
			"MICROSOFT_TEAMS": NotificationViewTypeName_MICROSOFT_TEAMS,
			"OPS_GENIE": NotificationViewTypeName_OPS_GENIE,
			"ORG": NotificationViewTypeName_ORG,
			"PAGER_DUTY": NotificationViewTypeName_PAGER_DUTY,
			"PROMETHEUS": NotificationViewTypeName_PROMETHEUS,
			"SLACK": NotificationViewTypeName_SLACK,
			"SMS": NotificationViewTypeName_SMS,
			"TEAM": NotificationViewTypeName_TEAM,
			"USER": NotificationViewTypeName_USER,
			"VICTOR_OPS": NotificationViewTypeName_VICTOR_OPS,
			"WEBHOOK": NotificationViewTypeName_WEBHOOK,
		},
	)
	_jsii_.RegisterClass(
		"awscdk-resources-mongodbatlas.PagerDutyIntegration",
		reflect.TypeOf((*PagerDutyIntegration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cfnThirdPartyIntegration", GoGetter: "CfnThirdPartyIntegration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PagerDutyIntegration{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.PagerDutyIntegrationProps",
		reflect.TypeOf((*PagerDutyIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.PagerDutyRegion",
		reflect.TypeOf((*PagerDutyRegion)(nil)).Elem(),
		map[string]interface{}{
			"US": PagerDutyRegion_US,
			"EU": PagerDutyRegion_EU,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.PartitionFieldView",
		reflect.TypeOf((*PartitionFieldView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.PartitionFieldViewFieldType",
		reflect.TypeOf((*PartitionFieldViewFieldType)(nil)).Elem(),
		map[string]interface{}{
			"DATE": PartitionFieldViewFieldType_DATE,
			"INT": PartitionFieldViewFieldType_INT,
			"LONG": PartitionFieldViewFieldType_LONG,
			"OBJECT_ID": PartitionFieldViewFieldType_OBJECT_ID,
			"STRING": PartitionFieldViewFieldType_STRING,
			"UUID": PartitionFieldViewFieldType_UUID,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.PrivateEndpoint",
		reflect.TypeOf((*PrivateEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.PrivateEndpointProps",
		reflect.TypeOf((*PrivateEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ProcessArgs",
		reflect.TypeOf((*ProcessArgs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ProjectApiKey",
		reflect.TypeOf((*ProjectApiKey)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ProjectProps",
		reflect.TypeOf((*ProjectProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ProjectSettings",
		reflect.TypeOf((*ProjectSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ProjectTeam",
		reflect.TypeOf((*ProjectTeam)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Resource",
		reflect.TypeOf((*Resource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.RoleAssignment",
		reflect.TypeOf((*RoleAssignment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.RoleDefinition",
		reflect.TypeOf((*RoleDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ScheduleConfig",
		reflect.TypeOf((*ScheduleConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ScheduleView",
		reflect.TypeOf((*ScheduleView)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ScheduleViewType",
		reflect.TypeOf((*ScheduleViewType)(nil)).Elem(),
		map[string]interface{}{
			"DAILY": ScheduleViewType_DAILY,
			"MONTHLY": ScheduleViewType_MONTHLY,
			"DEFAULT": ScheduleViewType_DEFAULT,
			"WEEKLY": ScheduleViewType_WEEKLY,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ScopeDefinition",
		reflect.TypeOf((*ScopeDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ScopeDefinitionType",
		reflect.TypeOf((*ScopeDefinitionType)(nil)).Elem(),
		map[string]interface{}{
			"CLUSTER": ScopeDefinitionType_CLUSTER,
			"DATA_LAKE": ScopeDefinitionType_DATA_LAKE,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ServerlessInstanceConnectionStrings",
		reflect.TypeOf((*ServerlessInstanceConnectionStrings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ServerlessInstancePrivateEndpoint",
		reflect.TypeOf((*ServerlessInstancePrivateEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ServerlessInstancePrivateEndpointEndpoint",
		reflect.TypeOf((*ServerlessInstancePrivateEndpointEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ServerlessInstancePrivateEndpointType",
		reflect.TypeOf((*ServerlessInstancePrivateEndpointType)(nil)).Elem(),
		map[string]interface{}{
			"MONGOS": ServerlessInstancePrivateEndpointType_MONGOS,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ServerlessInstanceProviderSettings",
		reflect.TypeOf((*ServerlessInstanceProviderSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"awscdk-resources-mongodbatlas.ServerlessInstanceProviderSettingsProviderName",
		reflect.TypeOf((*ServerlessInstanceProviderSettingsProviderName)(nil)).Elem(),
		map[string]interface{}{
			"SERVERLESS": ServerlessInstanceProviderSettingsProviderName_SERVERLESS,
		},
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Specs",
		reflect.TypeOf((*Specs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.StoreDetail",
		reflect.TypeOf((*StoreDetail)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.SynonymSource",
		reflect.TypeOf((*SynonymSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ThirdPartyIntegrationProps",
		reflect.TypeOf((*ThirdPartyIntegrationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.Validation",
		reflect.TypeOf((*Validation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"awscdk-resources-mongodbatlas.ZoneMapping",
		reflect.TypeOf((*ZoneMapping)(nil)).Elem(),
	)
}
