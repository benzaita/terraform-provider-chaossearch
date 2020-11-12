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

// Transform struct for Transform
type Transform struct {
	// A field transformer type  ## MaterializeRegexTransform \"Transforms\": [     {       \"_type\":\"MaterializeRegexTransform\",       \"pattern\":\".*(c).*(e)\",       \"inputField\":\"food\",       \"outputFields\": [         {\"name\":\"foo-number\",\"type\":\"NUMBER\"},{\"name\":\"foo-time\",\"type\":\"TIMEVAL\"}       ]     },     {       \"_type\":\"MaterializeRegexTransform\",       \"pattern\":\".*(c).*(e)\",       \"inputField\":\"first_name\",       \"outputFields\":[         {\"name\":\"foo-first\",\"type\":\"STRING\"},{\"name\":\"foo-second\",\"type\":\"STRING\"}       ]     }   ] }  The above transforms when given the following column values will produce new columns    column value \"regex-with-capture\" -> virtual column \"foo-number\": \"\", virtual column \"foo-time\": \"\"   column value \"regex-with-capture\" -> virtual column \"foo-first\": \"c\", virtual column \"foo-second\": \"e\"    The first row produces empty column values because the type(s) do not parse 
	Type string `json:"_type"`
	// A regular expression with capture groups to extract values to apply transformation to
	Pattern string `json:"pattern"`
	// The column name from the source index
	InputField string `json:"inputField"`
	OutputFields []OutputField `json:"outputFields"`
}

// NewTransform instantiates a new Transform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransform(type_ string, pattern string, inputField string, outputFields []OutputField, ) *Transform {
	this := Transform{}
	this.Type = type_
	this.Pattern = pattern
	this.InputField = inputField
	this.OutputFields = outputFields
	return &this
}

// NewTransformWithDefaults instantiates a new Transform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformWithDefaults() *Transform {
	this := Transform{}
	return &this
}

// GetType returns the Type field value
func (o *Transform) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Transform) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Transform) SetType(v string) {
	o.Type = v
}

// GetPattern returns the Pattern field value
func (o *Transform) GetPattern() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *Transform) GetPatternOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *Transform) SetPattern(v string) {
	o.Pattern = v
}

// GetInputField returns the InputField field value
func (o *Transform) GetInputField() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.InputField
}

// GetInputFieldOk returns a tuple with the InputField field value
// and a boolean to check if the value has been set.
func (o *Transform) GetInputFieldOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InputField, true
}

// SetInputField sets field value
func (o *Transform) SetInputField(v string) {
	o.InputField = v
}

// GetOutputFields returns the OutputFields field value
func (o *Transform) GetOutputFields() []OutputField {
	if o == nil  {
		var ret []OutputField
		return ret
	}

	return o.OutputFields
}

// GetOutputFieldsOk returns a tuple with the OutputFields field value
// and a boolean to check if the value has been set.
func (o *Transform) GetOutputFieldsOk() (*[]OutputField, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OutputFields, true
}

// SetOutputFields sets field value
func (o *Transform) SetOutputFields(v []OutputField) {
	o.OutputFields = v
}

func (o Transform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_type"] = o.Type
	}
	if true {
		toSerialize["pattern"] = o.Pattern
	}
	if true {
		toSerialize["inputField"] = o.InputField
	}
	if true {
		toSerialize["outputFields"] = o.OutputFields
	}
	return json.Marshal(toSerialize)
}

type NullableTransform struct {
	value *Transform
	isSet bool
}

func (v NullableTransform) Get() *Transform {
	return v.value
}

func (v *NullableTransform) Set(val *Transform) {
	v.value = val
	v.isSet = true
}

func (v NullableTransform) IsSet() bool {
	return v.isSet
}

func (v *NullableTransform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransform(val *Transform) *NullableTransform {
	return &NullableTransform{value: val, isSet: true}
}

func (v NullableTransform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


