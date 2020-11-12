# InlineResponse201

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** | The identifier of the sub-account  | [optional] 
**Exists** | Pointer to **bool** | true if the sub-account exists after call, false otherwise  | [optional] 

## Methods

### NewInlineResponse201

`func NewInlineResponse201() *InlineResponse201`

NewInlineResponse201 instantiates a new InlineResponse201 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse201WithDefaults

`func NewInlineResponse201WithDefaults() *InlineResponse201`

NewInlineResponse201WithDefaults instantiates a new InlineResponse201 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *InlineResponse201) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *InlineResponse201) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *InlineResponse201) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *InlineResponse201) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetExists

`func (o *InlineResponse201) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *InlineResponse201) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *InlineResponse201) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *InlineResponse201) HasExists() bool`

HasExists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


