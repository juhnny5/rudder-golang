# CampaignDetailsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Campaign id | [optional] 
**Name** | Pointer to **string** | Campaign name | [optional] 
**Description** | Pointer to **string** | Campaign description | [optional] 
**Status** | Pointer to [**CampaignDetailsInfoStatus**](CampaignDetailsInfoStatus.md) |  | [optional] 
**Schedule** | Pointer to [**CampaignDetailsInfoSchedule**](CampaignDetailsInfoSchedule.md) |  | [optional] 

## Methods

### NewCampaignDetailsInfo

`func NewCampaignDetailsInfo() *CampaignDetailsInfo`

NewCampaignDetailsInfo instantiates a new CampaignDetailsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDetailsInfoWithDefaults

`func NewCampaignDetailsInfoWithDefaults() *CampaignDetailsInfo`

NewCampaignDetailsInfoWithDefaults instantiates a new CampaignDetailsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignDetailsInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignDetailsInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignDetailsInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignDetailsInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CampaignDetailsInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignDetailsInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignDetailsInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignDetailsInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CampaignDetailsInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignDetailsInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignDetailsInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignDetailsInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *CampaignDetailsInfo) GetStatus() CampaignDetailsInfoStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignDetailsInfo) GetStatusOk() (*CampaignDetailsInfoStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignDetailsInfo) SetStatus(v CampaignDetailsInfoStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CampaignDetailsInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSchedule

`func (o *CampaignDetailsInfo) GetSchedule() CampaignDetailsInfoSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CampaignDetailsInfo) GetScheduleOk() (*CampaignDetailsInfoSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CampaignDetailsInfo) SetSchedule(v CampaignDetailsInfoSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CampaignDetailsInfo) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


