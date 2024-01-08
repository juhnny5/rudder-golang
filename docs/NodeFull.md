# NodeFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique id of the node | 
**Hostname** | **string** | Fully qualified name of the node | 
**Status** | **string** | Status of the node | 
**ArchitectureDescription** | Pointer to **string** | Information about CPU architecture (32/64 bits) | [optional] 
**Description** | Pointer to **string** | A human intended description of the node (not used) | [optional] 
**IpAddresses** | **[]string** | IP addresses of the node (IPv4 and IPv6) | 
**LastRunDate** | Pointer to **string** | Date and time of the latest run, if one is available (node time). Up to API v11, format was: \&quot;YYYY-MM-dd HH:mm\&quot; | [optional] 
**LastInventoryDate** | Pointer to **string** | Date and time of the latest generated inventory, if one is available (node time). Up to API v11, format was: \&quot;YYYY-MM-dd HH:mm\&quot; | [optional] 
**Machine** | Pointer to [**NodeFullMachine**](NodeFullMachine.md) |  | [optional] 
**Os** | Pointer to [**NodeFullOs**](NodeFullOs.md) |  | [optional] 
**ManagementTechnology** | [**[]NodeFullManagementTechnologyInner**](NodeFullManagementTechnologyInner.md) | Management agents running on the node | 
**PolicyServerId** | **string** | Rudder policy server managing the node | 
**Properties** | [**[]NodeFullPropertiesInner**](NodeFullPropertiesInner.md) | Node properties (either set by user or filled by third party sources) | 
**PolicyMode** | Pointer to **string** | Rudder policy mode for the node (&#x60;default&#x60; follows the global configuration) | [optional] 
**Ram** | Pointer to **int32** | Size of RAM in MB | [optional] 
**Timezone** | Pointer to [**NodeFullTimezone**](NodeFullTimezone.md) |  | [optional] 
**Accounts** | Pointer to **[]string** | User accounts declared in the node | [optional] 
**Bios** | Pointer to [**NodeFullBios**](NodeFullBios.md) |  | [optional] 
**Controllers** | Pointer to [**[]NodeFullControllersInner**](NodeFullControllersInner.md) | Physical controller information | [optional] 
**EnvironmentVariables** | Pointer to [**[]NodeFullEnvironmentVariablesInner**](NodeFullEnvironmentVariablesInner.md) | Environment variables defined on the node in the context of the agent | [optional] 
**FileSystems** | Pointer to [**[]NodeFullFileSystemsInner**](NodeFullFileSystemsInner.md) | File system declared on the node | [optional] 
**ManagementTechnologyDetails** | Pointer to [**NodeFullManagementTechnologyDetails**](NodeFullManagementTechnologyDetails.md) |  | [optional] 
**Memories** | Pointer to [**[]NodeFullMemoriesInner**](NodeFullMemoriesInner.md) | Memory slots | [optional] 
**NetworkInterfaces** | Pointer to [**[]NodeFullNetworkInterfacesInner**](NodeFullNetworkInterfacesInner.md) | Detailed information about registered network interfaces on the node | [optional] 
**Ports** | Pointer to [**[]NodeFullPortsInner**](NodeFullPortsInner.md) | Physical port information objects | [optional] 
**Processes** | Pointer to [**[]NodeFullProcessesInner**](NodeFullProcessesInner.md) | Process running (at inventory time) | [optional] 
**Processors** | Pointer to [**[]NodeFullProcessorsInner**](NodeFullProcessorsInner.md) | CPU information | [optional] 
**Slots** | Pointer to [**[]NodeFullSlotsInner**](NodeFullSlotsInner.md) | Physical extension slot information | [optional] 
**Software** | Pointer to [**[]NodeFullSoftwareInner**](NodeFullSoftwareInner.md) | Software installed on the node (can have thousands items) | [optional] 
**SoftwareUpdate** | Pointer to [**[]NodeFullSoftwareUpdateInner**](NodeFullSoftwareUpdateInner.md) | Software that can be updated on that machine | [optional] 
**Sound** | Pointer to [**[]NodeFullSoundInner**](NodeFullSoundInner.md) | Sound card information | [optional] 
**Storage** | Pointer to [**[]NodeFullStorageInner**](NodeFullStorageInner.md) | Storage (disks) information objects | [optional] 
**Videos** | Pointer to [**[]NodeFullVideosInner**](NodeFullVideosInner.md) | Video card information | [optional] 
**VirtualMachines** | Pointer to [**[]NodeFullVirtualMachinesInner**](NodeFullVirtualMachinesInner.md) | Hosted virtual machine information | [optional] 

## Methods

### NewNodeFull

`func NewNodeFull(id string, hostname string, status string, ipAddresses []string, managementTechnology []NodeFullManagementTechnologyInner, policyServerId string, properties []NodeFullPropertiesInner, ) *NodeFull`

NewNodeFull instantiates a new NodeFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFullWithDefaults

`func NewNodeFullWithDefaults() *NodeFull`

NewNodeFullWithDefaults instantiates a new NodeFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NodeFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodeFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodeFull) SetId(v string)`

SetId sets Id field to given value.


### GetHostname

`func (o *NodeFull) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NodeFull) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NodeFull) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetStatus

`func (o *NodeFull) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeFull) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeFull) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetArchitectureDescription

`func (o *NodeFull) GetArchitectureDescription() string`

GetArchitectureDescription returns the ArchitectureDescription field if non-nil, zero value otherwise.

### GetArchitectureDescriptionOk

`func (o *NodeFull) GetArchitectureDescriptionOk() (*string, bool)`

GetArchitectureDescriptionOk returns a tuple with the ArchitectureDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectureDescription

`func (o *NodeFull) SetArchitectureDescription(v string)`

SetArchitectureDescription sets ArchitectureDescription field to given value.

### HasArchitectureDescription

`func (o *NodeFull) HasArchitectureDescription() bool`

HasArchitectureDescription returns a boolean if a field has been set.

### GetDescription

`func (o *NodeFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NodeFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NodeFull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NodeFull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddresses

`func (o *NodeFull) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *NodeFull) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *NodeFull) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.


### GetLastRunDate

`func (o *NodeFull) GetLastRunDate() string`

GetLastRunDate returns the LastRunDate field if non-nil, zero value otherwise.

### GetLastRunDateOk

`func (o *NodeFull) GetLastRunDateOk() (*string, bool)`

GetLastRunDateOk returns a tuple with the LastRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunDate

`func (o *NodeFull) SetLastRunDate(v string)`

SetLastRunDate sets LastRunDate field to given value.

### HasLastRunDate

`func (o *NodeFull) HasLastRunDate() bool`

HasLastRunDate returns a boolean if a field has been set.

### GetLastInventoryDate

`func (o *NodeFull) GetLastInventoryDate() string`

GetLastInventoryDate returns the LastInventoryDate field if non-nil, zero value otherwise.

### GetLastInventoryDateOk

`func (o *NodeFull) GetLastInventoryDateOk() (*string, bool)`

GetLastInventoryDateOk returns a tuple with the LastInventoryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInventoryDate

`func (o *NodeFull) SetLastInventoryDate(v string)`

SetLastInventoryDate sets LastInventoryDate field to given value.

### HasLastInventoryDate

`func (o *NodeFull) HasLastInventoryDate() bool`

HasLastInventoryDate returns a boolean if a field has been set.

### GetMachine

`func (o *NodeFull) GetMachine() NodeFullMachine`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *NodeFull) GetMachineOk() (*NodeFullMachine, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *NodeFull) SetMachine(v NodeFullMachine)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *NodeFull) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOs

`func (o *NodeFull) GetOs() NodeFullOs`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *NodeFull) GetOsOk() (*NodeFullOs, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *NodeFull) SetOs(v NodeFullOs)`

SetOs sets Os field to given value.

### HasOs

`func (o *NodeFull) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetManagementTechnology

`func (o *NodeFull) GetManagementTechnology() []NodeFullManagementTechnologyInner`

GetManagementTechnology returns the ManagementTechnology field if non-nil, zero value otherwise.

### GetManagementTechnologyOk

`func (o *NodeFull) GetManagementTechnologyOk() (*[]NodeFullManagementTechnologyInner, bool)`

GetManagementTechnologyOk returns a tuple with the ManagementTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementTechnology

`func (o *NodeFull) SetManagementTechnology(v []NodeFullManagementTechnologyInner)`

SetManagementTechnology sets ManagementTechnology field to given value.


### GetPolicyServerId

`func (o *NodeFull) GetPolicyServerId() string`

GetPolicyServerId returns the PolicyServerId field if non-nil, zero value otherwise.

### GetPolicyServerIdOk

`func (o *NodeFull) GetPolicyServerIdOk() (*string, bool)`

GetPolicyServerIdOk returns a tuple with the PolicyServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyServerId

`func (o *NodeFull) SetPolicyServerId(v string)`

SetPolicyServerId sets PolicyServerId field to given value.


### GetProperties

`func (o *NodeFull) GetProperties() []NodeFullPropertiesInner`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *NodeFull) GetPropertiesOk() (*[]NodeFullPropertiesInner, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *NodeFull) SetProperties(v []NodeFullPropertiesInner)`

SetProperties sets Properties field to given value.


### GetPolicyMode

`func (o *NodeFull) GetPolicyMode() string`

GetPolicyMode returns the PolicyMode field if non-nil, zero value otherwise.

### GetPolicyModeOk

`func (o *NodeFull) GetPolicyModeOk() (*string, bool)`

GetPolicyModeOk returns a tuple with the PolicyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMode

`func (o *NodeFull) SetPolicyMode(v string)`

SetPolicyMode sets PolicyMode field to given value.

### HasPolicyMode

`func (o *NodeFull) HasPolicyMode() bool`

HasPolicyMode returns a boolean if a field has been set.

### GetRam

`func (o *NodeFull) GetRam() int32`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *NodeFull) GetRamOk() (*int32, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *NodeFull) SetRam(v int32)`

SetRam sets Ram field to given value.

### HasRam

`func (o *NodeFull) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetTimezone

`func (o *NodeFull) GetTimezone() NodeFullTimezone`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *NodeFull) GetTimezoneOk() (*NodeFullTimezone, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *NodeFull) SetTimezone(v NodeFullTimezone)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *NodeFull) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetAccounts

`func (o *NodeFull) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *NodeFull) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *NodeFull) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *NodeFull) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetBios

`func (o *NodeFull) GetBios() NodeFullBios`

GetBios returns the Bios field if non-nil, zero value otherwise.

### GetBiosOk

`func (o *NodeFull) GetBiosOk() (*NodeFullBios, bool)`

GetBiosOk returns a tuple with the Bios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBios

`func (o *NodeFull) SetBios(v NodeFullBios)`

SetBios sets Bios field to given value.

### HasBios

`func (o *NodeFull) HasBios() bool`

HasBios returns a boolean if a field has been set.

### GetControllers

`func (o *NodeFull) GetControllers() []NodeFullControllersInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *NodeFull) GetControllersOk() (*[]NodeFullControllersInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *NodeFull) SetControllers(v []NodeFullControllersInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *NodeFull) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *NodeFull) GetEnvironmentVariables() []NodeFullEnvironmentVariablesInner`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *NodeFull) GetEnvironmentVariablesOk() (*[]NodeFullEnvironmentVariablesInner, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *NodeFull) SetEnvironmentVariables(v []NodeFullEnvironmentVariablesInner)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *NodeFull) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetFileSystems

`func (o *NodeFull) GetFileSystems() []NodeFullFileSystemsInner`

GetFileSystems returns the FileSystems field if non-nil, zero value otherwise.

### GetFileSystemsOk

`func (o *NodeFull) GetFileSystemsOk() (*[]NodeFullFileSystemsInner, bool)`

GetFileSystemsOk returns a tuple with the FileSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSystems

`func (o *NodeFull) SetFileSystems(v []NodeFullFileSystemsInner)`

SetFileSystems sets FileSystems field to given value.

### HasFileSystems

`func (o *NodeFull) HasFileSystems() bool`

HasFileSystems returns a boolean if a field has been set.

### GetManagementTechnologyDetails

`func (o *NodeFull) GetManagementTechnologyDetails() NodeFullManagementTechnologyDetails`

GetManagementTechnologyDetails returns the ManagementTechnologyDetails field if non-nil, zero value otherwise.

### GetManagementTechnologyDetailsOk

`func (o *NodeFull) GetManagementTechnologyDetailsOk() (*NodeFullManagementTechnologyDetails, bool)`

GetManagementTechnologyDetailsOk returns a tuple with the ManagementTechnologyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementTechnologyDetails

`func (o *NodeFull) SetManagementTechnologyDetails(v NodeFullManagementTechnologyDetails)`

SetManagementTechnologyDetails sets ManagementTechnologyDetails field to given value.

### HasManagementTechnologyDetails

`func (o *NodeFull) HasManagementTechnologyDetails() bool`

HasManagementTechnologyDetails returns a boolean if a field has been set.

### GetMemories

`func (o *NodeFull) GetMemories() []NodeFullMemoriesInner`

GetMemories returns the Memories field if non-nil, zero value otherwise.

### GetMemoriesOk

`func (o *NodeFull) GetMemoriesOk() (*[]NodeFullMemoriesInner, bool)`

GetMemoriesOk returns a tuple with the Memories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemories

`func (o *NodeFull) SetMemories(v []NodeFullMemoriesInner)`

SetMemories sets Memories field to given value.

### HasMemories

`func (o *NodeFull) HasMemories() bool`

HasMemories returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *NodeFull) GetNetworkInterfaces() []NodeFullNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *NodeFull) GetNetworkInterfacesOk() (*[]NodeFullNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *NodeFull) SetNetworkInterfaces(v []NodeFullNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *NodeFull) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetPorts

`func (o *NodeFull) GetPorts() []NodeFullPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *NodeFull) GetPortsOk() (*[]NodeFullPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *NodeFull) SetPorts(v []NodeFullPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *NodeFull) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProcesses

`func (o *NodeFull) GetProcesses() []NodeFullProcessesInner`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *NodeFull) GetProcessesOk() (*[]NodeFullProcessesInner, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *NodeFull) SetProcesses(v []NodeFullProcessesInner)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *NodeFull) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### GetProcessors

`func (o *NodeFull) GetProcessors() []NodeFullProcessorsInner`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *NodeFull) GetProcessorsOk() (*[]NodeFullProcessorsInner, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *NodeFull) SetProcessors(v []NodeFullProcessorsInner)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *NodeFull) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetSlots

`func (o *NodeFull) GetSlots() []NodeFullSlotsInner`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *NodeFull) GetSlotsOk() (*[]NodeFullSlotsInner, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *NodeFull) SetSlots(v []NodeFullSlotsInner)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *NodeFull) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetSoftware

`func (o *NodeFull) GetSoftware() []NodeFullSoftwareInner`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *NodeFull) GetSoftwareOk() (*[]NodeFullSoftwareInner, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *NodeFull) SetSoftware(v []NodeFullSoftwareInner)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *NodeFull) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetSoftwareUpdate

`func (o *NodeFull) GetSoftwareUpdate() []NodeFullSoftwareUpdateInner`

GetSoftwareUpdate returns the SoftwareUpdate field if non-nil, zero value otherwise.

### GetSoftwareUpdateOk

`func (o *NodeFull) GetSoftwareUpdateOk() (*[]NodeFullSoftwareUpdateInner, bool)`

GetSoftwareUpdateOk returns a tuple with the SoftwareUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdate

`func (o *NodeFull) SetSoftwareUpdate(v []NodeFullSoftwareUpdateInner)`

SetSoftwareUpdate sets SoftwareUpdate field to given value.

### HasSoftwareUpdate

`func (o *NodeFull) HasSoftwareUpdate() bool`

HasSoftwareUpdate returns a boolean if a field has been set.

### GetSound

`func (o *NodeFull) GetSound() []NodeFullSoundInner`

GetSound returns the Sound field if non-nil, zero value otherwise.

### GetSoundOk

`func (o *NodeFull) GetSoundOk() (*[]NodeFullSoundInner, bool)`

GetSoundOk returns a tuple with the Sound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSound

`func (o *NodeFull) SetSound(v []NodeFullSoundInner)`

SetSound sets Sound field to given value.

### HasSound

`func (o *NodeFull) HasSound() bool`

HasSound returns a boolean if a field has been set.

### GetStorage

`func (o *NodeFull) GetStorage() []NodeFullStorageInner`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *NodeFull) GetStorageOk() (*[]NodeFullStorageInner, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *NodeFull) SetStorage(v []NodeFullStorageInner)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *NodeFull) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetVideos

`func (o *NodeFull) GetVideos() []NodeFullVideosInner`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *NodeFull) GetVideosOk() (*[]NodeFullVideosInner, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *NodeFull) SetVideos(v []NodeFullVideosInner)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *NodeFull) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetVirtualMachines

`func (o *NodeFull) GetVirtualMachines() []NodeFullVirtualMachinesInner`

GetVirtualMachines returns the VirtualMachines field if non-nil, zero value otherwise.

### GetVirtualMachinesOk

`func (o *NodeFull) GetVirtualMachinesOk() (*[]NodeFullVirtualMachinesInner, bool)`

GetVirtualMachinesOk returns a tuple with the VirtualMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachines

`func (o *NodeFull) SetVirtualMachines(v []NodeFullVirtualMachinesInner)`

SetVirtualMachines sets VirtualMachines field to given value.

### HasVirtualMachines

`func (o *NodeFull) HasVirtualMachines() bool`

HasVirtualMachines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


