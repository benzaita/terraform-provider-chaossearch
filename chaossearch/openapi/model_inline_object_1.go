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

// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	// Object Group Name
	Bucket string `json:"bucket"`
	// A time window (in terms of days) to determine whether a resource should be included (a value of -1 means no limit)
	IndexRetention *int32 `json:"indexRetention,omitempty"`
	// Allow overriding number of \"irrevokable\" leases on compute for tasks associated with in. Use with care.  Object Group indexing model uses a lease system to allocate workers (registered units) for indexing jobs. Each worker is allowed one lease. By default an index for a given Object Group will require 1 lease in order to make progress. By setting the target active index value to something larger than 1, the system will attempt to allocate more leases for indexing. If the system has no more leases to give, the index will progress as long as it has at least 1 lease allocated.  By default the value is 1  ChaosSearch support should be enagaged to help tune this value for latency sensitive indexes. 
	TargetActiveIndex *int32 `json:"targetActiveIndex,omitempty"`
}

// NewInlineObject1 instantiates a new InlineObject1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1(bucket string, ) *InlineObject1 {
	this := InlineObject1{}
	this.Bucket = bucket
	var indexRetention int32 = -1
	this.IndexRetention = &indexRetention
	var targetActiveIndex int32 = 1
	this.TargetActiveIndex = &targetActiveIndex
	return &this
}

// NewInlineObject1WithDefaults instantiates a new InlineObject1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1WithDefaults() *InlineObject1 {
	this := InlineObject1{}
	var indexRetention int32 = -1
	this.IndexRetention = &indexRetention
	var targetActiveIndex int32 = 1
	this.TargetActiveIndex = &targetActiveIndex
	return &this
}

// GetBucket returns the Bucket field value
func (o *InlineObject1) GetBucket() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetBucketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *InlineObject1) SetBucket(v string) {
	o.Bucket = v
}

// GetIndexRetention returns the IndexRetention field value if set, zero value otherwise.
func (o *InlineObject1) GetIndexRetention() int32 {
	if o == nil || o.IndexRetention == nil {
		var ret int32
		return ret
	}
	return *o.IndexRetention
}

// GetIndexRetentionOk returns a tuple with the IndexRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetIndexRetentionOk() (*int32, bool) {
	if o == nil || o.IndexRetention == nil {
		return nil, false
	}
	return o.IndexRetention, true
}

// HasIndexRetention returns a boolean if a field has been set.
func (o *InlineObject1) HasIndexRetention() bool {
	if o != nil && o.IndexRetention != nil {
		return true
	}

	return false
}

// SetIndexRetention gets a reference to the given int32 and assigns it to the IndexRetention field.
func (o *InlineObject1) SetIndexRetention(v int32) {
	o.IndexRetention = &v
}

// GetTargetActiveIndex returns the TargetActiveIndex field value if set, zero value otherwise.
func (o *InlineObject1) GetTargetActiveIndex() int32 {
	if o == nil || o.TargetActiveIndex == nil {
		var ret int32
		return ret
	}
	return *o.TargetActiveIndex
}

// GetTargetActiveIndexOk returns a tuple with the TargetActiveIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetTargetActiveIndexOk() (*int32, bool) {
	if o == nil || o.TargetActiveIndex == nil {
		return nil, false
	}
	return o.TargetActiveIndex, true
}

// HasTargetActiveIndex returns a boolean if a field has been set.
func (o *InlineObject1) HasTargetActiveIndex() bool {
	if o != nil && o.TargetActiveIndex != nil {
		return true
	}

	return false
}

// SetTargetActiveIndex gets a reference to the given int32 and assigns it to the TargetActiveIndex field.
func (o *InlineObject1) SetTargetActiveIndex(v int32) {
	o.TargetActiveIndex = &v
}

func (o InlineObject1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bucket"] = o.Bucket
	}
	if o.IndexRetention != nil {
		toSerialize["indexRetention"] = o.IndexRetention
	}
	if o.TargetActiveIndex != nil {
		toSerialize["targetActiveIndex"] = o.TargetActiveIndex
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1 struct {
	value *InlineObject1
	isSet bool
}

func (v NullableInlineObject1) Get() *InlineObject1 {
	return v.value
}

func (v *NullableInlineObject1) Set(val *InlineObject1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1(val *InlineObject1) *NullableInlineObject1 {
	return &NullableInlineObject1{value: val, isSet: true}
}

func (v NullableInlineObject1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


