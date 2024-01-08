# GetUserInfo200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**Users** | [**[]Users**](Users.md) |  | 

## Methods

### NewGetUserInfo200ResponseData

`func NewGetUserInfo200ResponseData(digest string, users []Users, ) *GetUserInfo200ResponseData`

NewGetUserInfo200ResponseData instantiates a new GetUserInfo200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserInfo200ResponseDataWithDefaults

`func NewGetUserInfo200ResponseDataWithDefaults() *GetUserInfo200ResponseData`

NewGetUserInfo200ResponseDataWithDefaults instantiates a new GetUserInfo200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *GetUserInfo200ResponseData) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *GetUserInfo200ResponseData) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *GetUserInfo200ResponseData) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetUsers

`func (o *GetUserInfo200ResponseData) GetUsers() []Users`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetUserInfo200ResponseData) GetUsersOk() (*[]Users, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetUserInfo200ResponseData) SetUsers(v []Users)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


