# ObjectGroupFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Filter predicates to be applied to an Object key (requires either prefix or regex property)  | 
**Regex** | Pointer to **string** | A regular expression predicate to be applied to a key  | [optional] 
**Prefix** | Pointer to **string** | A regular expression predicate to be applied to the prefix of a key  | [optional] 

## Methods

### NewObjectGroupFilter

`func NewObjectGroupFilter(field string, ) *ObjectGroupFilter`

NewObjectGroupFilter instantiates a new ObjectGroupFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectGroupFilterWithDefaults

`func NewObjectGroupFilterWithDefaults() *ObjectGroupFilter`

NewObjectGroupFilterWithDefaults instantiates a new ObjectGroupFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ObjectGroupFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ObjectGroupFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ObjectGroupFilter) SetField(v string)`

SetField sets Field field to given value.


### GetRegex

`func (o *ObjectGroupFilter) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *ObjectGroupFilter) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *ObjectGroupFilter) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *ObjectGroupFilter) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetPrefix

`func (o *ObjectGroupFilter) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ObjectGroupFilter) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ObjectGroupFilter) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ObjectGroupFilter) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


