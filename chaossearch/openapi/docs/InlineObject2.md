# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | View Name | 
**Sources** | **[]string** | The View&#39;s source of data, a list of Object Group Names.  IMPORTANT: An empty array means ALL Object Groups!  **NOTE:** View Name cannot be identical with an Object Group Name.  | 
**IndexPattern** | **string** | An Index Pattern (regex) that matches and filters source based on index names, practically speaking from Object Groups Names.  Example:      Object Group Name: foo     Resulting index names (for daily indexes):     _foo_2020-10-15_     _foo_2020-10-16_      Resulting index names:     _foo_  | 
**CaseInsensitive** | Pointer to **bool** | Flag to toggle case-sensitivity of searches of all values in the View  Example:      row: {\&quot;Foo\&quot;: \&quot;Bar\&quot;}     Searching for \&quot;Bar\&quot; is treated case-insensitive, but keys such as \&quot;Foo\&quot; are still case-sensitive.  | [optional] [default to false]
**IndexRetention** | Pointer to [**ViewTimeWindow**](ViewTimeWindow.md) |  | [optional] 
**Transforms** | Pointer to [**[]Transform**](Transform.md) | transformations on fields used to produce new fields from existing ones | [optional] [default to []]
**Filters** | Pointer to **interface{}** | TODO describe what and how filters work, especially predicates  Also, how to add AND, and OR logic here  Example log:      {       \&quot;system\&quot;: \&quot;foo\&quot;,       \&quot;type\&quot;: \&quot;elb-access\&quot;,       \&quot;stage\&quot;: \&quot;staging\&quot;,     }      User allowed to see all data where \&quot;system\&quot; &#x3D; \&quot;foo\&quot;, and \&quot;type\&quot; matches \&quot;.*(-access$)\&quot; but not     \&quot;stage\&quot; &#x3D; \&quot;production\&quot;.  Example:      \&quot;filter\&quot;: {       \&quot;predicate\&quot;: {           \&quot;field\&quot;: \&quot;attrs.Cluster\&quot;,           \&quot;query\&quot;: \&quot;au-production\&quot;,           \&quot;state\&quot;: {               \&quot;_type\&quot;: \&quot;chaossumo.query.QEP.Predicate.TextMatchState.Exact\&quot;           },           \&quot;_type\&quot;: \&quot;chaossumo.query.NIRFrontend.Request.Predicate.TextMatch\&quot;       }     }  | [optional] 

## Methods

### NewInlineObject2

`func NewInlineObject2(bucket string, sources []string, indexPattern string, ) *InlineObject2`

NewInlineObject2 instantiates a new InlineObject2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject2WithDefaults

`func NewInlineObject2WithDefaults() *InlineObject2`

NewInlineObject2WithDefaults instantiates a new InlineObject2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *InlineObject2) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *InlineObject2) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *InlineObject2) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetSources

`func (o *InlineObject2) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *InlineObject2) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *InlineObject2) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetIndexPattern

`func (o *InlineObject2) GetIndexPattern() string`

GetIndexPattern returns the IndexPattern field if non-nil, zero value otherwise.

### GetIndexPatternOk

`func (o *InlineObject2) GetIndexPatternOk() (*string, bool)`

GetIndexPatternOk returns a tuple with the IndexPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPattern

`func (o *InlineObject2) SetIndexPattern(v string)`

SetIndexPattern sets IndexPattern field to given value.


### GetCaseInsensitive

`func (o *InlineObject2) GetCaseInsensitive() bool`

GetCaseInsensitive returns the CaseInsensitive field if non-nil, zero value otherwise.

### GetCaseInsensitiveOk

`func (o *InlineObject2) GetCaseInsensitiveOk() (*bool, bool)`

GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitive

`func (o *InlineObject2) SetCaseInsensitive(v bool)`

SetCaseInsensitive sets CaseInsensitive field to given value.

### HasCaseInsensitive

`func (o *InlineObject2) HasCaseInsensitive() bool`

HasCaseInsensitive returns a boolean if a field has been set.

### GetIndexRetention

`func (o *InlineObject2) GetIndexRetention() ViewTimeWindow`

GetIndexRetention returns the IndexRetention field if non-nil, zero value otherwise.

### GetIndexRetentionOk

`func (o *InlineObject2) GetIndexRetentionOk() (*ViewTimeWindow, bool)`

GetIndexRetentionOk returns a tuple with the IndexRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexRetention

`func (o *InlineObject2) SetIndexRetention(v ViewTimeWindow)`

SetIndexRetention sets IndexRetention field to given value.

### HasIndexRetention

`func (o *InlineObject2) HasIndexRetention() bool`

HasIndexRetention returns a boolean if a field has been set.

### GetTransforms

`func (o *InlineObject2) GetTransforms() []Transform`

GetTransforms returns the Transforms field if non-nil, zero value otherwise.

### GetTransformsOk

`func (o *InlineObject2) GetTransformsOk() (*[]Transform, bool)`

GetTransformsOk returns a tuple with the Transforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransforms

`func (o *InlineObject2) SetTransforms(v []Transform)`

SetTransforms sets Transforms field to given value.

### HasTransforms

`func (o *InlineObject2) HasTransforms() bool`

HasTransforms returns a boolean if a field has been set.

### GetFilters

`func (o *InlineObject2) GetFilters() interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *InlineObject2) GetFiltersOk() (*interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *InlineObject2) SetFilters(v interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *InlineObject2) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *InlineObject2) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *InlineObject2) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


