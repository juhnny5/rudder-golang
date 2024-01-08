# CampaignSystemUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reboot** | Pointer to **string** | Define the behavior after update | [optional] 
**Targets** | Pointer to [**[][]RuleTargetsInner**]([]RuleTargetsInner.md) | List of all  groups of node to target the campaign | [optional] 

## Methods

### NewCampaignSystemUpdate

`func NewCampaignSystemUpdate() *CampaignSystemUpdate`

NewCampaignSystemUpdate instantiates a new CampaignSystemUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSystemUpdateWithDefaults

`func NewCampaignSystemUpdateWithDefaults() *CampaignSystemUpdate`

NewCampaignSystemUpdateWithDefaults instantiates a new CampaignSystemUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReboot

`func (o *CampaignSystemUpdate) GetReboot() string`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *CampaignSystemUpdate) GetRebootOk() (*string, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *CampaignSystemUpdate) SetReboot(v string)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *CampaignSystemUpdate) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetTargets

`func (o *CampaignSystemUpdate) GetTargets() [][]RuleTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CampaignSystemUpdate) GetTargetsOk() (*[][]RuleTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CampaignSystemUpdate) SetTargets(v [][]RuleTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *CampaignSystemUpdate) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


