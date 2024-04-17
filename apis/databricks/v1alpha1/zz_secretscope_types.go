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

type KeyvaultMetadataInitParameters struct {

	// Scope name requested by the user. Must be unique within a workspace. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// The id for the secret scope object.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type KeyvaultMetadataObservation struct {

	// Scope name requested by the user. Must be unique within a workspace. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// The id for the secret scope object.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type KeyvaultMetadataParameters struct {

	// Scope name requested by the user. Must be unique within a workspace. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.
	// +kubebuilder:validation:Optional
	DNSName *string `json:"dnsName" tf:"dns_name,omitempty"`

	// The id for the secret scope object.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`
}

type SecretScopeInitParameters struct {

	// Either DATABRICKS or AZURE_KEYVAULT
	BackendType *string `json:"backendType,omitempty" tf:"backend_type,omitempty"`

	// The principal with the only possible value users that is initially granted MANAGE permission to the created scope.  If it's omitted, then the databricks_secret_acl with MANAGE permission applied to the scope is assigned to the API request issuer's user identity (see documentation). This part of the state cannot be imported.
	InitialManagePrincipal *string `json:"initialManagePrincipal,omitempty" tf:"initial_manage_principal,omitempty"`

	KeyvaultMetadata []KeyvaultMetadataInitParameters `json:"keyvaultMetadata,omitempty" tf:"keyvault_metadata,omitempty"`

	// Scope name requested by the user. Must be unique within a workspace. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SecretScopeObservation struct {

	// Either DATABRICKS or AZURE_KEYVAULT
	BackendType *string `json:"backendType,omitempty" tf:"backend_type,omitempty"`

	// The id for the secret scope object.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The principal with the only possible value users that is initially granted MANAGE permission to the created scope.  If it's omitted, then the databricks_secret_acl with MANAGE permission applied to the scope is assigned to the API request issuer's user identity (see documentation). This part of the state cannot be imported.
	InitialManagePrincipal *string `json:"initialManagePrincipal,omitempty" tf:"initial_manage_principal,omitempty"`

	KeyvaultMetadata []KeyvaultMetadataObservation `json:"keyvaultMetadata,omitempty" tf:"keyvault_metadata,omitempty"`

	// Scope name requested by the user. Must be unique within a workspace. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type SecretScopeParameters struct {

	// Either DATABRICKS or AZURE_KEYVAULT
	// +kubebuilder:validation:Optional
	BackendType *string `json:"backendType,omitempty" tf:"backend_type,omitempty"`

	// The principal with the only possible value users that is initially granted MANAGE permission to the created scope.  If it's omitted, then the databricks_secret_acl with MANAGE permission applied to the scope is assigned to the API request issuer's user identity (see documentation). This part of the state cannot be imported.
	// +kubebuilder:validation:Optional
	InitialManagePrincipal *string `json:"initialManagePrincipal,omitempty" tf:"initial_manage_principal,omitempty"`

	// +kubebuilder:validation:Optional
	KeyvaultMetadata []KeyvaultMetadataParameters `json:"keyvaultMetadata,omitempty" tf:"keyvault_metadata,omitempty"`

	// Scope name requested by the user. Must be unique within a workspace. Must consist of alphanumeric characters, dashes, underscores, and periods, and may not exceed 128 characters.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// SecretScopeSpec defines the desired state of SecretScope
type SecretScopeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretScopeParameters `json:"forProvider"`
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
	InitProvider SecretScopeInitParameters `json:"initProvider,omitempty"`
}

// SecretScopeStatus defines the observed state of SecretScope.
type SecretScopeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretScopeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretScope is the Schema for the SecretScopes API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,databricks}
type SecretScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SecretScopeSpec   `json:"spec"`
	Status SecretScopeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretScopeList contains a list of SecretScopes
type SecretScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretScope `json:"items"`
}

// Repository type metadata.
var (
	SecretScope_Kind             = "SecretScope"
	SecretScope_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretScope_Kind}.String()
	SecretScope_KindAPIVersion   = SecretScope_Kind + "." + CRDGroupVersion.String()
	SecretScope_GroupVersionKind = CRDGroupVersion.WithKind(SecretScope_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretScope{}, &SecretScopeList{})
}
