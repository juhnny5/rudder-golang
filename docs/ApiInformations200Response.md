# ApiInformations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Result of the request | 
**Action** | **string** | The id of the action | 
**Data** | [**ApiInformations200ResponseData**](ApiInformations200ResponseData.md) |  | 

## Methods

### NewApiInformations200Response

`func NewApiInformations200Response(result string, action string, data ApiInformations200ResponseData, ) *ApiInformations200Response`

NewApiInformations200Response instantiates a new ApiInformations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInformations200ResponseWithDefaults

`func NewApiInformations200ResponseWithDefaults() *ApiInformations200Response`

NewApiInformations200ResponseWithDefaults instantiates a new ApiInformations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ApiInformations200Response) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ApiInformations200Response) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ApiInformations200Response) SetResult(v string)`

SetResult sets Result field to given value.


### GetAction

`func (o *ApiInformations200Response) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ApiInformations200Response) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ApiInformations200Response) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *ApiInformations200Response) GetData() ApiInformations200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiInformations200Response) GetDataOk() (*ApiInformations200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiInformations200Response) SetData(v ApiInformations200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


