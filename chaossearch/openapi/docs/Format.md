# Format

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The supported format types for parsing data to be indexed. The format will be used against all Objects in an Object Group. For any Objects that cannot be completely parsed.  If data of the wrong type is ingested and attempted to be indexed, that data is dropped and an event is emmitted (displayed in the UI or available via API). If a field in the data is corrupt or missing it is not an error.  JSON example      \&quot;format\&quot;:{\&quot;_type\&quot;:\&quot;JSON\&quot;,\&quot;stripPrefix\&quot;:true,\&quot;horizontal\&quot;:true}  CSV format      \&quot;format\&quot;:{\&quot;_type\&quot;:\&quot;CSV\&quot;,\&quot;columnDelimiter\&quot;:\&quot;,\&quot;,\&quot;rowDelimiter\&quot;:\&quot;\\n\&quot;,\&quot;headerRow\&quot;:true}  LOG example      \&quot;format\&quot;:{\&quot;_type\&quot;:\&quot;LOG\&quot;,\&quot;pattern\&quot;:\&quot;.*\&quot;}  | 

## Methods

### NewFormat

`func NewFormat(type_ string, ) *Format`

NewFormat instantiates a new Format object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormatWithDefaults

`func NewFormatWithDefaults() *Format`

NewFormatWithDefaults instantiates a new Format object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Format) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Format) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Format) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


