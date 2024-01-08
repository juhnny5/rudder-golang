# CampaignDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**CampaignDetailsInfo**](CampaignDetailsInfo.md) |  | [optional] 
**CampaignType** | Pointer to **string** | Type of campaign (with version determine what kind of campaign you handle and the details object) | [optional] 
**Version** | Pointer to **int32** | Version of campaign (with type determine what kind of campaign you handle and the details object) | [optional] 
**Details** | Pointer to [**CampaignDetailsDetails**](CampaignDetailsDetails.md) |  | [optional] 

## Methods

### NewCampaignDetails

`func NewCampaignDetails() *CampaignDetails`

NewCampaignDetails instantiates a new CampaignDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDetailsWithDefaults

`func NewCampaignDetailsWithDefaults() *CampaignDetails`

NewCampaignDetailsWithDefaults instantiates a new CampaignDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *CampaignDetails) GetInfo() CampaignDetailsInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CampaignDetails) GetInfoOk() (*CampaignDetailsInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CampaignDetails) SetInfo(v CampaignDetailsInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CampaignDetails) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCampaignType

`func (o *CampaignDetails) GetCampaignType() string`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *CampaignDetails) GetCampaignTypeOk() (*string, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *CampaignDetails) SetCampaignType(v string)`

SetCampaignType sets CampaignType field to given value.

### HasCampaignType

`func (o *CampaignDetails) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### GetVersion

`func (o *CampaignDetails) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CampaignDetails) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CampaignDetails) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CampaignDetails) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDetails

`func (o *CampaignDetails) GetDetails() CampaignDetailsDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CampaignDetails) GetDetailsOk() (*CampaignDetailsDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CampaignDetails) SetDetails(v CampaignDetailsDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CampaignDetails) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


