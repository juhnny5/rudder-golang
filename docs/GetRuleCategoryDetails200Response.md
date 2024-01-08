# GetRuleCategoryDetails200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetRuleCategoryDetails200ResponseData**](GetRuleCategoryDetails200ResponseData.md) |  | 

## Methods

### NewGetRuleCategoryDetails200Response

`func NewGetRuleCategoryDetails200Response(result string, action string, data GetRuleCategoryDetails200ResponseData, ) *GetRuleCategoryDetails200Response`

NewGetRuleCategoryDetails200Response instantiates a new GetRuleCategoryDetails200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRuleCategoryDetails200ResponseWithDefaults

`func NewGetRuleCategoryDetails200ResponseWithDefaults() *GetRuleCategoryDetails200Response`

NewGetRuleCategoryDetails200ResponseWithDefaults instantiates a new GetRuleCategoryDetails200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetRuleCategoryDetails200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRuleCategoryDetails200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRuleCategoryDetails200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *GetRuleCategoryDetails200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetRuleCategoryDetails200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetRuleCategoryDetails200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *GetRuleCategoryDetails200Response) GetData() GetRuleCategoryDetails200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRuleCategoryDetails200Response) GetDataOk() (*GetRuleCategoryDetails200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRuleCategoryDetails200Response) SetData(v GetRuleCategoryDetails200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


