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

// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	// View Name
	Bucket string `json:"bucket"`
	// The View's source of data, a list of Object Group Names.  IMPORTANT: An empty array means ALL Object Groups!  **NOTE:** View Name cannot be identical with an Object Group Name. 
	Sources []string `json:"sources"`
	// An Index Pattern (regex) that matches and filters source based on index names, practically speaking from Object Groups Names.  Example:      Object Group Name: foo     Resulting index names (for daily indexes):     _foo_2020-10-15_     _foo_2020-10-16_      Resulting index names:     _foo_ 
	IndexPattern string `json:"indexPattern"`
	// Flag to toggle case-sensitivity of searches of all values in the View  Example:      row: {\"Foo\": \"Bar\"}     Searching for \"Bar\" is treated case-insensitive, but keys such as \"Foo\" are still case-sensitive. 
	CaseInsensitive *bool `json:"caseInsensitive,omitempty"`
	IndexRetention *ViewTimeWindow `json:"indexRetention,omitempty"`
	// transformations on fields used to produce new fields from existing ones
	Transforms *[]Transform `json:"transforms,omitempty"`
	// TODO describe what and how filters work, especially predicates  Also, how to add AND, and OR logic here  Example log:      {       \"system\": \"foo\",       \"type\": \"elb-access\",       \"stage\": \"staging\",     }      User allowed to see all data where \"system\" = \"foo\", and \"type\" matches \".*(-access$)\" but not     \"stage\" = \"production\".  Example:      \"filter\": {       \"predicate\": {           \"field\": \"attrs.Cluster\",           \"query\": \"au-production\",           \"state\": {               \"_type\": \"chaossumo.query.QEP.Predicate.TextMatchState.Exact\"           },           \"_type\": \"chaossumo.query.NIRFrontend.Request.Predicate.TextMatch\"       }     } 
	Filters interface{} `json:"filters,omitempty"`
}

// NewInlineObject2 instantiates a new InlineObject2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2(bucket string, sources []string, indexPattern string, ) *InlineObject2 {
	this := InlineObject2{}
	this.Bucket = bucket
	this.Sources = sources
	this.IndexPattern = indexPattern
	var caseInsensitive bool = false
	this.CaseInsensitive = &caseInsensitive
	return &this
}

// NewInlineObject2WithDefaults instantiates a new InlineObject2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2WithDefaults() *InlineObject2 {
	this := InlineObject2{}
	var caseInsensitive bool = false
	this.CaseInsensitive = &caseInsensitive
	return &this
}

// GetBucket returns the Bucket field value
func (o *InlineObject2) GetBucket() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetBucketOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *InlineObject2) SetBucket(v string) {
	o.Bucket = v
}

// GetSources returns the Sources field value
func (o *InlineObject2) GetSources() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetSourcesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value
func (o *InlineObject2) SetSources(v []string) {
	o.Sources = v
}

// GetIndexPattern returns the IndexPattern field value
func (o *InlineObject2) GetIndexPattern() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.IndexPattern
}

// GetIndexPatternOk returns a tuple with the IndexPattern field value
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetIndexPatternOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IndexPattern, true
}

// SetIndexPattern sets field value
func (o *InlineObject2) SetIndexPattern(v string) {
	o.IndexPattern = v
}

// GetCaseInsensitive returns the CaseInsensitive field value if set, zero value otherwise.
func (o *InlineObject2) GetCaseInsensitive() bool {
	if o == nil || o.CaseInsensitive == nil {
		var ret bool
		return ret
	}
	return *o.CaseInsensitive
}

// GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetCaseInsensitiveOk() (*bool, bool) {
	if o == nil || o.CaseInsensitive == nil {
		return nil, false
	}
	return o.CaseInsensitive, true
}

// HasCaseInsensitive returns a boolean if a field has been set.
func (o *InlineObject2) HasCaseInsensitive() bool {
	if o != nil && o.CaseInsensitive != nil {
		return true
	}

	return false
}

// SetCaseInsensitive gets a reference to the given bool and assigns it to the CaseInsensitive field.
func (o *InlineObject2) SetCaseInsensitive(v bool) {
	o.CaseInsensitive = &v
}

// GetIndexRetention returns the IndexRetention field value if set, zero value otherwise.
func (o *InlineObject2) GetIndexRetention() ViewTimeWindow {
	if o == nil || o.IndexRetention == nil {
		var ret ViewTimeWindow
		return ret
	}
	return *o.IndexRetention
}

// GetIndexRetentionOk returns a tuple with the IndexRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetIndexRetentionOk() (*ViewTimeWindow, bool) {
	if o == nil || o.IndexRetention == nil {
		return nil, false
	}
	return o.IndexRetention, true
}

// HasIndexRetention returns a boolean if a field has been set.
func (o *InlineObject2) HasIndexRetention() bool {
	if o != nil && o.IndexRetention != nil {
		return true
	}

	return false
}

// SetIndexRetention gets a reference to the given ViewTimeWindow and assigns it to the IndexRetention field.
func (o *InlineObject2) SetIndexRetention(v ViewTimeWindow) {
	o.IndexRetention = &v
}

// GetTransforms returns the Transforms field value if set, zero value otherwise.
func (o *InlineObject2) GetTransforms() []Transform {
	if o == nil || o.Transforms == nil {
		var ret []Transform
		return ret
	}
	return *o.Transforms
}

// GetTransformsOk returns a tuple with the Transforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetTransformsOk() (*[]Transform, bool) {
	if o == nil || o.Transforms == nil {
		return nil, false
	}
	return o.Transforms, true
}

// HasTransforms returns a boolean if a field has been set.
func (o *InlineObject2) HasTransforms() bool {
	if o != nil && o.Transforms != nil {
		return true
	}

	return false
}

// SetTransforms gets a reference to the given []Transform and assigns it to the Transforms field.
func (o *InlineObject2) SetTransforms(v []Transform) {
	o.Transforms = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject2) GetFilters() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject2) GetFiltersOk() (*interface{}, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *InlineObject2) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given interface{} and assigns it to the Filters field.
func (o *InlineObject2) SetFilters(v interface{}) {
	o.Filters = v
}

func (o InlineObject2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bucket"] = o.Bucket
	}
	if true {
		toSerialize["sources"] = o.Sources
	}
	if true {
		toSerialize["indexPattern"] = o.IndexPattern
	}
	if o.CaseInsensitive != nil {
		toSerialize["caseInsensitive"] = o.CaseInsensitive
	}
	if o.IndexRetention != nil {
		toSerialize["indexRetention"] = o.IndexRetention
	}
	if o.Transforms != nil {
		toSerialize["transforms"] = o.Transforms
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject2 struct {
	value *InlineObject2
	isSet bool
}

func (v NullableInlineObject2) Get() *InlineObject2 {
	return v.value
}

func (v *NullableInlineObject2) Set(val *InlineObject2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2(val *InlineObject2) *NullableInlineObject2 {
	return &NullableInlineObject2{value: val, isSet: true}
}

func (v NullableInlineObject2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


