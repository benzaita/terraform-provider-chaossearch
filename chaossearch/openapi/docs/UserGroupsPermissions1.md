# UserGroupsPermissions1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | The type of effect for the permission of this container, Allow or Deny  | 
**Action** | **string** | The action(s) to which this container applies  | Action | Definition| |--------|-----------| | * | All-access | | s3:* | Ability to access S3 if given permissions in the IAM Policy and API calls to the Chaos Index | | s3:aws:* | Ability to access S3 if given permissions in the IAM Policy | | s3:chaos:* | Ability to access the ChaosSearch Admin API | | elastic:* | Ability to access the Elasticsearch API | | elastic:opendistro:* | Ability to access the Elasticsearch API | | chaos:* | Ability to access all replica, query, and theme settings | | chaos:replica:* | Ability to access all replica information (i.e. compute allocation), initiate burst and see the compute status | | chaos:replica:burst | Ability to click burst | | chaos:replica:status | Ability to see how many compute resources are allocated | | chaos:query:* | Full access to query permissions | | chaos:query:status | Ability to access the Query progress bar | | chaos:query:migrate |  | | chaos:query:cancel | Ability to Cancel a query | | chaos:query:pause |  | | chaos:theme:user | Ability to change the color scheme of the ChaosSearch UI | | kibana:* | Full access to Kibana permissions | | kibana-settings:read | Ability to access Visualizations and Dashboards | | kibana-settings:write | Ability to create Visualizations and Dashboards | | kibana-opendistro:* | Ability to create Alerts | | ui:* | Full access to the ChaosSearch UI | | ui:storage | Ability to access the Storage | | ui:refinery | Ability to access the Refinery | | ui:analytics | Ability to access the Analytics | | ui:dashboard | Ability to access the Dashboard |  | 
**Resources** | **string** | The resource(s) to which this container applies  | 

## Methods

### NewUserGroupsPermissions1

`func NewUserGroupsPermissions1(effect string, action string, resources string, ) *UserGroupsPermissions1`

NewUserGroupsPermissions1 instantiates a new UserGroupsPermissions1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupsPermissions1WithDefaults

`func NewUserGroupsPermissions1WithDefaults() *UserGroupsPermissions1`

NewUserGroupsPermissions1WithDefaults instantiates a new UserGroupsPermissions1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *UserGroupsPermissions1) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *UserGroupsPermissions1) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *UserGroupsPermissions1) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetAction

`func (o *UserGroupsPermissions1) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UserGroupsPermissions1) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UserGroupsPermissions1) SetAction(v string)`

SetAction sets Action field to given value.


### GetResources

`func (o *UserGroupsPermissions1) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UserGroupsPermissions1) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UserGroupsPermissions1) SetResources(v string)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


