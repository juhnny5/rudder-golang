# UpdateUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**UpdateUser200ResponseData**](UpdateUser200ResponseData.md) |  | 

## Methods

### NewUpdateUser200Response

`func NewUpdateUser200Response(result string, action string, data UpdateUser200ResponseData, ) *UpdateUser200Response`

NewUpdateUser200Response instantiates a new UpdateUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUser200ResponseWithDefaults

`func NewUpdateUser200ResponseWithDefaults() *UpdateUser200Response`

NewUpdateUser200ResponseWithDefaults instantiates a new UpdateUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UpdateUser200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateUser200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateUser200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *UpdateUser200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateUser200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateUser200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *UpdateUser200Response) GetData() UpdateUser200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateUser200Response) GetDataOk() (*UpdateUser200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateUser200Response) SetData(v UpdateUser200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


