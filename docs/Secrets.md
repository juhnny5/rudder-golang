# Secrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the secret used as a reference on the value | [optional] 
**Description** | Pointer to **string** | The description of the secret to identify it more easily | [optional] 
**Value** | Pointer to **string** | The value of the secret it will not be exposed in the interface | [optional] 

## Methods

### NewSecrets

`func NewSecrets() *Secrets`

NewSecrets instantiates a new Secrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretsWithDefaults

`func NewSecretsWithDefaults() *Secrets`

NewSecretsWithDefaults instantiates a new Secrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Secrets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Secrets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Secrets) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Secrets) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Secrets) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Secrets) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Secrets) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Secrets) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValue

`func (o *Secrets) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Secrets) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Secrets) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Secrets) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


