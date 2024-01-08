# Users

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPreHahed** | Pointer to **bool** | If you want to provide hashed password set this property to &#x60;true&#x60; otherwise we will hash the plain password and store the hash | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** | this password will be hashed for you if the &#x60;isPreHashed&#x60; is set on false | [optional] 
**Role** | Pointer to **[]string** | Defined user&#39;s permissions | [optional] 

## Methods

### NewUsers

`func NewUsers() *Users`

NewUsers instantiates a new Users object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersWithDefaults

`func NewUsersWithDefaults() *Users`

NewUsersWithDefaults instantiates a new Users object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPreHahed

`func (o *Users) GetIsPreHahed() bool`

GetIsPreHahed returns the IsPreHahed field if non-nil, zero value otherwise.

### GetIsPreHahedOk

`func (o *Users) GetIsPreHahedOk() (*bool, bool)`

GetIsPreHahedOk returns a tuple with the IsPreHahed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreHahed

`func (o *Users) SetIsPreHahed(v bool)`

SetIsPreHahed sets IsPreHahed field to given value.

### HasIsPreHahed

`func (o *Users) HasIsPreHahed() bool`

HasIsPreHahed returns a boolean if a field has been set.

### GetUsername

`func (o *Users) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Users) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Users) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Users) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *Users) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Users) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Users) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Users) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRole

`func (o *Users) GetRole() []string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Users) GetRoleOk() (*[]string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Users) SetRole(v []string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Users) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


