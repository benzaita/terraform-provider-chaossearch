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

// InlineObject4 struct for InlineObject4
type InlineObject4 struct {
	// Can be the name of either a View, Object Group, or index.  View and Object Groups all resolve to an index (i.e. they are built on indexes), this API accepts names of the resolved indexes or the by names of resources that can be used to resolve them. 
	Bucket *string `json:"bucket,omitempty"`
}

// NewInlineObject4 instantiates a new InlineObject4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject4() *InlineObject4 {
	this := InlineObject4{}
	return &this
}

// NewInlineObject4WithDefaults instantiates a new InlineObject4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject4WithDefaults() *InlineObject4 {
	this := InlineObject4{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *InlineObject4) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject4) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *InlineObject4) HasBucket() bool {
	if o != nil && o.Bucket != nil {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *InlineObject4) SetBucket(v string) {
	o.Bucket = &v
}

func (o InlineObject4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject4 struct {
	value *InlineObject4
	isSet bool
}

func (v NullableInlineObject4) Get() *InlineObject4 {
	return v.value
}

func (v *NullableInlineObject4) Set(val *InlineObject4) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject4) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject4(val *InlineObject4) *NullableInlineObject4 {
	return &NullableInlineObject4{value: val, isSet: true}
}

func (v NullableInlineObject4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


