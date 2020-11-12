# InlineResponse2001Users

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Hocon** | Pointer to **string** |  | [optional] 
**Activated** | Pointer to **bool** |  | [optional] 
**Regions** | Pointer to [**[]InlineResponse2001Regions**](inline_response_200_1_Regions.md) |  | [optional] 
**Deployed** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]InlineResponse2001Groups**](inline_response_200_1_Groups.md) |  | [optional] 
**Readonly** | Pointer to **bool** |  | [optional] 
**SubAccounts** | Pointer to [**[]InlineResponse2001SubAccounts**](inline_response_200_1_SubAccounts.md) |  | [optional] 
**ServiceType** | Pointer to **bool** |  | [optional] 
**FullName** | Pointer to **bool** |  | [optional] 

## Methods

### NewInlineResponse2001Users

`func NewInlineResponse2001Users() *InlineResponse2001Users`

NewInlineResponse2001Users instantiates a new InlineResponse2001Users object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001UsersWithDefaults

`func NewInlineResponse2001UsersWithDefaults() *InlineResponse2001Users`

NewInlineResponse2001UsersWithDefaults instantiates a new InlineResponse2001Users object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *InlineResponse2001Users) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *InlineResponse2001Users) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *InlineResponse2001Users) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *InlineResponse2001Users) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsername

`func (o *InlineResponse2001Users) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InlineResponse2001Users) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InlineResponse2001Users) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InlineResponse2001Users) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetHocon

`func (o *InlineResponse2001Users) GetHocon() string`

GetHocon returns the Hocon field if non-nil, zero value otherwise.

### GetHoconOk

`func (o *InlineResponse2001Users) GetHoconOk() (*string, bool)`

GetHoconOk returns a tuple with the Hocon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHocon

`func (o *InlineResponse2001Users) SetHocon(v string)`

SetHocon sets Hocon field to given value.

### HasHocon

`func (o *InlineResponse2001Users) HasHocon() bool`

HasHocon returns a boolean if a field has been set.

### GetActivated

`func (o *InlineResponse2001Users) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *InlineResponse2001Users) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *InlineResponse2001Users) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *InlineResponse2001Users) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetRegions

`func (o *InlineResponse2001Users) GetRegions() []InlineResponse2001Regions`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *InlineResponse2001Users) GetRegionsOk() (*[]InlineResponse2001Regions, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *InlineResponse2001Users) SetRegions(v []InlineResponse2001Regions)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *InlineResponse2001Users) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetDeployed

`func (o *InlineResponse2001Users) GetDeployed() bool`

GetDeployed returns the Deployed field if non-nil, zero value otherwise.

### GetDeployedOk

`func (o *InlineResponse2001Users) GetDeployedOk() (*bool, bool)`

GetDeployedOk returns a tuple with the Deployed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployed

`func (o *InlineResponse2001Users) SetDeployed(v bool)`

SetDeployed sets Deployed field to given value.

### HasDeployed

`func (o *InlineResponse2001Users) HasDeployed() bool`

HasDeployed returns a boolean if a field has been set.

### GetEmail

`func (o *InlineResponse2001Users) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse2001Users) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse2001Users) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineResponse2001Users) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGroups

`func (o *InlineResponse2001Users) GetGroups() []InlineResponse2001Groups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *InlineResponse2001Users) GetGroupsOk() (*[]InlineResponse2001Groups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *InlineResponse2001Users) SetGroups(v []InlineResponse2001Groups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *InlineResponse2001Users) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetReadonly

`func (o *InlineResponse2001Users) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *InlineResponse2001Users) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *InlineResponse2001Users) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *InlineResponse2001Users) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.

### GetSubAccounts

`func (o *InlineResponse2001Users) GetSubAccounts() []InlineResponse2001SubAccounts`

GetSubAccounts returns the SubAccounts field if non-nil, zero value otherwise.

### GetSubAccountsOk

`func (o *InlineResponse2001Users) GetSubAccountsOk() (*[]InlineResponse2001SubAccounts, bool)`

GetSubAccountsOk returns a tuple with the SubAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccounts

`func (o *InlineResponse2001Users) SetSubAccounts(v []InlineResponse2001SubAccounts)`

SetSubAccounts sets SubAccounts field to given value.

### HasSubAccounts

`func (o *InlineResponse2001Users) HasSubAccounts() bool`

HasSubAccounts returns a boolean if a field has been set.

### GetServiceType

`func (o *InlineResponse2001Users) GetServiceType() bool`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *InlineResponse2001Users) GetServiceTypeOk() (*bool, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *InlineResponse2001Users) SetServiceType(v bool)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *InlineResponse2001Users) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetFullName

`func (o *InlineResponse2001Users) GetFullName() bool`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *InlineResponse2001Users) GetFullNameOk() (*bool, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *InlineResponse2001Users) SetFullName(v bool)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *InlineResponse2001Users) HasFullName() bool`

HasFullName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


