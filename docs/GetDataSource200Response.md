# GetDataSource200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetAllDataSources200ResponseData**](GetAllDataSources200ResponseData.md) |  | 

## Methods

### NewGetDataSource200Response

`func NewGetDataSource200Response(result string, action string, data GetAllDataSources200ResponseData, ) *GetDataSource200Response`

NewGetDataSource200Response instantiates a new GetDataSource200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDataSource200ResponseWithDefaults

`func NewGetDataSource200ResponseWithDefaults() *GetDataSource200Response`

NewGetDataSource200ResponseWithDefaults instantiates a new GetDataSource200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *GetDataSource200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDataSource200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDataSource200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *GetDataSource200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetDataSource200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetDataSource200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *GetDataSource200Response) GetData() GetAllDataSources200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDataSource200Response) GetDataOk() (*GetAllDataSources200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDataSource200Response) SetData(v GetAllDataSources200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


