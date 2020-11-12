# InlineObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | Object Group Name | 
**IndexRetention** | Pointer to **int32** | A time window (in terms of days) to determine whether a resource should be included (a value of -1 means no limit) | [optional] [default to -1]
**TargetActiveIndex** | Pointer to **int32** | Allow overriding number of \&quot;irrevokable\&quot; leases on compute for tasks associated with in. Use with care.  Object Group indexing model uses a lease system to allocate workers (registered units) for indexing jobs. Each worker is allowed one lease. By default an index for a given Object Group will require 1 lease in order to make progress. By setting the target active index value to something larger than 1, the system will attempt to allocate more leases for indexing. If the system has no more leases to give, the index will progress as long as it has at least 1 lease allocated.  By default the value is 1  ChaosSearch support should be enagaged to help tune this value for latency sensitive indexes.  | [optional] [default to 1]

## Methods

### NewInlineObject1

`func NewInlineObject1(bucket string, ) *InlineObject1`

NewInlineObject1 instantiates a new InlineObject1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject1WithDefaults

`func NewInlineObject1WithDefaults() *InlineObject1`

NewInlineObject1WithDefaults instantiates a new InlineObject1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *InlineObject1) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *InlineObject1) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *InlineObject1) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetIndexRetention

`func (o *InlineObject1) GetIndexRetention() int32`

GetIndexRetention returns the IndexRetention field if non-nil, zero value otherwise.

### GetIndexRetentionOk

`func (o *InlineObject1) GetIndexRetentionOk() (*int32, bool)`

GetIndexRetentionOk returns a tuple with the IndexRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexRetention

`func (o *InlineObject1) SetIndexRetention(v int32)`

SetIndexRetention sets IndexRetention field to given value.

### HasIndexRetention

`func (o *InlineObject1) HasIndexRetention() bool`

HasIndexRetention returns a boolean if a field has been set.

### GetTargetActiveIndex

`func (o *InlineObject1) GetTargetActiveIndex() int32`

GetTargetActiveIndex returns the TargetActiveIndex field if non-nil, zero value otherwise.

### GetTargetActiveIndexOk

`func (o *InlineObject1) GetTargetActiveIndexOk() (*int32, bool)`

GetTargetActiveIndexOk returns a tuple with the TargetActiveIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetActiveIndex

`func (o *InlineObject1) SetTargetActiveIndex(v int32)`

SetTargetActiveIndex sets TargetActiveIndex field to given value.

### HasTargetActiveIndex

`func (o *InlineObject1) HasTargetActiveIndex() bool`

HasTargetActiveIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


