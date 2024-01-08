# NodeFullNetworkInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Interface name (for ex \&quot;eth0\&quot;) | [optional] 
**Mask** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** | Information about the type of network interface | [optional] 
**Speed** | Pointer to **string** | Information about synchronization speed | [optional] 
**Status** | Pointer to **string** | network interface status (enabled or not, up or down) | [optional] 
**DhcpServer** | Pointer to **string** | DHCP server managing that network interface | [optional] 
**MacAddress** | Pointer to **string** | MAC address of the network interface | [optional] 
**IpAddresses** | Pointer to **[]string** | IP addresses of the network interface | [optional] 

## Methods

### NewNodeFullNetworkInterfacesInner

`func NewNodeFullNetworkInterfacesInner() *NodeFullNetworkInterfacesInner`

NewNodeFullNetworkInterfacesInner instantiates a new NodeFullNetworkInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullNetworkInterfacesInnerWithDefaults

`func NewNodeFullNetworkInterfacesInnerWithDefaults() *NodeFullNetworkInterfacesInner`

NewNodeFullNetworkInterfacesInnerWithDefaults instantiates a new NodeFullNetworkInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodeFullNetworkInterfacesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeFullNetworkInterfacesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeFullNetworkInterfacesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeFullNetworkInterfacesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMask

`func (o *NodeFullNetworkInterfacesInner) GetMask() []string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *NodeFullNetworkInterfacesInner) GetMaskOk() (*[]string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *NodeFullNetworkInterfacesInner) SetMask(v []string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *NodeFullNetworkInterfacesInner) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetType

`func (o *NodeFullNetworkInterfacesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeFullNetworkInterfacesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeFullNetworkInterfacesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodeFullNetworkInterfacesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSpeed

`func (o *NodeFullNetworkInterfacesInner) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *NodeFullNetworkInterfacesInner) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *NodeFullNetworkInterfacesInner) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *NodeFullNetworkInterfacesInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *NodeFullNetworkInterfacesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeFullNetworkInterfacesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeFullNetworkInterfacesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeFullNetworkInterfacesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDhcpServer

`func (o *NodeFullNetworkInterfacesInner) GetDhcpServer() string`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *NodeFullNetworkInterfacesInner) GetDhcpServerOk() (*string, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *NodeFullNetworkInterfacesInner) SetDhcpServer(v string)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *NodeFullNetworkInterfacesInner) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetMacAddress

`func (o *NodeFullNetworkInterfacesInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NodeFullNetworkInterfacesInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NodeFullNetworkInterfacesInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NodeFullNetworkInterfacesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetIpAddresses

`func (o *NodeFullNetworkInterfacesInner) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *NodeFullNetworkInterfacesInner) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *NodeFullNetworkInterfacesInner) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *NodeFullNetworkInterfacesInner) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


