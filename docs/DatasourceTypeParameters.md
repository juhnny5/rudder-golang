# DatasourceTypeParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL to contact. Rudder expansion available. | [optional] 
**RequestMethod** | Pointer to **string** | HTTP method to use to contact the URL. | [optional] 
**Headers** | Pointer to [**[]DatasourceTypeParametersHeadersInner**](DatasourceTypeParametersHeadersInner.md) | Represent HTTP headers for the query. Rudder expansion available. | [optional] 
**Path** | Pointer to **string** | JSON path (as defined in [the specification](https://github.com/jayway/JsonPath/), without the leading &#x60;$.&#x60;) to find the interesting sub-json or string/number/boolean value in the answer. Let empty to use the whole answer as value. | [optional] 
**CheckSsl** | Pointer to **bool** | Check SSL certificate validity for https. Must be set to false for self-signed certificate | [optional] 
**RequestTimeout** | Pointer to **int32** | Timeout in seconds for each HTTP request | [optional] 
**RequestMode** | Pointer to [**DatasourceTypeParametersRequestMode**](DatasourceTypeParametersRequestMode.md) |  | [optional] 

## Methods

### NewDatasourceTypeParameters

`func NewDatasourceTypeParameters() *DatasourceTypeParameters`

NewDatasourceTypeParameters instantiates a new DatasourceTypeParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceTypeParametersWithDefaults

`func NewDatasourceTypeParametersWithDefaults() *DatasourceTypeParameters`

NewDatasourceTypeParametersWithDefaults instantiates a new DatasourceTypeParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DatasourceTypeParameters) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DatasourceTypeParameters) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DatasourceTypeParameters) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DatasourceTypeParameters) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequestMethod

`func (o *DatasourceTypeParameters) GetRequestMethod() string`

GetRequestMethod returns the RequestMethod field if non-nil, zero value otherwise.

### GetRequestMethodOk

`func (o *DatasourceTypeParameters) GetRequestMethodOk() (*string, bool)`

GetRequestMethodOk returns a tuple with the RequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMethod

`func (o *DatasourceTypeParameters) SetRequestMethod(v string)`

SetRequestMethod sets RequestMethod field to given value.

### HasRequestMethod

`func (o *DatasourceTypeParameters) HasRequestMethod() bool`

HasRequestMethod returns a boolean if a field has been set.

### GetHeaders

`func (o *DatasourceTypeParameters) GetHeaders() []DatasourceTypeParametersHeadersInner`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DatasourceTypeParameters) GetHeadersOk() (*[]DatasourceTypeParametersHeadersInner, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DatasourceTypeParameters) SetHeaders(v []DatasourceTypeParametersHeadersInner)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DatasourceTypeParameters) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPath

`func (o *DatasourceTypeParameters) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DatasourceTypeParameters) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DatasourceTypeParameters) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DatasourceTypeParameters) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCheckSsl

`func (o *DatasourceTypeParameters) GetCheckSsl() bool`

GetCheckSsl returns the CheckSsl field if non-nil, zero value otherwise.

### GetCheckSslOk

`func (o *DatasourceTypeParameters) GetCheckSslOk() (*bool, bool)`

GetCheckSslOk returns a tuple with the CheckSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSsl

`func (o *DatasourceTypeParameters) SetCheckSsl(v bool)`

SetCheckSsl sets CheckSsl field to given value.

### HasCheckSsl

`func (o *DatasourceTypeParameters) HasCheckSsl() bool`

HasCheckSsl returns a boolean if a field has been set.

### GetRequestTimeout

`func (o *DatasourceTypeParameters) GetRequestTimeout() int32`

GetRequestTimeout returns the RequestTimeout field if non-nil, zero value otherwise.

### GetRequestTimeoutOk

`func (o *DatasourceTypeParameters) GetRequestTimeoutOk() (*int32, bool)`

GetRequestTimeoutOk returns a tuple with the RequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeout

`func (o *DatasourceTypeParameters) SetRequestTimeout(v int32)`

SetRequestTimeout sets RequestTimeout field to given value.

### HasRequestTimeout

`func (o *DatasourceTypeParameters) HasRequestTimeout() bool`

HasRequestTimeout returns a boolean if a field has been set.

### GetRequestMode

`func (o *DatasourceTypeParameters) GetRequestMode() DatasourceTypeParametersRequestMode`

GetRequestMode returns the RequestMode field if non-nil, zero value otherwise.

### GetRequestModeOk

`func (o *DatasourceTypeParameters) GetRequestModeOk() (*DatasourceTypeParametersRequestMode, bool)`

GetRequestModeOk returns a tuple with the RequestMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMode

`func (o *DatasourceTypeParameters) SetRequestMode(v DatasourceTypeParametersRequestMode)`

SetRequestMode sets RequestMode field to given value.

### HasRequestMode

`func (o *DatasourceTypeParameters) HasRequestMode() bool`

HasRequestMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


