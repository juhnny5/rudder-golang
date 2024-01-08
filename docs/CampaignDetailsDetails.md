# CampaignDetailsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reboot** | Pointer to **string** | Define the behavior after update | [optional] 
**Targets** | Pointer to [**[][]RuleTargetsInner**]([]RuleTargetsInner.md) | List of all  groups of node to target the campaign | [optional] 
**SoftwareUpdate** | Pointer to [**[]CampaignSoftwareUpdateSoftwareUpdateInner**](CampaignSoftwareUpdateSoftwareUpdateInner.md) | List of all software to update | [optional] 

## Methods

### NewCampaignDetailsDetails

`func NewCampaignDetailsDetails() *CampaignDetailsDetails`

NewCampaignDetailsDetails instantiates a new CampaignDetailsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDetailsDetailsWithDefaults

`func NewCampaignDetailsDetailsWithDefaults() *CampaignDetailsDetails`

NewCampaignDetailsDetailsWithDefaults instantiates a new CampaignDetailsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReboot

`func (o *CampaignDetailsDetails) GetReboot() string`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *CampaignDetailsDetails) GetRebootOk() (*string, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *CampaignDetailsDetails) SetReboot(v string)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *CampaignDetailsDetails) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetTargets

`func (o *CampaignDetailsDetails) GetTargets() [][]RuleTargetsInner`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CampaignDetailsDetails) GetTargetsOk() (*[][]RuleTargetsInner, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CampaignDetailsDetails) SetTargets(v [][]RuleTargetsInner)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *CampaignDetailsDetails) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetSoftwareUpdate

`func (o *CampaignDetailsDetails) GetSoftwareUpdate() []CampaignSoftwareUpdateSoftwareUpdateInner`

GetSoftwareUpdate returns the SoftwareUpdate field if non-nil, zero value otherwise.

### GetSoftwareUpdateOk

`func (o *CampaignDetailsDetails) GetSoftwareUpdateOk() (*[]CampaignSoftwareUpdateSoftwareUpdateInner, bool)`

GetSoftwareUpdateOk returns a tuple with the SoftwareUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdate

`func (o *CampaignDetailsDetails) SetSoftwareUpdate(v []CampaignSoftwareUpdateSoftwareUpdateInner)`

SetSoftwareUpdate sets SoftwareUpdate field to given value.

### HasSoftwareUpdate

`func (o *CampaignDetailsDetails) HasSoftwareUpdate() bool`

HasSoftwareUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


