# UpdateCVERequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Url used to update CVE, will default to one set in config | [optional] 
**Years** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateCVERequest

`func NewUpdateCVERequest() *UpdateCVERequest`

NewUpdateCVERequest instantiates a new UpdateCVERequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCVERequestWithDefaults

`func NewUpdateCVERequestWithDefaults() *UpdateCVERequest`

NewUpdateCVERequestWithDefaults instantiates a new UpdateCVERequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UpdateCVERequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateCVERequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateCVERequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateCVERequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetYears

`func (o *UpdateCVERequest) GetYears() []string`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *UpdateCVERequest) GetYearsOk() (*[]string, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *UpdateCVERequest) SetYears(v []string)`

SetYears sets Years field to given value.

### HasYears

`func (o *UpdateCVERequest) HasYears() bool`

HasYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


