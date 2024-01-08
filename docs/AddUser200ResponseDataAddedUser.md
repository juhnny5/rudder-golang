# AddUser200ResponseDataAddedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** | this password will be hashed for you if the &#x60;isPreHashed&#x60; is set on false | [optional] 
**Role** | Pointer to **[]string** | defined user&#39;s permissions | [optional] 

## Methods

### NewAddUser200ResponseDataAddedUser

`func NewAddUser200ResponseDataAddedUser() *AddUser200ResponseDataAddedUser`

NewAddUser200ResponseDataAddedUser instantiates a new AddUser200ResponseDataAddedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUser200ResponseDataAddedUserWithDefaults

`func NewAddUser200ResponseDataAddedUserWithDefaults() *AddUser200ResponseDataAddedUser`

NewAddUser200ResponseDataAddedUserWithDefaults instantiates a new AddUser200ResponseDataAddedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AddUser200ResponseDataAddedUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddUser200ResponseDataAddedUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddUser200ResponseDataAddedUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddUser200ResponseDataAddedUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AddUser200ResponseDataAddedUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddUser200ResponseDataAddedUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddUser200ResponseDataAddedUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddUser200ResponseDataAddedUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRole

`func (o *AddUser200ResponseDataAddedUser) GetRole() []string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddUser200ResponseDataAddedUser) GetRoleOk() (*[]string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddUser200ResponseDataAddedUser) SetRole(v []string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AddUser200ResponseDataAddedUser) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


