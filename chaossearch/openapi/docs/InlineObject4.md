# InlineObject4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Can be the name of either a View, Object Group, or index.  View and Object Groups all resolve to an index (i.e. they are built on indexes), this API accepts names of the resolved indexes or the by names of resources that can be used to resolve them.  | [optional] 

## Methods

### NewInlineObject4

`func NewInlineObject4() *InlineObject4`

NewInlineObject4 instantiates a new InlineObject4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject4WithDefaults

`func NewInlineObject4WithDefaults() *InlineObject4`

NewInlineObject4WithDefaults instantiates a new InlineObject4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *InlineObject4) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *InlineObject4) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *InlineObject4) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *InlineObject4) HasBucket() bool`

HasBucket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


