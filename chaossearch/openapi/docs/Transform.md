# Transform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A field transformer type  ## MaterializeRegexTransform \&quot;Transforms\&quot;: [     {       \&quot;_type\&quot;:\&quot;MaterializeRegexTransform\&quot;,       \&quot;pattern\&quot;:\&quot;.*(c).*(e)\&quot;,       \&quot;inputField\&quot;:\&quot;food\&quot;,       \&quot;outputFields\&quot;: [         {\&quot;name\&quot;:\&quot;foo-number\&quot;,\&quot;type\&quot;:\&quot;NUMBER\&quot;},{\&quot;name\&quot;:\&quot;foo-time\&quot;,\&quot;type\&quot;:\&quot;TIMEVAL\&quot;}       ]     },     {       \&quot;_type\&quot;:\&quot;MaterializeRegexTransform\&quot;,       \&quot;pattern\&quot;:\&quot;.*(c).*(e)\&quot;,       \&quot;inputField\&quot;:\&quot;first_name\&quot;,       \&quot;outputFields\&quot;:[         {\&quot;name\&quot;:\&quot;foo-first\&quot;,\&quot;type\&quot;:\&quot;STRING\&quot;},{\&quot;name\&quot;:\&quot;foo-second\&quot;,\&quot;type\&quot;:\&quot;STRING\&quot;}       ]     }   ] }  The above transforms when given the following column values will produce new columns    column value \&quot;regex-with-capture\&quot; -&gt; virtual column \&quot;foo-number\&quot;: \&quot;\&quot;, virtual column \&quot;foo-time\&quot;: \&quot;\&quot;   column value \&quot;regex-with-capture\&quot; -&gt; virtual column \&quot;foo-first\&quot;: \&quot;c\&quot;, virtual column \&quot;foo-second\&quot;: \&quot;e\&quot;    The first row produces empty column values because the type(s) do not parse  | 
**Pattern** | **string** | A regular expression with capture groups to extract values to apply transformation to | 
**InputField** | **string** | The column name from the source index | 
**OutputFields** | [**[]OutputField**](OutputField.md) |  | 

## Methods

### NewTransform

`func NewTransform(type_ string, pattern string, inputField string, outputFields []OutputField, ) *Transform`

NewTransform instantiates a new Transform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformWithDefaults

`func NewTransformWithDefaults() *Transform`

NewTransformWithDefaults instantiates a new Transform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Transform) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transform) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transform) SetType(v string)`

SetType sets Type field to given value.


### GetPattern

`func (o *Transform) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *Transform) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *Transform) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetInputField

`func (o *Transform) GetInputField() string`

GetInputField returns the InputField field if non-nil, zero value otherwise.

### GetInputFieldOk

`func (o *Transform) GetInputFieldOk() (*string, bool)`

GetInputFieldOk returns a tuple with the InputField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputField

`func (o *Transform) SetInputField(v string)`

SetInputField sets InputField field to given value.


### GetOutputFields

`func (o *Transform) GetOutputFields() []OutputField`

GetOutputFields returns the OutputFields field if non-nil, zero value otherwise.

### GetOutputFieldsOk

`func (o *Transform) GetOutputFieldsOk() (*[]OutputField, bool)`

GetOutputFieldsOk returns a tuple with the OutputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFields

`func (o *Transform) SetOutputFields(v []OutputField)`

SetOutputFields sets OutputFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


