# DeleteDataSource200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**GetAllDataSources200ResponseData**](GetAllDataSources200ResponseData.md) |  | 

## Methods

### NewDeleteDataSource200Response

`func NewDeleteDataSource200Response(result string, action string, data GetAllDataSources200ResponseData, ) *DeleteDataSource200Response`

NewDeleteDataSource200Response instantiates a new DeleteDataSource200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDataSource200ResponseWithDefaults

`func NewDeleteDataSource200ResponseWithDefaults() *DeleteDataSource200Response`

NewDeleteDataSource200ResponseWithDefaults instantiates a new DeleteDataSource200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DeleteDataSource200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeleteDataSource200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeleteDataSource200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *DeleteDataSource200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeleteDataSource200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeleteDataSource200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *DeleteDataSource200Response) GetData() GetAllDataSources200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteDataSource200Response) GetDataOk() (*GetAllDataSources200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteDataSource200Response) SetData(v GetAllDataSources200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


