# InlineObject5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | Object Group name  | [optional] 
**ModelMode** | Pointer to **int32** | Integer to represent the state(s) of the model  1 &#x3D;&gt; Restart Indexing  0 &#x3D;&gt; Start Indexing  -1 &#x3D;&gt; Stop Indexing  -2 &#x3D;&gt; Terminate Indexing  | [optional] 

## Methods

### NewInlineObject5

`func NewInlineObject5() *InlineObject5`

NewInlineObject5 instantiates a new InlineObject5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject5WithDefaults

`func NewInlineObject5WithDefaults() *InlineObject5`

NewInlineObject5WithDefaults instantiates a new InlineObject5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *InlineObject5) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *InlineObject5) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *InlineObject5) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *InlineObject5) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetModelMode

`func (o *InlineObject5) GetModelMode() int32`

GetModelMode returns the ModelMode field if non-nil, zero value otherwise.

### GetModelModeOk

`func (o *InlineObject5) GetModelModeOk() (*int32, bool)`

GetModelModeOk returns a tuple with the ModelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelMode

`func (o *InlineObject5) SetModelMode(v int32)`

SetModelMode sets ModelMode field to given value.

### HasModelMode

`func (o *InlineObject5) HasModelMode() bool`

HasModelMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


