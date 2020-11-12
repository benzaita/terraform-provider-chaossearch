# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]UserGroupsPermissions1**](_user_groups_permissions_1.md) |  | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *InlineObject) GetPermissions() []UserGroupsPermissions1`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InlineObject) GetPermissionsOk() (*[]UserGroupsPermissions1, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InlineObject) SetPermissions(v []UserGroupsPermissions1)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InlineObject) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


