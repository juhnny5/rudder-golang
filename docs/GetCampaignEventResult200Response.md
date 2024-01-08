# GetCampaignEventResult200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Id** | Pointer to **string** | Campaign event id | [optional] 
**Data** | [**GetCampaignEventResult200ResponseData**](GetCampaignEventResult200ResponseData.md) |  | 

## Methods

### NewGetCampaignEventResult200Response

`func NewGetCampaignEventResult200Response(result string, action string, data GetCampaignEventResult200ResponseData, ) *GetCampaignEventResult200Response`

NewGetCampaignEventResult200Response instantiates a new GetCampaignEventResult200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignEventResult200ResponseWithDefaults

`func NewGetCampaignEventResult200ResponseWithDefaults() *GetCampaignEventResult200Response`

NewGetCampaignEventResult200ResponseWithDefaults instantiates a new GetCampaignEventResult200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetCampaignEventResult200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetCampaignEventResult200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetCampaignEventResult200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *GetCampaignEventResult200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetCampaignEventResult200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetCampaignEventResult200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *GetCampaignEventResult200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCampaignEventResult200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCampaignEventResult200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetCampaignEventResult200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetData

`func (o *GetCampaignEventResult200Response) GetData() GetCampaignEventResult200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCampaignEventResult200Response) GetDataOk() (*GetCampaignEventResult200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCampaignEventResult200Response) SetData(v GetCampaignEventResult200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


