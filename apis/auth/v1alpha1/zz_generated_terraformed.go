/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ConfigActivedirectory
func (mg *ConfigActivedirectory) GetTerraformResourceType() string {
	return "rancher2_auth_config_activedirectory"
}

// GetConnectionDetailsMapping for this ConfigActivedirectory
func (tr *ConfigActivedirectory) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"certificate": "spec.forProvider.certificateSecretRef", "service_account_password": "spec.forProvider.serviceAccountPasswordSecretRef", "service_account_username": "spec.forProvider.serviceAccountUsernameSecretRef", "test_password": "spec.forProvider.testPasswordSecretRef"}
}

// GetObservation of this ConfigActivedirectory
func (tr *ConfigActivedirectory) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConfigActivedirectory
func (tr *ConfigActivedirectory) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConfigActivedirectory
func (tr *ConfigActivedirectory) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConfigActivedirectory
func (tr *ConfigActivedirectory) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConfigActivedirectory
func (tr *ConfigActivedirectory) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ConfigActivedirectory using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConfigActivedirectory) LateInitialize(attrs []byte) (bool, error) {
	params := &ConfigActivedirectoryParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConfigActivedirectory) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ConfigAdfs
func (mg *ConfigAdfs) GetTerraformResourceType() string {
	return "rancher2_auth_config_adfs"
}

// GetConnectionDetailsMapping for this ConfigAdfs
func (tr *ConfigAdfs) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"idp_metadata_content": "spec.forProvider.idpMetadataContentSecretRef", "sp_cert": "spec.forProvider.spCertSecretRef", "sp_key": "spec.forProvider.spKeySecretRef"}
}

// GetObservation of this ConfigAdfs
func (tr *ConfigAdfs) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConfigAdfs
func (tr *ConfigAdfs) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConfigAdfs
func (tr *ConfigAdfs) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConfigAdfs
func (tr *ConfigAdfs) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConfigAdfs
func (tr *ConfigAdfs) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ConfigAdfs using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConfigAdfs) LateInitialize(attrs []byte) (bool, error) {
	params := &ConfigAdfsParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConfigAdfs) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ConfigAzuread
func (mg *ConfigAzuread) GetTerraformResourceType() string {
	return "rancher2_auth_config_azuread"
}

// GetConnectionDetailsMapping for this ConfigAzuread
func (tr *ConfigAzuread) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"application_id": "spec.forProvider.applicationIdSecretRef", "application_secret": "spec.forProvider.applicationSecretSecretRef"}
}

// GetObservation of this ConfigAzuread
func (tr *ConfigAzuread) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConfigAzuread
func (tr *ConfigAzuread) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConfigAzuread
func (tr *ConfigAzuread) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConfigAzuread
func (tr *ConfigAzuread) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConfigAzuread
func (tr *ConfigAzuread) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ConfigAzuread using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConfigAzuread) LateInitialize(attrs []byte) (bool, error) {
	params := &ConfigAzureadParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConfigAzuread) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ConfigFreeipa
func (mg *ConfigFreeipa) GetTerraformResourceType() string {
	return "rancher2_auth_config_freeipa"
}

// GetConnectionDetailsMapping for this ConfigFreeipa
func (tr *ConfigFreeipa) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"certificate": "spec.forProvider.certificateSecretRef", "service_account_distinguished_name": "spec.forProvider.serviceAccountDistinguishedNameSecretRef", "service_account_password": "spec.forProvider.serviceAccountPasswordSecretRef", "test_password": "spec.forProvider.testPasswordSecretRef"}
}

// GetObservation of this ConfigFreeipa
func (tr *ConfigFreeipa) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConfigFreeipa
func (tr *ConfigFreeipa) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConfigFreeipa
func (tr *ConfigFreeipa) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConfigFreeipa
func (tr *ConfigFreeipa) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConfigFreeipa
func (tr *ConfigFreeipa) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ConfigFreeipa using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConfigFreeipa) LateInitialize(attrs []byte) (bool, error) {
	params := &ConfigFreeipaParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConfigFreeipa) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ConfigGithub
func (mg *ConfigGithub) GetTerraformResourceType() string {
	return "rancher2_auth_config_github"
}

// GetConnectionDetailsMapping for this ConfigGithub
func (tr *ConfigGithub) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"client_id": "spec.forProvider.clientIdSecretRef", "client_secret": "spec.forProvider.clientSecretSecretRef"}
}

// GetObservation of this ConfigGithub
func (tr *ConfigGithub) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConfigGithub
func (tr *ConfigGithub) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConfigGithub
func (tr *ConfigGithub) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConfigGithub
func (tr *ConfigGithub) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConfigGithub
func (tr *ConfigGithub) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ConfigGithub using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConfigGithub) LateInitialize(attrs []byte) (bool, error) {
	params := &ConfigGithubParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConfigGithub) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ConfigKeycloak
func (mg *ConfigKeycloak) GetTerraformResourceType() string {
	return "rancher2_auth_config_keycloak"
}

// GetConnectionDetailsMapping for this ConfigKeycloak
func (tr *ConfigKeycloak) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"idp_metadata_content": "spec.forProvider.idpMetadataContentSecretRef", "sp_cert": "spec.forProvider.spCertSecretRef", "sp_key": "spec.forProvider.spKeySecretRef"}
}

// GetObservation of this ConfigKeycloak
func (tr *ConfigKeycloak) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConfigKeycloak
func (tr *ConfigKeycloak) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConfigKeycloak
func (tr *ConfigKeycloak) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConfigKeycloak
func (tr *ConfigKeycloak) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConfigKeycloak
func (tr *ConfigKeycloak) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ConfigKeycloak using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConfigKeycloak) LateInitialize(attrs []byte) (bool, error) {
	params := &ConfigKeycloakParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConfigKeycloak) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ConfigOkta
func (mg *ConfigOkta) GetTerraformResourceType() string {
	return "rancher2_auth_config_okta"
}

// GetConnectionDetailsMapping for this ConfigOkta
func (tr *ConfigOkta) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"idp_metadata_content": "spec.forProvider.idpMetadataContentSecretRef", "sp_cert": "spec.forProvider.spCertSecretRef", "sp_key": "spec.forProvider.spKeySecretRef"}
}

// GetObservation of this ConfigOkta
func (tr *ConfigOkta) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConfigOkta
func (tr *ConfigOkta) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConfigOkta
func (tr *ConfigOkta) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConfigOkta
func (tr *ConfigOkta) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConfigOkta
func (tr *ConfigOkta) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ConfigOkta using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConfigOkta) LateInitialize(attrs []byte) (bool, error) {
	params := &ConfigOktaParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConfigOkta) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ConfigOpenldap
func (mg *ConfigOpenldap) GetTerraformResourceType() string {
	return "rancher2_auth_config_openldap"
}

// GetConnectionDetailsMapping for this ConfigOpenldap
func (tr *ConfigOpenldap) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"certificate": "spec.forProvider.certificateSecretRef", "service_account_distinguished_name": "spec.forProvider.serviceAccountDistinguishedNameSecretRef", "service_account_password": "spec.forProvider.serviceAccountPasswordSecretRef", "test_password": "spec.forProvider.testPasswordSecretRef"}
}

// GetObservation of this ConfigOpenldap
func (tr *ConfigOpenldap) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConfigOpenldap
func (tr *ConfigOpenldap) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConfigOpenldap
func (tr *ConfigOpenldap) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConfigOpenldap
func (tr *ConfigOpenldap) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConfigOpenldap
func (tr *ConfigOpenldap) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ConfigOpenldap using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConfigOpenldap) LateInitialize(attrs []byte) (bool, error) {
	params := &ConfigOpenldapParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConfigOpenldap) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ConfigPing
func (mg *ConfigPing) GetTerraformResourceType() string {
	return "rancher2_auth_config_ping"
}

// GetConnectionDetailsMapping for this ConfigPing
func (tr *ConfigPing) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"idp_metadata_content": "spec.forProvider.idpMetadataContentSecretRef", "sp_cert": "spec.forProvider.spCertSecretRef", "sp_key": "spec.forProvider.spKeySecretRef"}
}

// GetObservation of this ConfigPing
func (tr *ConfigPing) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ConfigPing
func (tr *ConfigPing) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ConfigPing
func (tr *ConfigPing) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ConfigPing
func (tr *ConfigPing) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ConfigPing
func (tr *ConfigPing) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ConfigPing using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ConfigPing) LateInitialize(attrs []byte) (bool, error) {
	params := &ConfigPingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ConfigPing) GetTerraformSchemaVersion() int {
	return 0
}
