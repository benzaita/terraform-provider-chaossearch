# InlineObject3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | View Name | [optional] 
**IndexRetention** | Pointer to [**ViewTimeWindow**](ViewTimeWindow.md) |  | [optional] 

## Methods

### NewInlineObject3

`func NewInlineObject3() *InlineObject3`

NewInlineObject3 instantiates a new InlineObject3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject3WithDefaults

`func NewInlineObject3WithDefaults() *InlineObject3`

NewInlineObject3WithDefaults instantiates a new InlineObject3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *InlineObject3) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *InlineObject3) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *InlineObject3) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *InlineObject3) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetIndexRetention

`func (o *InlineObject3) GetIndexRetention() ViewTimeWindow`

GetIndexRetention returns the IndexRetention field if non-nil, zero value otherwise.

### GetIndexRetentionOk

`func (o *InlineObject3) GetIndexRetentionOk() (*ViewTimeWindow, bool)`

GetIndexRetentionOk returns a tuple with the IndexRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexRetention

`func (o *InlineObject3) SetIndexRetention(v ViewTimeWindow)`

SetIndexRetention sets IndexRetention field to given value.

### HasIndexRetention

`func (o *InlineObject3) HasIndexRetention() bool`

HasIndexRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


