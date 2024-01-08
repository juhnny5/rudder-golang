# GetAllCve200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetAllCve200ResponseData**](GetAllCve200ResponseData.md) |  | 

## Methods

### NewGetAllCve200Response

`func NewGetAllCve200Response(result string, action string, data GetAllCve200ResponseData, ) *GetAllCve200Response`

NewGetAllCve200Response instantiates a new GetAllCve200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllCve200ResponseWithDefaults

`func NewGetAllCve200ResponseWithDefaults() *GetAllCve200Response`

NewGetAllCve200ResponseWithDefaults instantiates a new GetAllCve200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetAllCve200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAllCve200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAllCve200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *GetAllCve200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetAllCve200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetAllCve200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *GetAllCve200Response) GetData() GetAllCve200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllCve200Response) GetDataOk() (*GetAllCve200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllCve200Response) SetData(v GetAllCve200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


