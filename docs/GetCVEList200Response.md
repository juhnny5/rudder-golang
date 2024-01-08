# GetCVEList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetAllCve200ResponseData**](GetAllCve200ResponseData.md) |  | 

## Methods

### NewGetCVEList200Response

`func NewGetCVEList200Response(result string, action string, data GetAllCve200ResponseData, ) *GetCVEList200Response`

NewGetCVEList200Response instantiates a new GetCVEList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCVEList200ResponseWithDefaults

`func NewGetCVEList200ResponseWithDefaults() *GetCVEList200Response`

NewGetCVEList200ResponseWithDefaults instantiates a new GetCVEList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetCVEList200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetCVEList200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetCVEList200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *GetCVEList200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetCVEList200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetCVEList200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *GetCVEList200Response) GetData() GetAllCve200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCVEList200Response) GetDataOk() (*GetAllCve200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCVEList200Response) SetData(v GetAllCve200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


