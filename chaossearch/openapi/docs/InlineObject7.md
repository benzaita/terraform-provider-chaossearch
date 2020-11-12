# InlineObject7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserInfoBlock** | Pointer to [**UserCreateSubAccountUserInfoBlock**](_user_createSubAccount_UserInfoBlock.md) |  | [optional] 
**GroupsIds** | Pointer to **[]string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Hocon** | Pointer to **[]string** | used to override or set internal fields of the sub-account (please work with ChaosSearch support to use)  | [optional] 

## Methods

### NewInlineObject7

`func NewInlineObject7() *InlineObject7`

NewInlineObject7 instantiates a new InlineObject7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject7WithDefaults

`func NewInlineObject7WithDefaults() *InlineObject7`

NewInlineObject7WithDefaults instantiates a new InlineObject7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserInfoBlock

`func (o *InlineObject7) GetUserInfoBlock() UserCreateSubAccountUserInfoBlock`

GetUserInfoBlock returns the UserInfoBlock field if non-nil, zero value otherwise.

### GetUserInfoBlockOk

`func (o *InlineObject7) GetUserInfoBlockOk() (*UserCreateSubAccountUserInfoBlock, bool)`

GetUserInfoBlockOk returns a tuple with the UserInfoBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoBlock

`func (o *InlineObject7) SetUserInfoBlock(v UserCreateSubAccountUserInfoBlock)`

SetUserInfoBlock sets UserInfoBlock field to given value.

### HasUserInfoBlock

`func (o *InlineObject7) HasUserInfoBlock() bool`

HasUserInfoBlock returns a boolean if a field has been set.

### GetGroupsIds

`func (o *InlineObject7) GetGroupsIds() []string`

GetGroupsIds returns the GroupsIds field if non-nil, zero value otherwise.

### GetGroupsIdsOk

`func (o *InlineObject7) GetGroupsIdsOk() (*[]string, bool)`

GetGroupsIdsOk returns a tuple with the GroupsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsIds

`func (o *InlineObject7) SetGroupsIds(v []string)`

SetGroupsIds sets GroupsIds field to given value.

### HasGroupsIds

`func (o *InlineObject7) HasGroupsIds() bool`

HasGroupsIds returns a boolean if a field has been set.

### GetPassword

`func (o *InlineObject7) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject7) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject7) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineObject7) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHocon

`func (o *InlineObject7) GetHocon() []string`

GetHocon returns the Hocon field if non-nil, zero value otherwise.

### GetHoconOk

`func (o *InlineObject7) GetHoconOk() (*[]string, bool)`

GetHoconOk returns a tuple with the Hocon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHocon

`func (o *InlineObject7) SetHocon(v []string)`

SetHocon sets Hocon field to given value.

### HasHocon

`func (o *InlineObject7) HasHocon() bool`

HasHocon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


