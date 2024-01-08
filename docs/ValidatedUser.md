# ValidatedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**IsValidated** | **bool** | whether the user&#39;s actions generate chanque-request or not | 
**UserExists** | **bool** | whether the user exists in file or not | 

## Methods

### NewValidatedUser

`func NewValidatedUser(username string, isValidated bool, userExists bool, ) *ValidatedUser`

NewValidatedUser instantiates a new ValidatedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatedUserWithDefaults

`func NewValidatedUserWithDefaults() *ValidatedUser`

NewValidatedUserWithDefaults instantiates a new ValidatedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ValidatedUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ValidatedUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ValidatedUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetIsValidated

`func (o *ValidatedUser) GetIsValidated() bool`

GetIsValidated returns the IsValidated field if non-nil, zero value otherwise.

### GetIsValidatedOk

`func (o *ValidatedUser) GetIsValidatedOk() (*bool, bool)`

GetIsValidatedOk returns a tuple with the IsValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValidated

`func (o *ValidatedUser) SetIsValidated(v bool)`

SetIsValidated sets IsValidated field to given value.


### GetUserExists

`func (o *ValidatedUser) GetUserExists() bool`

GetUserExists returns the UserExists field if non-nil, zero value otherwise.

### GetUserExistsOk

`func (o *ValidatedUser) GetUserExistsOk() (*bool, bool)`

GetUserExistsOk returns a tuple with the UserExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExists

`func (o *ValidatedUser) SetUserExists(v bool)`

SetUserExists sets UserExists field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


