# ApiSubInformations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**ApiGeneralInformations200ResponseData**](ApiGeneralInformations200ResponseData.md) |  | 

## Methods

### NewApiSubInformations200Response

`func NewApiSubInformations200Response(result string, action string, data ApiGeneralInformations200ResponseData, ) *ApiSubInformations200Response`

NewApiSubInformations200Response instantiates a new ApiSubInformations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSubInformations200ResponseWithDefaults

`func NewApiSubInformations200ResponseWithDefaults() *ApiSubInformations200Response`

NewApiSubInformations200ResponseWithDefaults instantiates a new ApiSubInformations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ApiSubInformations200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ApiSubInformations200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ApiSubInformations200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *ApiSubInformations200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ApiSubInformations200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ApiSubInformations200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *ApiSubInformations200Response) GetData() ApiGeneralInformations200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiSubInformations200Response) GetDataOk() (*ApiGeneralInformations200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiSubInformations200Response) SetData(v ApiGeneralInformations200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


