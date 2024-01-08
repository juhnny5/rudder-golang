# ScheduleCampaign200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetAllCampaignEvents200ResponseData**](GetAllCampaignEvents200ResponseData.md) |  | 

## Methods

### NewScheduleCampaign200Response

`func NewScheduleCampaign200Response(result string, action string, data GetAllCampaignEvents200ResponseData, ) *ScheduleCampaign200Response`

NewScheduleCampaign200Response instantiates a new ScheduleCampaign200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCampaign200ResponseWithDefaults

`func NewScheduleCampaign200ResponseWithDefaults() *ScheduleCampaign200Response`

NewScheduleCampaign200ResponseWithDefaults instantiates a new ScheduleCampaign200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ScheduleCampaign200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ScheduleCampaign200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ScheduleCampaign200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *ScheduleCampaign200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ScheduleCampaign200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ScheduleCampaign200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *ScheduleCampaign200Response) GetData() GetAllCampaignEvents200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ScheduleCampaign200Response) GetDataOk() (*GetAllCampaignEvents200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ScheduleCampaign200Response) SetData(v GetAllCampaignEvents200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


