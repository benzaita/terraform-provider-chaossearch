# BucketCreateObjectGroupOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreIrregular** | **bool** | Ignore Irregular Objects while indexing a bucket (deprecated)  This setting should always be set &#x60;true&#x60;.  | 
**Compression** | Pointer to **string** | Compression algorithm used to compress the Object stored in S3, used by ChaosSearch to decompress the archive for indexing.  **NOTE:** Do not supply this key if the Objects in the bucket are not compressed  | [optional] 
**ColTypes** | Pointer to **map[string]interface{}** | Column Types can be used to override a given column&#39;s data type. Parsing of Objects will produce values of the overridden type  Example      \&quot;colTypes\&quot;: { \&quot;foo\&quot;: \&quot;NUMBER\&quot; }  The above example will override the data type of the column named \&quot;foo\&quot; so its value will be a NUMBER  | [optional] 

## Methods

### NewBucketCreateObjectGroupOptions

`func NewBucketCreateObjectGroupOptions(ignoreIrregular bool, ) *BucketCreateObjectGroupOptions`

NewBucketCreateObjectGroupOptions instantiates a new BucketCreateObjectGroupOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketCreateObjectGroupOptionsWithDefaults

`func NewBucketCreateObjectGroupOptionsWithDefaults() *BucketCreateObjectGroupOptions`

NewBucketCreateObjectGroupOptionsWithDefaults instantiates a new BucketCreateObjectGroupOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreIrregular

`func (o *BucketCreateObjectGroupOptions) GetIgnoreIrregular() bool`

GetIgnoreIrregular returns the IgnoreIrregular field if non-nil, zero value otherwise.

### GetIgnoreIrregularOk

`func (o *BucketCreateObjectGroupOptions) GetIgnoreIrregularOk() (*bool, bool)`

GetIgnoreIrregularOk returns a tuple with the IgnoreIrregular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIrregular

`func (o *BucketCreateObjectGroupOptions) SetIgnoreIrregular(v bool)`

SetIgnoreIrregular sets IgnoreIrregular field to given value.


### GetCompression

`func (o *BucketCreateObjectGroupOptions) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *BucketCreateObjectGroupOptions) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *BucketCreateObjectGroupOptions) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *BucketCreateObjectGroupOptions) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetColTypes

`func (o *BucketCreateObjectGroupOptions) GetColTypes() map[string]interface{}`

GetColTypes returns the ColTypes field if non-nil, zero value otherwise.

### GetColTypesOk

`func (o *BucketCreateObjectGroupOptions) GetColTypesOk() (*map[string]interface{}, bool)`

GetColTypesOk returns a tuple with the ColTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColTypes

`func (o *BucketCreateObjectGroupOptions) SetColTypes(v map[string]interface{})`

SetColTypes sets ColTypes field to given value.

### HasColTypes

`func (o *BucketCreateObjectGroupOptions) HasColTypes() bool`

HasColTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


