# SaveCampaignEvent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetAllCampaignEvents200ResponseData**](GetAllCampaignEvents200ResponseData.md) |  | 

## Methods

### NewSaveCampaignEvent200Response

`func NewSaveCampaignEvent200Response(result string, action string, data GetAllCampaignEvents200ResponseData, ) *SaveCampaignEvent200Response`

NewSaveCampaignEvent200Response instantiates a new SaveCampaignEvent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveCampaignEvent200ResponseWithDefaults

`func NewSaveCampaignEvent200ResponseWithDefaults() *SaveCampaignEvent200Response`

NewSaveCampaignEvent200ResponseWithDefaults instantiates a new SaveCampaignEvent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *SaveCampaignEvent200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SaveCampaignEvent200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SaveCampaignEvent200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *SaveCampaignEvent200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SaveCampaignEvent200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SaveCampaignEvent200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *SaveCampaignEvent200Response) GetData() GetAllCampaignEvents200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SaveCampaignEvent200Response) GetDataOk() (*GetAllCampaignEvents200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SaveCampaignEvent200Response) SetData(v GetAllCampaignEvents200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


