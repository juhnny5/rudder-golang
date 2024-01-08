# ModifyAllowedNetworksRequestAllowedNetworks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Add** | Pointer to **[]string** | List of networks to add to existing allowed networks for that server | [optional] 
**Delete** | Pointer to **[]string** | List of networks to remove from existing allowed networks for that server | [optional] 

## Methods

### NewModifyAllowedNetworksRequestAllowedNetworks

`func NewModifyAllowedNetworksRequestAllowedNetworks() *ModifyAllowedNetworksRequestAllowedNetworks`

NewModifyAllowedNetworksRequestAllowedNetworks instantiates a new ModifyAllowedNetworksRequestAllowedNetworks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyAllowedNetworksRequestAllowedNetworksWithDefaults

`func NewModifyAllowedNetworksRequestAllowedNetworksWithDefaults() *ModifyAllowedNetworksRequestAllowedNetworks`

NewModifyAllowedNetworksRequestAllowedNetworksWithDefaults instantiates a new ModifyAllowedNetworksRequestAllowedNetworks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdd

`func (o *ModifyAllowedNetworksRequestAllowedNetworks) GetAdd() []string`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *ModifyAllowedNetworksRequestAllowedNetworks) GetAddOk() (*[]string, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *ModifyAllowedNetworksRequestAllowedNetworks) SetAdd(v []string)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *ModifyAllowedNetworksRequestAllowedNetworks) HasAdd() bool`

HasAdd returns a boolean if a field has been set.

### GetDelete

`func (o *ModifyAllowedNetworksRequestAllowedNetworks) GetDelete() []string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ModifyAllowedNetworksRequestAllowedNetworks) GetDeleteOk() (*[]string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ModifyAllowedNetworksRequestAllowedNetworks) SetDelete(v []string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ModifyAllowedNetworksRequestAllowedNetworks) HasDelete() bool`

HasDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


