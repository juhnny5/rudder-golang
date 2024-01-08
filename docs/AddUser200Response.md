# AddUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**AddUser200ResponseData**](AddUser200ResponseData.md) |  | 

## Methods

### NewAddUser200Response

`func NewAddUser200Response(result string, action string, data AddUser200ResponseData, ) *AddUser200Response`

NewAddUser200Response instantiates a new AddUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUser200ResponseWithDefaults

`func NewAddUser200ResponseWithDefaults() *AddUser200Response`

NewAddUser200ResponseWithDefaults instantiates a new AddUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *AddUser200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AddUser200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AddUser200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *AddUser200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AddUser200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AddUser200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *AddUser200Response) GetData() AddUser200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddUser200Response) GetDataOk() (*AddUser200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddUser200Response) SetData(v AddUser200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


