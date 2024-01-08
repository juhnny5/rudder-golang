# CampaignDetailsInfoSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to [**CampaignScheduleMonthlyStart**](CampaignScheduleMonthlyStart.md) |  | [optional] 
**End** | Pointer to [**CampaignScheduleMonthlyEnd**](CampaignScheduleMonthlyEnd.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** | Week during the month in which our campaign start, 1/2/3 means first/second/last, -1/-2 means Last/secondLast | [optional] 

## Methods

### NewCampaignDetailsInfoSchedule

`func NewCampaignDetailsInfoSchedule() *CampaignDetailsInfoSchedule`

NewCampaignDetailsInfoSchedule instantiates a new CampaignDetailsInfoSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDetailsInfoScheduleWithDefaults

`func NewCampaignDetailsInfoScheduleWithDefaults() *CampaignDetailsInfoSchedule`

NewCampaignDetailsInfoScheduleWithDefaults instantiates a new CampaignDetailsInfoSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *CampaignDetailsInfoSchedule) GetStart() CampaignScheduleMonthlyStart`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CampaignDetailsInfoSchedule) GetStartOk() (*CampaignScheduleMonthlyStart, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CampaignDetailsInfoSchedule) SetStart(v CampaignScheduleMonthlyStart)`

SetStart sets Start field to given value.

### HasStart

`func (o *CampaignDetailsInfoSchedule) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *CampaignDetailsInfoSchedule) GetEnd() CampaignScheduleMonthlyEnd`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CampaignDetailsInfoSchedule) GetEndOk() (*CampaignScheduleMonthlyEnd, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CampaignDetailsInfoSchedule) SetEnd(v CampaignScheduleMonthlyEnd)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CampaignDetailsInfoSchedule) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetType

`func (o *CampaignDetailsInfoSchedule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignDetailsInfoSchedule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignDetailsInfoSchedule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CampaignDetailsInfoSchedule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPosition

`func (o *CampaignDetailsInfoSchedule) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CampaignDetailsInfoSchedule) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CampaignDetailsInfoSchedule) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CampaignDetailsInfoSchedule) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


