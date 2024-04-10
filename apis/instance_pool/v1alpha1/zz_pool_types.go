// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AwsAttributesInitParameters struct {

	// Availability type used for all nodes. Valid values are SPOT_AZURE and ON_DEMAND_AZURE.
	Availability *string `json:"availability,omitempty" tf:"availability,omitempty"`

	// (Integer) The max price for AWS spot instances, as a percentage of the corresponding instance type’s on-demand price. For example, if this field is set to 50, and the instance pool needs a new i3.xlarge spot instance, then the max price is half of the price of on-demand i3.xlarge instances. Similarly, if this field is set to 200, the max price is twice the price of on-demand i3.xlarge instances. If not specified, the default value is 100. When spot instances are requested for this instance pool, only spot instances whose max price percentage matches this field are considered. For safety, this field cannot be greater than 10000.
	SpotBidPricePercent *float64 `json:"spotBidPricePercent,omitempty" tf:"spot_bid_price_percent,omitempty"`

	// (String) Identifier for the availability zone/datacenter in which the instance pool resides. This string is of the form like "us-west-2a". The provided availability zone must be in the same region as the Databricks deployment. For example, "us-west-2a" is not a valid zone ID if the Databricks deployment resides in the "us-east-1" region. If not specified, a default zone is used. You can find the list of available zones as well as the default value by using the List Zones API.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type AwsAttributesObservation struct {

	// Availability type used for all nodes. Valid values are SPOT_AZURE and ON_DEMAND_AZURE.
	Availability *string `json:"availability,omitempty" tf:"availability,omitempty"`

	// (Integer) The max price for AWS spot instances, as a percentage of the corresponding instance type’s on-demand price. For example, if this field is set to 50, and the instance pool needs a new i3.xlarge spot instance, then the max price is half of the price of on-demand i3.xlarge instances. Similarly, if this field is set to 200, the max price is twice the price of on-demand i3.xlarge instances. If not specified, the default value is 100. When spot instances are requested for this instance pool, only spot instances whose max price percentage matches this field are considered. For safety, this field cannot be greater than 10000.
	SpotBidPricePercent *float64 `json:"spotBidPricePercent,omitempty" tf:"spot_bid_price_percent,omitempty"`

	// (String) Identifier for the availability zone/datacenter in which the instance pool resides. This string is of the form like "us-west-2a". The provided availability zone must be in the same region as the Databricks deployment. For example, "us-west-2a" is not a valid zone ID if the Databricks deployment resides in the "us-east-1" region. If not specified, a default zone is used. You can find the list of available zones as well as the default value by using the List Zones API.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type AwsAttributesParameters struct {

	// Availability type used for all nodes. Valid values are SPOT_AZURE and ON_DEMAND_AZURE.
	// +kubebuilder:validation:Optional
	Availability *string `json:"availability,omitempty" tf:"availability,omitempty"`

	// (Integer) The max price for AWS spot instances, as a percentage of the corresponding instance type’s on-demand price. For example, if this field is set to 50, and the instance pool needs a new i3.xlarge spot instance, then the max price is half of the price of on-demand i3.xlarge instances. Similarly, if this field is set to 200, the max price is twice the price of on-demand i3.xlarge instances. If not specified, the default value is 100. When spot instances are requested for this instance pool, only spot instances whose max price percentage matches this field are considered. For safety, this field cannot be greater than 10000.
	// +kubebuilder:validation:Optional
	SpotBidPricePercent *float64 `json:"spotBidPricePercent,omitempty" tf:"spot_bid_price_percent,omitempty"`

	// (String) Identifier for the availability zone/datacenter in which the instance pool resides. This string is of the form like "us-west-2a". The provided availability zone must be in the same region as the Databricks deployment. For example, "us-west-2a" is not a valid zone ID if the Databricks deployment resides in the "us-east-1" region. If not specified, a default zone is used. You can find the list of available zones as well as the default value by using the List Zones API.
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type AzureAttributesInitParameters struct {

	// Availability type used for all nodes. Valid values are SPOT_AZURE and ON_DEMAND_AZURE.
	Availability *string `json:"availability,omitempty" tf:"availability,omitempty"`

	// The max price for Azure spot instances.  Use -1 to specify the lowest price.
	SpotBidMaxPrice *float64 `json:"spotBidMaxPrice,omitempty" tf:"spot_bid_max_price,omitempty"`
}

type AzureAttributesObservation struct {

	// Availability type used for all nodes. Valid values are SPOT_AZURE and ON_DEMAND_AZURE.
	Availability *string `json:"availability,omitempty" tf:"availability,omitempty"`

	// The max price for Azure spot instances.  Use -1 to specify the lowest price.
	SpotBidMaxPrice *float64 `json:"spotBidMaxPrice,omitempty" tf:"spot_bid_max_price,omitempty"`
}

type AzureAttributesParameters struct {

	// Availability type used for all nodes. Valid values are SPOT_AZURE and ON_DEMAND_AZURE.
	// +kubebuilder:validation:Optional
	Availability *string `json:"availability,omitempty" tf:"availability,omitempty"`

	// The max price for Azure spot instances.  Use -1 to specify the lowest price.
	// +kubebuilder:validation:Optional
	SpotBidMaxPrice *float64 `json:"spotBidMaxPrice,omitempty" tf:"spot_bid_max_price,omitempty"`
}

type BasicAuthInitParameters struct {
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type BasicAuthObservation struct {
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type BasicAuthParameters struct {

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type DiskSpecInitParameters struct {

	// (Integer) The number of disks to attach to each instance. This feature is only enabled for supported node types. Users can choose up to the limit of the disks supported by the node type. For node types with no local disk, at least one disk needs to be specified.
	DiskCount *float64 `json:"diskCount,omitempty" tf:"disk_count,omitempty"`

	// (Integer) The size of each disk (in GiB) to attach.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	DiskType []DiskTypeInitParameters `json:"diskType,omitempty" tf:"disk_type,omitempty"`
}

type DiskSpecObservation struct {

	// (Integer) The number of disks to attach to each instance. This feature is only enabled for supported node types. Users can choose up to the limit of the disks supported by the node type. For node types with no local disk, at least one disk needs to be specified.
	DiskCount *float64 `json:"diskCount,omitempty" tf:"disk_count,omitempty"`

	// (Integer) The size of each disk (in GiB) to attach.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	DiskType []DiskTypeObservation `json:"diskType,omitempty" tf:"disk_type,omitempty"`
}

type DiskSpecParameters struct {

	// (Integer) The number of disks to attach to each instance. This feature is only enabled for supported node types. Users can choose up to the limit of the disks supported by the node type. For node types with no local disk, at least one disk needs to be specified.
	// +kubebuilder:validation:Optional
	DiskCount *float64 `json:"diskCount,omitempty" tf:"disk_count,omitempty"`

	// (Integer) The size of each disk (in GiB) to attach.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Optional
	DiskType []DiskTypeParameters `json:"diskType,omitempty" tf:"disk_type,omitempty"`
}

type DiskTypeInitParameters struct {
	AzureDiskVolumeType *string `json:"azureDiskVolumeType,omitempty" tf:"azure_disk_volume_type,omitempty"`

	EBSVolumeType *string `json:"ebsVolumeType,omitempty" tf:"ebs_volume_type,omitempty"`
}

type DiskTypeObservation struct {
	AzureDiskVolumeType *string `json:"azureDiskVolumeType,omitempty" tf:"azure_disk_volume_type,omitempty"`

	EBSVolumeType *string `json:"ebsVolumeType,omitempty" tf:"ebs_volume_type,omitempty"`
}

type DiskTypeParameters struct {

	// +kubebuilder:validation:Optional
	AzureDiskVolumeType *string `json:"azureDiskVolumeType,omitempty" tf:"azure_disk_volume_type,omitempty"`

	// +kubebuilder:validation:Optional
	EBSVolumeType *string `json:"ebsVolumeType,omitempty" tf:"ebs_volume_type,omitempty"`
}

type FleetOnDemandOptionInitParameters struct {
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`

	InstancePoolsToUseCount *float64 `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count,omitempty"`
}

type FleetOnDemandOptionObservation struct {
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`

	InstancePoolsToUseCount *float64 `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count,omitempty"`
}

type FleetOnDemandOptionParameters struct {

	// +kubebuilder:validation:Optional
	AllocationStrategy *string `json:"allocationStrategy" tf:"allocation_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	InstancePoolsToUseCount *float64 `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count,omitempty"`
}

type FleetSpotOptionInitParameters struct {
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`

	InstancePoolsToUseCount *float64 `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count,omitempty"`
}

type FleetSpotOptionObservation struct {
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy,omitempty"`

	InstancePoolsToUseCount *float64 `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count,omitempty"`
}

type FleetSpotOptionParameters struct {

	// +kubebuilder:validation:Optional
	AllocationStrategy *string `json:"allocationStrategy" tf:"allocation_strategy,omitempty"`

	// +kubebuilder:validation:Optional
	InstancePoolsToUseCount *float64 `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count,omitempty"`
}

type GCPAttributesInitParameters struct {

	// Availability type used for all nodes. Valid values are PREEMPTIBLE_GCP, PREEMPTIBLE_WITH_FALLBACK_GCP and ON_DEMAND_GCP, default: ON_DEMAND_GCP.
	GCPAvailability *string `json:"gcpAvailability,omitempty" tf:"gcp_availability,omitempty"`

	// Number of local SSD disks (each is 375GB in size) that will be attached to each node of the cluster.
	LocalSsdCount *float64 `json:"localSsdCount,omitempty" tf:"local_ssd_count,omitempty"`
}

type GCPAttributesObservation struct {

	// Availability type used for all nodes. Valid values are PREEMPTIBLE_GCP, PREEMPTIBLE_WITH_FALLBACK_GCP and ON_DEMAND_GCP, default: ON_DEMAND_GCP.
	GCPAvailability *string `json:"gcpAvailability,omitempty" tf:"gcp_availability,omitempty"`

	// Number of local SSD disks (each is 375GB in size) that will be attached to each node of the cluster.
	LocalSsdCount *float64 `json:"localSsdCount,omitempty" tf:"local_ssd_count,omitempty"`
}

type GCPAttributesParameters struct {

	// Availability type used for all nodes. Valid values are PREEMPTIBLE_GCP, PREEMPTIBLE_WITH_FALLBACK_GCP and ON_DEMAND_GCP, default: ON_DEMAND_GCP.
	// +kubebuilder:validation:Optional
	GCPAvailability *string `json:"gcpAvailability,omitempty" tf:"gcp_availability,omitempty"`

	// Number of local SSD disks (each is 375GB in size) that will be attached to each node of the cluster.
	// +kubebuilder:validation:Optional
	LocalSsdCount *float64 `json:"localSsdCount,omitempty" tf:"local_ssd_count,omitempty"`
}

type InstancePoolFleetAttributesInitParameters struct {
	FleetOnDemandOption []FleetOnDemandOptionInitParameters `json:"fleetOnDemandOption,omitempty" tf:"fleet_on_demand_option,omitempty"`

	FleetSpotOption []FleetSpotOptionInitParameters `json:"fleetSpotOption,omitempty" tf:"fleet_spot_option,omitempty"`

	LaunchTemplateOverride []LaunchTemplateOverrideInitParameters `json:"launchTemplateOverride,omitempty" tf:"launch_template_override,omitempty"`
}

type InstancePoolFleetAttributesObservation struct {
	FleetOnDemandOption []FleetOnDemandOptionObservation `json:"fleetOnDemandOption,omitempty" tf:"fleet_on_demand_option,omitempty"`

	FleetSpotOption []FleetSpotOptionObservation `json:"fleetSpotOption,omitempty" tf:"fleet_spot_option,omitempty"`

	LaunchTemplateOverride []LaunchTemplateOverrideObservation `json:"launchTemplateOverride,omitempty" tf:"launch_template_override,omitempty"`
}

type InstancePoolFleetAttributesParameters struct {

	// +kubebuilder:validation:Optional
	FleetOnDemandOption []FleetOnDemandOptionParameters `json:"fleetOnDemandOption,omitempty" tf:"fleet_on_demand_option,omitempty"`

	// +kubebuilder:validation:Optional
	FleetSpotOption []FleetSpotOptionParameters `json:"fleetSpotOption,omitempty" tf:"fleet_spot_option,omitempty"`

	// +kubebuilder:validation:Optional
	LaunchTemplateOverride []LaunchTemplateOverrideParameters `json:"launchTemplateOverride" tf:"launch_template_override,omitempty"`
}

type LaunchTemplateOverrideInitParameters struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`
}

type LaunchTemplateOverrideObservation struct {
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`
}

type LaunchTemplateOverrideParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`
}

type PoolInitParameters struct {
	AwsAttributes []AwsAttributesInitParameters `json:"awsAttributes,omitempty" tf:"aws_attributes,omitempty"`

	AzureAttributes []AzureAttributesInitParameters `json:"azureAttributes,omitempty" tf:"azure_attributes,omitempty"`

	// (Map) Additional tags for instance pool resources. Databricks tags all pool resources (e.g. AWS & Azure instances and Disk volumes). The tags of the instance pool will propagate to the clusters using the pool (see the official documentation). Attempting to set the same tags in both cluster and instance pool will raise an error. Databricks allows at most 43 custom tags.
	CustomTags map[string]*string `json:"customTags,omitempty" tf:"custom_tags,omitempty"`

	DiskSpec []DiskSpecInitParameters `json:"diskSpec,omitempty" tf:"disk_spec,omitempty"`

	// (Bool) Autoscaling Local Storage: when enabled, the instances in the pool dynamically acquire additional disk space when they are running low on disk space.
	EnableElasticDisk *bool `json:"enableElasticDisk,omitempty" tf:"enable_elastic_disk,omitempty"`

	GCPAttributes []GCPAttributesInitParameters `json:"gcpAttributes,omitempty" tf:"gcp_attributes,omitempty"`

	// (Integer) The number of minutes that idle instances in excess of the min_idle_instances are maintained by the pool before being terminated. If not specified, excess idle instances are terminated automatically after a default timeout period. If specified, the time must be between 0 and 10000 minutes. If you specify 0, excess idle instances are removed as soon as possible.
	IdleInstanceAutoterminationMinutes *float64 `json:"idleInstanceAutoterminationMinutes,omitempty" tf:"idle_instance_autotermination_minutes,omitempty"`

	InstancePoolFleetAttributes []InstancePoolFleetAttributesInitParameters `json:"instancePoolFleetAttributes,omitempty" tf:"instance_pool_fleet_attributes,omitempty"`

	// Canonical unique identifier for the instance pool.
	InstancePoolID *string `json:"instancePoolId,omitempty" tf:"instance_pool_id,omitempty"`

	// (String) The name of the instance pool. This is required for create and edit operations. It must be unique, non-empty, and less than 100 characters.
	InstancePoolName *string `json:"instancePoolName,omitempty" tf:"instance_pool_name,omitempty"`

	// (Integer) The maximum number of instances the pool can contain, including both idle instances and ones in use by clusters. Once the maximum capacity is reached, you cannot create new clusters from the pool and existing clusters cannot autoscale up until some instances are made idle in the pool via cluster termination or down-scaling. There is no default limit, but as a best practice, this should be set based on anticipated usage.
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// (Integer) The minimum number of idle instances maintained by the pool. This is in addition to any instances in use by active clusters.
	MinIdleInstances *float64 `json:"minIdleInstances,omitempty" tf:"min_idle_instances,omitempty"`

	// (String) The node type for the instances in the pool. All clusters attached to the pool inherit this node type and the pool’s idle instances are allocated based on this type. You can retrieve a list of available node types by using the List Node Types API call.
	NodeTypeID *string `json:"nodeTypeId,omitempty" tf:"node_type_id,omitempty"`

	PreloadedDockerImage []PreloadedDockerImageInitParameters `json:"preloadedDockerImage,omitempty" tf:"preloaded_docker_image,omitempty"`

	// (List) A list with at most one runtime version the pool installs on each instance. Pool clusters that use a preloaded runtime version start faster as they do not have to wait for the image to download. You can retrieve them via databricks_spark_version data source or via  Runtime Versions API call.
	PreloadedSparkVersions []*string `json:"preloadedSparkVersions,omitempty" tf:"preloaded_spark_versions,omitempty"`
}

type PoolObservation struct {
	AwsAttributes []AwsAttributesObservation `json:"awsAttributes,omitempty" tf:"aws_attributes,omitempty"`

	AzureAttributes []AzureAttributesObservation `json:"azureAttributes,omitempty" tf:"azure_attributes,omitempty"`

	// (Map) Additional tags for instance pool resources. Databricks tags all pool resources (e.g. AWS & Azure instances and Disk volumes). The tags of the instance pool will propagate to the clusters using the pool (see the official documentation). Attempting to set the same tags in both cluster and instance pool will raise an error. Databricks allows at most 43 custom tags.
	CustomTags map[string]*string `json:"customTags,omitempty" tf:"custom_tags,omitempty"`

	DiskSpec []DiskSpecObservation `json:"diskSpec,omitempty" tf:"disk_spec,omitempty"`

	// (Bool) Autoscaling Local Storage: when enabled, the instances in the pool dynamically acquire additional disk space when they are running low on disk space.
	EnableElasticDisk *bool `json:"enableElasticDisk,omitempty" tf:"enable_elastic_disk,omitempty"`

	GCPAttributes []GCPAttributesObservation `json:"gcpAttributes,omitempty" tf:"gcp_attributes,omitempty"`

	// Canonical unique identifier for the instance pool.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Integer) The number of minutes that idle instances in excess of the min_idle_instances are maintained by the pool before being terminated. If not specified, excess idle instances are terminated automatically after a default timeout period. If specified, the time must be between 0 and 10000 minutes. If you specify 0, excess idle instances are removed as soon as possible.
	IdleInstanceAutoterminationMinutes *float64 `json:"idleInstanceAutoterminationMinutes,omitempty" tf:"idle_instance_autotermination_minutes,omitempty"`

	InstancePoolFleetAttributes []InstancePoolFleetAttributesObservation `json:"instancePoolFleetAttributes,omitempty" tf:"instance_pool_fleet_attributes,omitempty"`

	// Canonical unique identifier for the instance pool.
	InstancePoolID *string `json:"instancePoolId,omitempty" tf:"instance_pool_id,omitempty"`

	// (String) The name of the instance pool. This is required for create and edit operations. It must be unique, non-empty, and less than 100 characters.
	InstancePoolName *string `json:"instancePoolName,omitempty" tf:"instance_pool_name,omitempty"`

	// (Integer) The maximum number of instances the pool can contain, including both idle instances and ones in use by clusters. Once the maximum capacity is reached, you cannot create new clusters from the pool and existing clusters cannot autoscale up until some instances are made idle in the pool via cluster termination or down-scaling. There is no default limit, but as a best practice, this should be set based on anticipated usage.
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// (Integer) The minimum number of idle instances maintained by the pool. This is in addition to any instances in use by active clusters.
	MinIdleInstances *float64 `json:"minIdleInstances,omitempty" tf:"min_idle_instances,omitempty"`

	// (String) The node type for the instances in the pool. All clusters attached to the pool inherit this node type and the pool’s idle instances are allocated based on this type. You can retrieve a list of available node types by using the List Node Types API call.
	NodeTypeID *string `json:"nodeTypeId,omitempty" tf:"node_type_id,omitempty"`

	PreloadedDockerImage []PreloadedDockerImageObservation `json:"preloadedDockerImage,omitempty" tf:"preloaded_docker_image,omitempty"`

	// (List) A list with at most one runtime version the pool installs on each instance. Pool clusters that use a preloaded runtime version start faster as they do not have to wait for the image to download. You can retrieve them via databricks_spark_version data source or via  Runtime Versions API call.
	PreloadedSparkVersions []*string `json:"preloadedSparkVersions,omitempty" tf:"preloaded_spark_versions,omitempty"`
}

type PoolParameters struct {

	// +kubebuilder:validation:Optional
	AwsAttributes []AwsAttributesParameters `json:"awsAttributes,omitempty" tf:"aws_attributes,omitempty"`

	// +kubebuilder:validation:Optional
	AzureAttributes []AzureAttributesParameters `json:"azureAttributes,omitempty" tf:"azure_attributes,omitempty"`

	// (Map) Additional tags for instance pool resources. Databricks tags all pool resources (e.g. AWS & Azure instances and Disk volumes). The tags of the instance pool will propagate to the clusters using the pool (see the official documentation). Attempting to set the same tags in both cluster and instance pool will raise an error. Databricks allows at most 43 custom tags.
	// +kubebuilder:validation:Optional
	CustomTags map[string]*string `json:"customTags,omitempty" tf:"custom_tags,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSpec []DiskSpecParameters `json:"diskSpec,omitempty" tf:"disk_spec,omitempty"`

	// (Bool) Autoscaling Local Storage: when enabled, the instances in the pool dynamically acquire additional disk space when they are running low on disk space.
	// +kubebuilder:validation:Optional
	EnableElasticDisk *bool `json:"enableElasticDisk,omitempty" tf:"enable_elastic_disk,omitempty"`

	// +kubebuilder:validation:Optional
	GCPAttributes []GCPAttributesParameters `json:"gcpAttributes,omitempty" tf:"gcp_attributes,omitempty"`

	// (Integer) The number of minutes that idle instances in excess of the min_idle_instances are maintained by the pool before being terminated. If not specified, excess idle instances are terminated automatically after a default timeout period. If specified, the time must be between 0 and 10000 minutes. If you specify 0, excess idle instances are removed as soon as possible.
	// +kubebuilder:validation:Optional
	IdleInstanceAutoterminationMinutes *float64 `json:"idleInstanceAutoterminationMinutes,omitempty" tf:"idle_instance_autotermination_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	InstancePoolFleetAttributes []InstancePoolFleetAttributesParameters `json:"instancePoolFleetAttributes,omitempty" tf:"instance_pool_fleet_attributes,omitempty"`

	// Canonical unique identifier for the instance pool.
	// +kubebuilder:validation:Optional
	InstancePoolID *string `json:"instancePoolId,omitempty" tf:"instance_pool_id,omitempty"`

	// (String) The name of the instance pool. This is required for create and edit operations. It must be unique, non-empty, and less than 100 characters.
	// +kubebuilder:validation:Optional
	InstancePoolName *string `json:"instancePoolName,omitempty" tf:"instance_pool_name,omitempty"`

	// (Integer) The maximum number of instances the pool can contain, including both idle instances and ones in use by clusters. Once the maximum capacity is reached, you cannot create new clusters from the pool and existing clusters cannot autoscale up until some instances are made idle in the pool via cluster termination or down-scaling. There is no default limit, but as a best practice, this should be set based on anticipated usage.
	// +kubebuilder:validation:Optional
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// (Integer) The minimum number of idle instances maintained by the pool. This is in addition to any instances in use by active clusters.
	// +kubebuilder:validation:Optional
	MinIdleInstances *float64 `json:"minIdleInstances,omitempty" tf:"min_idle_instances,omitempty"`

	// (String) The node type for the instances in the pool. All clusters attached to the pool inherit this node type and the pool’s idle instances are allocated based on this type. You can retrieve a list of available node types by using the List Node Types API call.
	// +kubebuilder:validation:Optional
	NodeTypeID *string `json:"nodeTypeId,omitempty" tf:"node_type_id,omitempty"`

	// +kubebuilder:validation:Optional
	PreloadedDockerImage []PreloadedDockerImageParameters `json:"preloadedDockerImage,omitempty" tf:"preloaded_docker_image,omitempty"`

	// (List) A list with at most one runtime version the pool installs on each instance. Pool clusters that use a preloaded runtime version start faster as they do not have to wait for the image to download. You can retrieve them via databricks_spark_version data source or via  Runtime Versions API call.
	// +kubebuilder:validation:Optional
	PreloadedSparkVersions []*string `json:"preloadedSparkVersions,omitempty" tf:"preloaded_spark_versions,omitempty"`
}

type PreloadedDockerImageInitParameters struct {

	// basic_auth.username and basic_auth.password for Docker repository. Docker registry credentials are encrypted when they are stored in Databricks internal storage and when they are passed to a registry upon fetching Docker images at cluster launch. However, other authenticated and authorized API users of this workspace can access the username and password.
	BasicAuth []BasicAuthInitParameters `json:"basicAuth,omitempty" tf:"basic_auth,omitempty"`

	// URL for the Docker image
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type PreloadedDockerImageObservation struct {

	// basic_auth.username and basic_auth.password for Docker repository. Docker registry credentials are encrypted when they are stored in Databricks internal storage and when they are passed to a registry upon fetching Docker images at cluster launch. However, other authenticated and authorized API users of this workspace can access the username and password.
	BasicAuth []BasicAuthObservation `json:"basicAuth,omitempty" tf:"basic_auth,omitempty"`

	// URL for the Docker image
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type PreloadedDockerImageParameters struct {

	// basic_auth.username and basic_auth.password for Docker repository. Docker registry credentials are encrypted when they are stored in Databricks internal storage and when they are passed to a registry upon fetching Docker images at cluster launch. However, other authenticated and authorized API users of this workspace can access the username and password.
	// +kubebuilder:validation:Optional
	BasicAuth []BasicAuthParameters `json:"basicAuth,omitempty" tf:"basic_auth,omitempty"`

	// URL for the Docker image
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`
}

// PoolSpec defines the desired state of Pool
type PoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PoolParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PoolInitParameters `json:"initProvider,omitempty"`
}

// PoolStatus defines the observed state of Pool.
type PoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Pool is the Schema for the Pools API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type Pool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.idleInstanceAutoterminationMinutes) || (has(self.initProvider) && has(self.initProvider.idleInstanceAutoterminationMinutes))",message="spec.forProvider.idleInstanceAutoterminationMinutes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instancePoolName) || (has(self.initProvider) && has(self.initProvider.instancePoolName))",message="spec.forProvider.instancePoolName is a required parameter"
	Spec   PoolSpec   `json:"spec"`
	Status PoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PoolList contains a list of Pools
type PoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pool `json:"items"`
}

// Repository type metadata.
var (
	Pool_Kind             = "Pool"
	Pool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pool_Kind}.String()
	Pool_KindAPIVersion   = Pool_Kind + "." + CRDGroupVersion.String()
	Pool_GroupVersionKind = CRDGroupVersion.WithKind(Pool_Kind)
)

func init() {
	SchemeBuilder.Register(&Pool{}, &PoolList{})
}
