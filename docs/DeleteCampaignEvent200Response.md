# DeleteCampaignEvent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetAllCampaignEvents200ResponseData**](GetAllCampaignEvents200ResponseData.md) |  | 

## Methods

### NewDeleteCampaignEvent200Response

`func NewDeleteCampaignEvent200Response(result string, action string, data GetAllCampaignEvents200ResponseData, ) *DeleteCampaignEvent200Response`

NewDeleteCampaignEvent200Response instantiates a new DeleteCampaignEvent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCampaignEvent200ResponseWithDefaults

`func NewDeleteCampaignEvent200ResponseWithDefaults() *DeleteCampaignEvent200Response`

NewDeleteCampaignEvent200ResponseWithDefaults instantiates a new DeleteCampaignEvent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DeleteCampaignEvent200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeleteCampaignEvent200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeleteCampaignEvent200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *DeleteCampaignEvent200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeleteCampaignEvent200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeleteCampaignEvent200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *DeleteCampaignEvent200Response) GetData() GetAllCampaignEvents200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteCampaignEvent200Response) GetDataOk() (*GetAllCampaignEvents200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteCampaignEvent200Response) SetData(v GetAllCampaignEvents200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


