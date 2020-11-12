/*
 * ChaosSearch API
 *
 * The ChaosSearch API is the administrative interface for the ChaosSearch service.  It is composed of ChaosSearch and AWS style interfaces (built to provide interoperability with the S3 service).  ChaosSearch admin API was originally modeled as an extension to the S3 API. Most of the API calls, parameters, etc will look familiar to those familiar with creating and managing S3 Buckets. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UserGroupsPermissions struct for UserGroupsPermissions
type UserGroupsPermissions struct {
	Effect *string `json:"Effect,omitempty"`
	Action *string `json:"Action,omitempty"`
	Resources *string `json:"Resources,omitempty"`
}

// NewUserGroupsPermissions instantiates a new UserGroupsPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupsPermissions() *UserGroupsPermissions {
	this := UserGroupsPermissions{}
	return &this
}

// NewUserGroupsPermissionsWithDefaults instantiates a new UserGroupsPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupsPermissionsWithDefaults() *UserGroupsPermissions {
	this := UserGroupsPermissions{}
	return &this
}

// GetEffect returns the Effect field value if set, zero value otherwise.
func (o *UserGroupsPermissions) GetEffect() string {
	if o == nil || o.Effect == nil {
		var ret string
		return ret
	}
	return *o.Effect
}

// GetEffectOk returns a tuple with the Effect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupsPermissions) GetEffectOk() (*string, bool) {
	if o == nil || o.Effect == nil {
		return nil, false
	}
	return o.Effect, true
}

// HasEffect returns a boolean if a field has been set.
func (o *UserGroupsPermissions) HasEffect() bool {
	if o != nil && o.Effect != nil {
		return true
	}

	return false
}

// SetEffect gets a reference to the given string and assigns it to the Effect field.
func (o *UserGroupsPermissions) SetEffect(v string) {
	o.Effect = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UserGroupsPermissions) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupsPermissions) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UserGroupsPermissions) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *UserGroupsPermissions) SetAction(v string) {
	o.Action = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *UserGroupsPermissions) GetResources() string {
	if o == nil || o.Resources == nil {
		var ret string
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupsPermissions) GetResourcesOk() (*string, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *UserGroupsPermissions) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given string and assigns it to the Resources field.
func (o *UserGroupsPermissions) SetResources(v string) {
	o.Resources = &v
}

func (o UserGroupsPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Effect != nil {
		toSerialize["Effect"] = o.Effect
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.Resources != nil {
		toSerialize["Resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableUserGroupsPermissions struct {
	value *UserGroupsPermissions
	isSet bool
}

func (v NullableUserGroupsPermissions) Get() *UserGroupsPermissions {
	return v.value
}

func (v *NullableUserGroupsPermissions) Set(val *UserGroupsPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupsPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupsPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupsPermissions(val *UserGroupsPermissions) *NullableUserGroupsPermissions {
	return &NullableUserGroupsPermissions{value: val, isSet: true}
}

func (v NullableUserGroupsPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupsPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

