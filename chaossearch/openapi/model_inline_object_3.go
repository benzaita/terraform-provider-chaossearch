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

// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	// View Name
	Bucket *string `json:"bucket,omitempty"`
	IndexRetention *ViewTimeWindow `json:"indexRetention,omitempty"`
}

// NewInlineObject3 instantiates a new InlineObject3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject3() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// NewInlineObject3WithDefaults instantiates a new InlineObject3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject3WithDefaults() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *InlineObject3) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *InlineObject3) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *InlineObject3) SetBucket(v string) {
	o.Bucket = &v
}

// GetIndexRetention returns the IndexRetention field value if set, zero value otherwise.
func (o *InlineObject3) GetIndexRetention() ViewTimeWindow {
	if o == nil || o.IndexRetention == nil {
		var ret ViewTimeWindow
		return ret
	}
	return *o.IndexRetention
}

// GetIndexRetentionOk returns a tuple with the IndexRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetIndexRetentionOk() (*ViewTimeWindow, bool) {
	if o == nil || o.IndexRetention == nil {
		return nil, false
	}
	return o.IndexRetention, true
}

// HasIndexRetention returns a boolean if a field has been set.
func (o *InlineObject3) HasIndexRetention() bool {
	if o != nil && o.IndexRetention != nil {
		return true
	}

	return false
}

// SetIndexRetention gets a reference to the given ViewTimeWindow and assigns it to the IndexRetention field.
func (o *InlineObject3) SetIndexRetention(v ViewTimeWindow) {
	o.IndexRetention = &v
}

func (o InlineObject3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.IndexRetention != nil {
		toSerialize["indexRetention"] = o.IndexRetention
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject3 struct {
	value *InlineObject3
	isSet bool
}

func (v NullableInlineObject3) Get() *InlineObject3 {
	return v.value
}

func (v *NullableInlineObject3) Set(val *InlineObject3) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject3) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject3(val *InlineObject3) *NullableInlineObject3 {
	return &NullableInlineObject3{value: val, isSet: true}
}

func (v NullableInlineObject3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


