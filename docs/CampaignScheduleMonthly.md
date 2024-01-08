# CampaignScheduleMonthly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to [**CampaignScheduleMonthlyStart**](CampaignScheduleMonthlyStart.md) |  | [optional] 
**End** | Pointer to [**CampaignScheduleMonthlyEnd**](CampaignScheduleMonthlyEnd.md) |  | [optional] 
**Position** | Pointer to **int32** | Week during the month in which our campaign start, 1/2/3 means first/second/last, -1/-2 means Last/secondLast | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCampaignScheduleMonthly

`func NewCampaignScheduleMonthly() *CampaignScheduleMonthly`

NewCampaignScheduleMonthly instantiates a new CampaignScheduleMonthly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignScheduleMonthlyWithDefaults

`func NewCampaignScheduleMonthlyWithDefaults() *CampaignScheduleMonthly`

NewCampaignScheduleMonthlyWithDefaults instantiates a new CampaignScheduleMonthly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *CampaignScheduleMonthly) GetStart() CampaignScheduleMonthlyStart`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CampaignScheduleMonthly) GetStartOk() (*CampaignScheduleMonthlyStart, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CampaignScheduleMonthly) SetStart(v CampaignScheduleMonthlyStart)`

SetStart sets Start field to given value.

### HasStart

`func (o *CampaignScheduleMonthly) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *CampaignScheduleMonthly) GetEnd() CampaignScheduleMonthlyEnd`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CampaignScheduleMonthly) GetEndOk() (*CampaignScheduleMonthlyEnd, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CampaignScheduleMonthly) SetEnd(v CampaignScheduleMonthlyEnd)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CampaignScheduleMonthly) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetPosition

`func (o *CampaignScheduleMonthly) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CampaignScheduleMonthly) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CampaignScheduleMonthly) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CampaignScheduleMonthly) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetType

`func (o *CampaignScheduleMonthly) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignScheduleMonthly) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignScheduleMonthly) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CampaignScheduleMonthly) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


