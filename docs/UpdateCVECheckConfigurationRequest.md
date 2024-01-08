# UpdateCVECheckConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Url used to check CVE | [optional] 
**ApiKey** | Pointer to **string** | Token used by to contact the API to check CVE | [optional] 

## Methods

### NewUpdateCVECheckConfigurationRequest

`func NewUpdateCVECheckConfigurationRequest() *UpdateCVECheckConfigurationRequest`

NewUpdateCVECheckConfigurationRequest instantiates a new UpdateCVECheckConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCVECheckConfigurationRequestWithDefaults

`func NewUpdateCVECheckConfigurationRequestWithDefaults() *UpdateCVECheckConfigurationRequest`

NewUpdateCVECheckConfigurationRequestWithDefaults instantiates a new UpdateCVECheckConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UpdateCVECheckConfigurationRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateCVECheckConfigurationRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateCVECheckConfigurationRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateCVECheckConfigurationRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetApiKey

`func (o *UpdateCVECheckConfigurationRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateCVECheckConfigurationRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateCVECheckConfigurationRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UpdateCVECheckConfigurationRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


