# GetSystemUpdateResultForNode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Id** | Pointer to **string** | Campaign event id | [optional] 
**Data** | [**GetSystemUpdateResultForNode200ResponseData**](GetSystemUpdateResultForNode200ResponseData.md) |  | 

## Methods

### NewGetSystemUpdateResultForNode200Response

`func NewGetSystemUpdateResultForNode200Response(result string, action string, data GetSystemUpdateResultForNode200ResponseData, ) *GetSystemUpdateResultForNode200Response`

NewGetSystemUpdateResultForNode200Response instantiates a new GetSystemUpdateResultForNode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSystemUpdateResultForNode200ResponseWithDefaults

`func NewGetSystemUpdateResultForNode200ResponseWithDefaults() *GetSystemUpdateResultForNode200Response`

NewGetSystemUpdateResultForNode200ResponseWithDefaults instantiates a new GetSystemUpdateResultForNode200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetSystemUpdateResultForNode200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSystemUpdateResultForNode200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSystemUpdateResultForNode200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *GetSystemUpdateResultForNode200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetSystemUpdateResultForNode200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetSystemUpdateResultForNode200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *GetSystemUpdateResultForNode200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSystemUpdateResultForNode200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSystemUpdateResultForNode200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetSystemUpdateResultForNode200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetData

`func (o *GetSystemUpdateResultForNode200Response) GetData() GetSystemUpdateResultForNode200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSystemUpdateResultForNode200Response) GetDataOk() (*GetSystemUpdateResultForNode200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSystemUpdateResultForNode200Response) SetData(v GetSystemUpdateResultForNode200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


