# DeleteUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**DeleteUser200ResponseData**](DeleteUser200ResponseData.md) |  | 

## Methods

### NewDeleteUser200Response

`func NewDeleteUser200Response(result string, action string, data DeleteUser200ResponseData, ) *DeleteUser200Response`

NewDeleteUser200Response instantiates a new DeleteUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUser200ResponseWithDefaults

`func NewDeleteUser200ResponseWithDefaults() *DeleteUser200Response`

NewDeleteUser200ResponseWithDefaults instantiates a new DeleteUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *DeleteUser200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DeleteUser200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DeleteUser200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *DeleteUser200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DeleteUser200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DeleteUser200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *DeleteUser200Response) GetData() DeleteUser200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteUser200Response) GetDataOk() (*DeleteUser200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteUser200Response) SetData(v DeleteUser200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


