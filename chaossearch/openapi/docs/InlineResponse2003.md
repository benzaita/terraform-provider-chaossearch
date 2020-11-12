# InlineResponse2003

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]UserGroupsPermissions**](_user_groups_permissions.md) |  | [optional] 

## Methods

### NewInlineResponse2003

`func NewInlineResponse2003() *InlineResponse2003`

NewInlineResponse2003 instantiates a new InlineResponse2003 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2003WithDefaults

`func NewInlineResponse2003WithDefaults() *InlineResponse2003`

NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse2003) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse2003) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse2003) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse2003) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse2003) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2003) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2003) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse2003) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *InlineResponse2003) GetPermissions() []UserGroupsPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InlineResponse2003) GetPermissionsOk() (*[]UserGroupsPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InlineResponse2003) SetPermissions(v []UserGroupsPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InlineResponse2003) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


