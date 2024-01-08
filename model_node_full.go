/*
Rudder API

Download OpenAPI specification: [openapi.yml](openapi.yml)  **Other documentation sources**:  * [Main documentation](https://docs.rudder.io) * [Internal relay API](https://docs.rudder.io/api/relay/)  # Introduction  Rudder exposes a REST API, enabling the user to interact with Rudder without using the webapp, for example, in scripts or cron jobs.  ## Authentication  The Rudder REST API uses simple API keys for authentication. All requests must be authenticated (except from a generic status API). The tokens are 32-character strings, passed in a `X-API-Token` header, like in:   ```bash curl --header \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules ```  The tokens are the API equivalent of a password, and must be secured just like a password would be.  ### API accounts  The accounts are managed in the Web interface. There are two possible types of accounts:  * **Global API accounts**: they are not linked to a Rudder user, and are managed by Rudder administrators in the _Administration -> API accounts_ page. You should define an expiration date whenever possible.  ![General API tokens settings](assets/api-tokens.png \"General API tokens settings\")  * **User tokens**: they are linked to a Rudder user, and give the same rights the user has. There can be only one token by user. This feature is provided by the `api-authorizatons` plugin.  ![User API token](assets/api-user.png \"User API token\")  When an action produces a change of configuration on the server, the API account that made it will be recorded in the event log, like for a Web interaction.  ### Authorization  When using Rudder without the `api-authorizatons` plugin, only global accounts are available, with two possible privilege levels, read-only or write. With the `api-authorizatons` plugin, you also get access to:  * User tokens, which have the same permissions as the user, using the Rudder roles and permissions feature. * Custom ACLs on global API accounts. They provide fine-grained permissions on every endpoint:  ![Custom API ACL](assets/custom-acl.png \"Custom API ACL\")  As a general principle, you should create dedicated tokens with the least privilege level for each different interaction you have with the API. This limits the risks of exploitation if a token is stolen, and allows tracking the activity of each token separately. Token renewal is also easier when they are only used for a limited purpose.  ## Versioning  Each time the API is extended with new features (new functions, new parameters, new responses, ...), it will be assigned a new version number. This will allow you to keep your existing scripts (based on previous behavior). Versions will always be integers (no 2.1 or 3.3, just 2, 3, 4, ...) or `latest`.  You can change the version of the API used by setting it either within the url or in a header:  * the URL: each URL is prefixed by its version id, like `/api/version/function`.  ```bash # Version 10 curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/10/rules # Latest curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules # Wrong (not an integer) => 404 not found curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/3.14/rules ```  * the HTTP headers. You can add the **X-API-Version** header to your request. The value needs to be an integer or `latest`.  ```bash # Version 10 curl -X GET -H \"X-API-Token: yourToken\" -H \"X-API-Version: 10\" https://rudder.example.com/rudder/api/rules # Wrong => Error response indicating which versions are available curl -X GET -H \"X-API-Token: yourToken\" -H \"X-API-Version: 3.14\" https://rudder.example.com/rudder/api/rules ```  In the future, we may declare some versions as deprecated, in order to remove them in a later version of Rudder, but we will never remove any versions without warning, or without a safe period of time to allow migration from previous versions.   <h4>Existing versions</h4> <table>   <thead>     <tr>       <th style=\"width: 20%\">Version</th>       <th style=\"width: 20%\">Rudder versions it appeared in</th>       <th style=\"width: 70%\">Description</th>     </tr>   </thead>   <tbody>     <tr>       <td class=\"code\">1</td>       <td class=\"code\">Never released (for internal use only)</td>       <td>Experimental version</td>     </tr>     <tr>       <td class=\"code\">2 to 10 (deprecated)</td>       <td class=\"code\">4.3 and before</td>       <td>These versions provided the core set of API features for rules, directives, nodes global parameters, change requests and compliance, rudder settings, and system API</td>     </tr>     <tr>       <td class=\"code\">11</td>       <td class=\"code\">5.0</td>       <td>New system API (replacing old localhost v1 api): status, maintenance operations and server behavior</td>     </tr>     <tr>       <td class=\"code\">12</td>       <td class=\"code\">6.0 and 6.1</td>       <td>Node key management</td>     </tr>     <tr>       <td class=\"code\">13</td>       <td class=\"code\">6.2</td>       <td><ul>         <li>Node status endpoint</li>         <li>System health check</li>         <li>System maintenance job to purge software [that endpoint was back-ported in 6.1]</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">14</td>       <td class=\"code\">7.0</td>       <td><ul>         <li>Secret management</li>         <li>Directive tree</li>         <li>Improve techniques management</li>         <li>Demote a relay</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">15</td>       <td class=\"code\">7.1</td>       <td><ul>         <li>Package updates in nodes</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">16</td>       <td class=\"code\">7.2</td>       <td><ul>         <li>Create node API included from plugin</li>         <li>Configuration archive import/export</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">17</td>       <td class=\"code\">7.3</td>       <td><ul>         <li>Compliance by directive</li>         <li>Path campaigns API included</li>       </ul></td>     </tr>     <tr>       <td class=\"code\">18</td>       <td class=\"code\">8.0</td>       <td><ul>         <li>Allowed network </li>         <li>Improve the structure of `/settings/allowed_networks` output</li>       </ul></td>     </tr>   </tbody> </table>   ## Response format  All responses from the API are in the JSON format.  ```json {   \"action\": \"The name of the called function\",   \"id\": \"The ID of the element you want, if relevant\",   \"result\": \"The result of your action: success or error\",   \"data\": \"Only present if this is a success and depends on the function, it's usually a JSON object\",   \"errorDetails\": \"Only present if this is an error, it contains the error message\" } ```   * __Success__ responses are sent with the 200 HTTP (Success) code  * __Error__ responses are sent with a HTTP error code (mostly 5xx...)   ## HTTP method  Rudder's REST API is based on the usage of [HTTP methods](http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html). We use them to indicate what action will be done by the request. Currently, we use four of them:   * **GET**: search or retrieve information (get rule details, get a group, ...)  * **PUT**: add new objects (create a directive, clone a Rule, ...)  * **DELETE**: remove objects (delete a node, delete a parameter, ...)  * **POST**: update existing objects (update a directive, reload a group, ...)   ## Parameters  ### General parameters  Some parameters are available for almost all API functions. They will be described in this section. They must be part of the query and can't be submitted in a JSON form.  #### Available for all requests  <table>   <thead>     <tr>       <th style=\"width: 30%\">Field</th>       <th style=\"width: 10%\">Type</th>       <th style=\"width: 70%\">Description</th>     </tr>   </thead>   <tbody>     <tr>       <td class=\"code\">prettify</td>       <td><b>boolean</b><br><i>optional</i></td>       <td>         Determine if the answer should be prettified (human friendly) or not. We recommend using this for debugging purposes, but not for general script usage as this does add some unnecessary load on the server side.         <p class=\"default-value\">Default value: <code>false</code></p>       </td>     </tr>   </tbody> </table>   #### Available for modification requests (PUT/POST/DELETE)  <table>   <thead>     <tr>       <th style=\"width: 25%\">Field</th>       <th style=\"width: 12%\">Type</th>       <th style=\"width: 70%\">Description</th>     </tr>   </thead>   <tbody>     <tr>       <td class=\"code\">reason</td>       <td><b>string</b><br><i>optional</i> or <i>required</i></td>       <td>         Set a message to explain the change. If you set the reason messages to be mandatory in the web interface, failing to supply this value will lead to an error.         <p class=\"default-value\">Default value: <code>\"\"</code></p>       </td>     </tr>     <tr>       <td class=\"code\">changeRequestName</td>       <td><b>string</b><br><i>optional</i></td>       <td>         Set the change request name, is used only if workflows are enabled. The default value depends on the function called         <p class=\"default-value\">Default value: <code>A default string for each function</code></p>       </td>     </tr>     <tr>       <td class=\"code\">changeRequestDescription</td>       <td><b>string</b><br><i>optional</i></td>       <td>         Set the change request description, is used only if workflows are enabled.         <p class=\"default-value\">Default value: <code>\"\"</code></p>       </td>     </tr>   </tbody> </table>   ### Passing parameters  Parameters to the API can be sent:  * As part of the URL for resource identification  * As data for POST/PUT requests    * Directly in JSON format    * As request arguments  #### As part of the URL for resource identification  Parameters in URLs are used to indicate which resource you want to interact with. The function will not work if this resource is missing.  ```bash # Get the Rule of ID \"id\" curl -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules/id ```  CAUTION: To avoid surprising behavior, do not put a '/' at the end of a URL: it would be interpreted as '/[empty string parameter]' and redirected to '/index', likely not what you wanted to do.   #### Sending data for POST/PUT requests  ##### Directly in JSON format  JSON format is the preferred way to interact with Rudder API for creating or updating resources. You'll also have to set the *Content-Type* header to **application/json** (without it the JSON content would be ignored). In a `curl` `POST` request, that header can be provided with the `-H` parameter:  ```bash curl -X POST -H \"Content-Type: application/json\" ... ```  The supplied file must contain a valid JSON: strings need quotes, booleans and integers don't, etc.  The (human-readable) format is:  ```json {   \"key1\": \"value1\",   \"key2\": false,   \"key3\": 42 } ```  Here is an example with inlined data:  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive curl -X POST -H \"X-API-Token: yourToken\" -H  \"Content-Type: application/json\" https://rudder.example.com/rudder/api/rules/latest/{id}   -d '{ \"displayName\": \"new name\", \"enabled\": false, \"directives\": \"directiveId\"}' ```  You can also pass a supply the JSON in a file:  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive curl -X POST -H \"X-API-Token: yourToken\" -H \"Content-Type: application/json\" https://rudder.example.com/rudder/api/rules/latest/{id} -d @jsonParam ```  Note that the general parameters view in the previous chapter cannot be passed in a JSON, and you will need to pass them a URL parameters if you want them to be taken into account (you can't mix JSON and request parameters):  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive with reason message \"Reason used\" curl -X POST -H \"X-API-Token: yourToken\" -H \"Content-Type: application/json\" \"https://rudder.example.com/rudder/api/rules/latest/{id}?reason=Reason used\" -d @jsonParam -d \"reason=Reason ignored\" ```  ##### Request parameters  In some cases, when you have little, simple data to update, JSON can feel bloated. In such cases, you can use request parameters. You will need to pass one parameter for each data you want to change.  Parameters follow the following schema:  ``` key=value ```  You can pass parameters by two means:  * As query parameters: At the end of your url, put a **?** then your first parameter and then a **&** before next parameters. In that case, parameters need to be https://en.wikipedia.org/wiki/Percent-encoding[URL encoded]  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive curl -X POST -H \"X-API-Token: yourToken\"  https://rudder.example.com/rudder/api/rules/latest/{id}?\"displayName=my new name\"&\"enabled=false\"&\"directives=aDirectiveId\" ```  * As request data: You can pass those parameters in the request data, they won't figure in the URL, making it lighter to read, You can pass a file that contains data.  ```bash # Update the Rule 'id' with a new name, disabled, and setting it one directive (in file directive-info.json) curl -X POST -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/rules/latest/{id} -d \"displayName=my new name\" -d \"enabled=false\" -d @directive-info.json ``` 

API version: 18
Contact: dev@rudder.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rudder-golang

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the NodeFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeFull{}

// NodeFull struct for NodeFull
type NodeFull struct {
	// Unique id of the node
	Id string `json:"id"`
	// Fully qualified name of the node
	Hostname string `json:"hostname"`
	// Status of the node
	Status string `json:"status"`
	// Information about CPU architecture (32/64 bits)
	ArchitectureDescription *string `json:"architectureDescription,omitempty"`
	// A human intended description of the node (not used)
	Description *string `json:"description,omitempty"`
	// IP addresses of the node (IPv4 and IPv6)
	IpAddresses []string `json:"ipAddresses"`
	// Date and time of the latest run, if one is available (node time). Up to API v11, format was: \"YYYY-MM-dd HH:mm\"
	LastRunDate *string `json:"lastRunDate,omitempty"`
	// Date and time of the latest generated inventory, if one is available (node time). Up to API v11, format was: \"YYYY-MM-dd HH:mm\"
	LastInventoryDate *string `json:"lastInventoryDate,omitempty"`
	Machine *NodeFullMachine `json:"machine,omitempty"`
	Os *NodeFullOs `json:"os,omitempty"`
	// Management agents running on the node
	ManagementTechnology []NodeFullManagementTechnologyInner `json:"managementTechnology"`
	// Rudder policy server managing the node
	PolicyServerId string `json:"policyServerId"`
	// Node properties (either set by user or filled by third party sources)
	Properties []NodeFullPropertiesInner `json:"properties"`
	// Rudder policy mode for the node (`default` follows the global configuration)
	PolicyMode *string `json:"policyMode,omitempty"`
	// Size of RAM in MB
	Ram *int32 `json:"ram,omitempty"`
	Timezone *NodeFullTimezone `json:"timezone,omitempty"`
	// User accounts declared in the node
	Accounts []string `json:"accounts,omitempty"`
	Bios *NodeFullBios `json:"bios,omitempty"`
	// Physical controller information
	Controllers []NodeFullControllersInner `json:"controllers,omitempty"`
	// Environment variables defined on the node in the context of the agent
	EnvironmentVariables []NodeFullEnvironmentVariablesInner `json:"environmentVariables,omitempty"`
	// File system declared on the node
	FileSystems []NodeFullFileSystemsInner `json:"fileSystems,omitempty"`
	ManagementTechnologyDetails *NodeFullManagementTechnologyDetails `json:"managementTechnologyDetails,omitempty"`
	// Memory slots
	Memories []NodeFullMemoriesInner `json:"memories,omitempty"`
	// Detailed information about registered network interfaces on the node
	NetworkInterfaces []NodeFullNetworkInterfacesInner `json:"networkInterfaces,omitempty"`
	// Physical port information objects
	Ports []NodeFullPortsInner `json:"ports,omitempty"`
	// Process running (at inventory time)
	Processes []NodeFullProcessesInner `json:"processes,omitempty"`
	// CPU information
	Processors []NodeFullProcessorsInner `json:"processors,omitempty"`
	// Physical extension slot information
	Slots []NodeFullSlotsInner `json:"slots,omitempty"`
	// Software installed on the node (can have thousands items)
	Software []NodeFullSoftwareInner `json:"software,omitempty"`
	// Software that can be updated on that machine
	SoftwareUpdate []NodeFullSoftwareUpdateInner `json:"softwareUpdate,omitempty"`
	// Sound card information
	Sound []NodeFullSoundInner `json:"sound,omitempty"`
	// Storage (disks) information objects
	Storage []NodeFullStorageInner `json:"storage,omitempty"`
	// Video card information
	Videos []NodeFullVideosInner `json:"videos,omitempty"`
	// Hosted virtual machine information
	VirtualMachines []NodeFullVirtualMachinesInner `json:"virtualMachines,omitempty"`
}

type _NodeFull NodeFull

// NewNodeFull instantiates a new NodeFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeFull(id string, hostname string, status string, ipAddresses []string, managementTechnology []NodeFullManagementTechnologyInner, policyServerId string, properties []NodeFullPropertiesInner) *NodeFull {
	this := NodeFull{}
	this.Id = id
	this.Hostname = hostname
	this.Status = status
	this.IpAddresses = ipAddresses
	this.ManagementTechnology = managementTechnology
	this.PolicyServerId = policyServerId
	this.Properties = properties
	return &this
}

// NewNodeFullWithDefaults instantiates a new NodeFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeFullWithDefaults() *NodeFull {
	this := NodeFull{}
	return &this
}

// GetId returns the Id field value
func (o *NodeFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NodeFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NodeFull) SetId(v string) {
	o.Id = v
}

// GetHostname returns the Hostname field value
func (o *NodeFull) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *NodeFull) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *NodeFull) SetHostname(v string) {
	o.Hostname = v
}

// GetStatus returns the Status field value
func (o *NodeFull) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NodeFull) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NodeFull) SetStatus(v string) {
	o.Status = v
}

// GetArchitectureDescription returns the ArchitectureDescription field value if set, zero value otherwise.
func (o *NodeFull) GetArchitectureDescription() string {
	if o == nil || IsNil(o.ArchitectureDescription) {
		var ret string
		return ret
	}
	return *o.ArchitectureDescription
}

// GetArchitectureDescriptionOk returns a tuple with the ArchitectureDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetArchitectureDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ArchitectureDescription) {
		return nil, false
	}
	return o.ArchitectureDescription, true
}

// HasArchitectureDescription returns a boolean if a field has been set.
func (o *NodeFull) HasArchitectureDescription() bool {
	if o != nil && !IsNil(o.ArchitectureDescription) {
		return true
	}

	return false
}

// SetArchitectureDescription gets a reference to the given string and assigns it to the ArchitectureDescription field.
func (o *NodeFull) SetArchitectureDescription(v string) {
	o.ArchitectureDescription = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NodeFull) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NodeFull) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NodeFull) SetDescription(v string) {
	o.Description = &v
}

// GetIpAddresses returns the IpAddresses field value
func (o *NodeFull) GetIpAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value
// and a boolean to check if the value has been set.
func (o *NodeFull) GetIpAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddresses, true
}

// SetIpAddresses sets field value
func (o *NodeFull) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetLastRunDate returns the LastRunDate field value if set, zero value otherwise.
func (o *NodeFull) GetLastRunDate() string {
	if o == nil || IsNil(o.LastRunDate) {
		var ret string
		return ret
	}
	return *o.LastRunDate
}

// GetLastRunDateOk returns a tuple with the LastRunDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetLastRunDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastRunDate) {
		return nil, false
	}
	return o.LastRunDate, true
}

// HasLastRunDate returns a boolean if a field has been set.
func (o *NodeFull) HasLastRunDate() bool {
	if o != nil && !IsNil(o.LastRunDate) {
		return true
	}

	return false
}

// SetLastRunDate gets a reference to the given string and assigns it to the LastRunDate field.
func (o *NodeFull) SetLastRunDate(v string) {
	o.LastRunDate = &v
}

// GetLastInventoryDate returns the LastInventoryDate field value if set, zero value otherwise.
func (o *NodeFull) GetLastInventoryDate() string {
	if o == nil || IsNil(o.LastInventoryDate) {
		var ret string
		return ret
	}
	return *o.LastInventoryDate
}

// GetLastInventoryDateOk returns a tuple with the LastInventoryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetLastInventoryDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastInventoryDate) {
		return nil, false
	}
	return o.LastInventoryDate, true
}

// HasLastInventoryDate returns a boolean if a field has been set.
func (o *NodeFull) HasLastInventoryDate() bool {
	if o != nil && !IsNil(o.LastInventoryDate) {
		return true
	}

	return false
}

// SetLastInventoryDate gets a reference to the given string and assigns it to the LastInventoryDate field.
func (o *NodeFull) SetLastInventoryDate(v string) {
	o.LastInventoryDate = &v
}

// GetMachine returns the Machine field value if set, zero value otherwise.
func (o *NodeFull) GetMachine() NodeFullMachine {
	if o == nil || IsNil(o.Machine) {
		var ret NodeFullMachine
		return ret
	}
	return *o.Machine
}

// GetMachineOk returns a tuple with the Machine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetMachineOk() (*NodeFullMachine, bool) {
	if o == nil || IsNil(o.Machine) {
		return nil, false
	}
	return o.Machine, true
}

// HasMachine returns a boolean if a field has been set.
func (o *NodeFull) HasMachine() bool {
	if o != nil && !IsNil(o.Machine) {
		return true
	}

	return false
}

// SetMachine gets a reference to the given NodeFullMachine and assigns it to the Machine field.
func (o *NodeFull) SetMachine(v NodeFullMachine) {
	o.Machine = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *NodeFull) GetOs() NodeFullOs {
	if o == nil || IsNil(o.Os) {
		var ret NodeFullOs
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetOsOk() (*NodeFullOs, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *NodeFull) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given NodeFullOs and assigns it to the Os field.
func (o *NodeFull) SetOs(v NodeFullOs) {
	o.Os = &v
}

// GetManagementTechnology returns the ManagementTechnology field value
func (o *NodeFull) GetManagementTechnology() []NodeFullManagementTechnologyInner {
	if o == nil {
		var ret []NodeFullManagementTechnologyInner
		return ret
	}

	return o.ManagementTechnology
}

// GetManagementTechnologyOk returns a tuple with the ManagementTechnology field value
// and a boolean to check if the value has been set.
func (o *NodeFull) GetManagementTechnologyOk() ([]NodeFullManagementTechnologyInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagementTechnology, true
}

// SetManagementTechnology sets field value
func (o *NodeFull) SetManagementTechnology(v []NodeFullManagementTechnologyInner) {
	o.ManagementTechnology = v
}

// GetPolicyServerId returns the PolicyServerId field value
func (o *NodeFull) GetPolicyServerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyServerId
}

// GetPolicyServerIdOk returns a tuple with the PolicyServerId field value
// and a boolean to check if the value has been set.
func (o *NodeFull) GetPolicyServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyServerId, true
}

// SetPolicyServerId sets field value
func (o *NodeFull) SetPolicyServerId(v string) {
	o.PolicyServerId = v
}

// GetProperties returns the Properties field value
func (o *NodeFull) GetProperties() []NodeFullPropertiesInner {
	if o == nil {
		var ret []NodeFullPropertiesInner
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *NodeFull) GetPropertiesOk() ([]NodeFullPropertiesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *NodeFull) SetProperties(v []NodeFullPropertiesInner) {
	o.Properties = v
}

// GetPolicyMode returns the PolicyMode field value if set, zero value otherwise.
func (o *NodeFull) GetPolicyMode() string {
	if o == nil || IsNil(o.PolicyMode) {
		var ret string
		return ret
	}
	return *o.PolicyMode
}

// GetPolicyModeOk returns a tuple with the PolicyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetPolicyModeOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyMode) {
		return nil, false
	}
	return o.PolicyMode, true
}

// HasPolicyMode returns a boolean if a field has been set.
func (o *NodeFull) HasPolicyMode() bool {
	if o != nil && !IsNil(o.PolicyMode) {
		return true
	}

	return false
}

// SetPolicyMode gets a reference to the given string and assigns it to the PolicyMode field.
func (o *NodeFull) SetPolicyMode(v string) {
	o.PolicyMode = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *NodeFull) GetRam() int32 {
	if o == nil || IsNil(o.Ram) {
		var ret int32
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetRamOk() (*int32, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *NodeFull) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given int32 and assigns it to the Ram field.
func (o *NodeFull) SetRam(v int32) {
	o.Ram = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *NodeFull) GetTimezone() NodeFullTimezone {
	if o == nil || IsNil(o.Timezone) {
		var ret NodeFullTimezone
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetTimezoneOk() (*NodeFullTimezone, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *NodeFull) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given NodeFullTimezone and assigns it to the Timezone field.
func (o *NodeFull) SetTimezone(v NodeFullTimezone) {
	o.Timezone = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *NodeFull) GetAccounts() []string {
	if o == nil || IsNil(o.Accounts) {
		var ret []string
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetAccountsOk() ([]string, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *NodeFull) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []string and assigns it to the Accounts field.
func (o *NodeFull) SetAccounts(v []string) {
	o.Accounts = v
}

// GetBios returns the Bios field value if set, zero value otherwise.
func (o *NodeFull) GetBios() NodeFullBios {
	if o == nil || IsNil(o.Bios) {
		var ret NodeFullBios
		return ret
	}
	return *o.Bios
}

// GetBiosOk returns a tuple with the Bios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetBiosOk() (*NodeFullBios, bool) {
	if o == nil || IsNil(o.Bios) {
		return nil, false
	}
	return o.Bios, true
}

// HasBios returns a boolean if a field has been set.
func (o *NodeFull) HasBios() bool {
	if o != nil && !IsNil(o.Bios) {
		return true
	}

	return false
}

// SetBios gets a reference to the given NodeFullBios and assigns it to the Bios field.
func (o *NodeFull) SetBios(v NodeFullBios) {
	o.Bios = &v
}

// GetControllers returns the Controllers field value if set, zero value otherwise.
func (o *NodeFull) GetControllers() []NodeFullControllersInner {
	if o == nil || IsNil(o.Controllers) {
		var ret []NodeFullControllersInner
		return ret
	}
	return o.Controllers
}

// GetControllersOk returns a tuple with the Controllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetControllersOk() ([]NodeFullControllersInner, bool) {
	if o == nil || IsNil(o.Controllers) {
		return nil, false
	}
	return o.Controllers, true
}

// HasControllers returns a boolean if a field has been set.
func (o *NodeFull) HasControllers() bool {
	if o != nil && !IsNil(o.Controllers) {
		return true
	}

	return false
}

// SetControllers gets a reference to the given []NodeFullControllersInner and assigns it to the Controllers field.
func (o *NodeFull) SetControllers(v []NodeFullControllersInner) {
	o.Controllers = v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *NodeFull) GetEnvironmentVariables() []NodeFullEnvironmentVariablesInner {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret []NodeFullEnvironmentVariablesInner
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetEnvironmentVariablesOk() ([]NodeFullEnvironmentVariablesInner, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *NodeFull) HasEnvironmentVariables() bool {
	if o != nil && !IsNil(o.EnvironmentVariables) {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given []NodeFullEnvironmentVariablesInner and assigns it to the EnvironmentVariables field.
func (o *NodeFull) SetEnvironmentVariables(v []NodeFullEnvironmentVariablesInner) {
	o.EnvironmentVariables = v
}

// GetFileSystems returns the FileSystems field value if set, zero value otherwise.
func (o *NodeFull) GetFileSystems() []NodeFullFileSystemsInner {
	if o == nil || IsNil(o.FileSystems) {
		var ret []NodeFullFileSystemsInner
		return ret
	}
	return o.FileSystems
}

// GetFileSystemsOk returns a tuple with the FileSystems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetFileSystemsOk() ([]NodeFullFileSystemsInner, bool) {
	if o == nil || IsNil(o.FileSystems) {
		return nil, false
	}
	return o.FileSystems, true
}

// HasFileSystems returns a boolean if a field has been set.
func (o *NodeFull) HasFileSystems() bool {
	if o != nil && !IsNil(o.FileSystems) {
		return true
	}

	return false
}

// SetFileSystems gets a reference to the given []NodeFullFileSystemsInner and assigns it to the FileSystems field.
func (o *NodeFull) SetFileSystems(v []NodeFullFileSystemsInner) {
	o.FileSystems = v
}

// GetManagementTechnologyDetails returns the ManagementTechnologyDetails field value if set, zero value otherwise.
func (o *NodeFull) GetManagementTechnologyDetails() NodeFullManagementTechnologyDetails {
	if o == nil || IsNil(o.ManagementTechnologyDetails) {
		var ret NodeFullManagementTechnologyDetails
		return ret
	}
	return *o.ManagementTechnologyDetails
}

// GetManagementTechnologyDetailsOk returns a tuple with the ManagementTechnologyDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetManagementTechnologyDetailsOk() (*NodeFullManagementTechnologyDetails, bool) {
	if o == nil || IsNil(o.ManagementTechnologyDetails) {
		return nil, false
	}
	return o.ManagementTechnologyDetails, true
}

// HasManagementTechnologyDetails returns a boolean if a field has been set.
func (o *NodeFull) HasManagementTechnologyDetails() bool {
	if o != nil && !IsNil(o.ManagementTechnologyDetails) {
		return true
	}

	return false
}

// SetManagementTechnologyDetails gets a reference to the given NodeFullManagementTechnologyDetails and assigns it to the ManagementTechnologyDetails field.
func (o *NodeFull) SetManagementTechnologyDetails(v NodeFullManagementTechnologyDetails) {
	o.ManagementTechnologyDetails = &v
}

// GetMemories returns the Memories field value if set, zero value otherwise.
func (o *NodeFull) GetMemories() []NodeFullMemoriesInner {
	if o == nil || IsNil(o.Memories) {
		var ret []NodeFullMemoriesInner
		return ret
	}
	return o.Memories
}

// GetMemoriesOk returns a tuple with the Memories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetMemoriesOk() ([]NodeFullMemoriesInner, bool) {
	if o == nil || IsNil(o.Memories) {
		return nil, false
	}
	return o.Memories, true
}

// HasMemories returns a boolean if a field has been set.
func (o *NodeFull) HasMemories() bool {
	if o != nil && !IsNil(o.Memories) {
		return true
	}

	return false
}

// SetMemories gets a reference to the given []NodeFullMemoriesInner and assigns it to the Memories field.
func (o *NodeFull) SetMemories(v []NodeFullMemoriesInner) {
	o.Memories = v
}

// GetNetworkInterfaces returns the NetworkInterfaces field value if set, zero value otherwise.
func (o *NodeFull) GetNetworkInterfaces() []NodeFullNetworkInterfacesInner {
	if o == nil || IsNil(o.NetworkInterfaces) {
		var ret []NodeFullNetworkInterfacesInner
		return ret
	}
	return o.NetworkInterfaces
}

// GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetNetworkInterfacesOk() ([]NodeFullNetworkInterfacesInner, bool) {
	if o == nil || IsNil(o.NetworkInterfaces) {
		return nil, false
	}
	return o.NetworkInterfaces, true
}

// HasNetworkInterfaces returns a boolean if a field has been set.
func (o *NodeFull) HasNetworkInterfaces() bool {
	if o != nil && !IsNil(o.NetworkInterfaces) {
		return true
	}

	return false
}

// SetNetworkInterfaces gets a reference to the given []NodeFullNetworkInterfacesInner and assigns it to the NetworkInterfaces field.
func (o *NodeFull) SetNetworkInterfaces(v []NodeFullNetworkInterfacesInner) {
	o.NetworkInterfaces = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *NodeFull) GetPorts() []NodeFullPortsInner {
	if o == nil || IsNil(o.Ports) {
		var ret []NodeFullPortsInner
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetPortsOk() ([]NodeFullPortsInner, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *NodeFull) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []NodeFullPortsInner and assigns it to the Ports field.
func (o *NodeFull) SetPorts(v []NodeFullPortsInner) {
	o.Ports = v
}

// GetProcesses returns the Processes field value if set, zero value otherwise.
func (o *NodeFull) GetProcesses() []NodeFullProcessesInner {
	if o == nil || IsNil(o.Processes) {
		var ret []NodeFullProcessesInner
		return ret
	}
	return o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetProcessesOk() ([]NodeFullProcessesInner, bool) {
	if o == nil || IsNil(o.Processes) {
		return nil, false
	}
	return o.Processes, true
}

// HasProcesses returns a boolean if a field has been set.
func (o *NodeFull) HasProcesses() bool {
	if o != nil && !IsNil(o.Processes) {
		return true
	}

	return false
}

// SetProcesses gets a reference to the given []NodeFullProcessesInner and assigns it to the Processes field.
func (o *NodeFull) SetProcesses(v []NodeFullProcessesInner) {
	o.Processes = v
}

// GetProcessors returns the Processors field value if set, zero value otherwise.
func (o *NodeFull) GetProcessors() []NodeFullProcessorsInner {
	if o == nil || IsNil(o.Processors) {
		var ret []NodeFullProcessorsInner
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetProcessorsOk() ([]NodeFullProcessorsInner, bool) {
	if o == nil || IsNil(o.Processors) {
		return nil, false
	}
	return o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *NodeFull) HasProcessors() bool {
	if o != nil && !IsNil(o.Processors) {
		return true
	}

	return false
}

// SetProcessors gets a reference to the given []NodeFullProcessorsInner and assigns it to the Processors field.
func (o *NodeFull) SetProcessors(v []NodeFullProcessorsInner) {
	o.Processors = v
}

// GetSlots returns the Slots field value if set, zero value otherwise.
func (o *NodeFull) GetSlots() []NodeFullSlotsInner {
	if o == nil || IsNil(o.Slots) {
		var ret []NodeFullSlotsInner
		return ret
	}
	return o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetSlotsOk() ([]NodeFullSlotsInner, bool) {
	if o == nil || IsNil(o.Slots) {
		return nil, false
	}
	return o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *NodeFull) HasSlots() bool {
	if o != nil && !IsNil(o.Slots) {
		return true
	}

	return false
}

// SetSlots gets a reference to the given []NodeFullSlotsInner and assigns it to the Slots field.
func (o *NodeFull) SetSlots(v []NodeFullSlotsInner) {
	o.Slots = v
}

// GetSoftware returns the Software field value if set, zero value otherwise.
func (o *NodeFull) GetSoftware() []NodeFullSoftwareInner {
	if o == nil || IsNil(o.Software) {
		var ret []NodeFullSoftwareInner
		return ret
	}
	return o.Software
}

// GetSoftwareOk returns a tuple with the Software field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetSoftwareOk() ([]NodeFullSoftwareInner, bool) {
	if o == nil || IsNil(o.Software) {
		return nil, false
	}
	return o.Software, true
}

// HasSoftware returns a boolean if a field has been set.
func (o *NodeFull) HasSoftware() bool {
	if o != nil && !IsNil(o.Software) {
		return true
	}

	return false
}

// SetSoftware gets a reference to the given []NodeFullSoftwareInner and assigns it to the Software field.
func (o *NodeFull) SetSoftware(v []NodeFullSoftwareInner) {
	o.Software = v
}

// GetSoftwareUpdate returns the SoftwareUpdate field value if set, zero value otherwise.
func (o *NodeFull) GetSoftwareUpdate() []NodeFullSoftwareUpdateInner {
	if o == nil || IsNil(o.SoftwareUpdate) {
		var ret []NodeFullSoftwareUpdateInner
		return ret
	}
	return o.SoftwareUpdate
}

// GetSoftwareUpdateOk returns a tuple with the SoftwareUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetSoftwareUpdateOk() ([]NodeFullSoftwareUpdateInner, bool) {
	if o == nil || IsNil(o.SoftwareUpdate) {
		return nil, false
	}
	return o.SoftwareUpdate, true
}

// HasSoftwareUpdate returns a boolean if a field has been set.
func (o *NodeFull) HasSoftwareUpdate() bool {
	if o != nil && !IsNil(o.SoftwareUpdate) {
		return true
	}

	return false
}

// SetSoftwareUpdate gets a reference to the given []NodeFullSoftwareUpdateInner and assigns it to the SoftwareUpdate field.
func (o *NodeFull) SetSoftwareUpdate(v []NodeFullSoftwareUpdateInner) {
	o.SoftwareUpdate = v
}

// GetSound returns the Sound field value if set, zero value otherwise.
func (o *NodeFull) GetSound() []NodeFullSoundInner {
	if o == nil || IsNil(o.Sound) {
		var ret []NodeFullSoundInner
		return ret
	}
	return o.Sound
}

// GetSoundOk returns a tuple with the Sound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetSoundOk() ([]NodeFullSoundInner, bool) {
	if o == nil || IsNil(o.Sound) {
		return nil, false
	}
	return o.Sound, true
}

// HasSound returns a boolean if a field has been set.
func (o *NodeFull) HasSound() bool {
	if o != nil && !IsNil(o.Sound) {
		return true
	}

	return false
}

// SetSound gets a reference to the given []NodeFullSoundInner and assigns it to the Sound field.
func (o *NodeFull) SetSound(v []NodeFullSoundInner) {
	o.Sound = v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *NodeFull) GetStorage() []NodeFullStorageInner {
	if o == nil || IsNil(o.Storage) {
		var ret []NodeFullStorageInner
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetStorageOk() ([]NodeFullStorageInner, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *NodeFull) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []NodeFullStorageInner and assigns it to the Storage field.
func (o *NodeFull) SetStorage(v []NodeFullStorageInner) {
	o.Storage = v
}

// GetVideos returns the Videos field value if set, zero value otherwise.
func (o *NodeFull) GetVideos() []NodeFullVideosInner {
	if o == nil || IsNil(o.Videos) {
		var ret []NodeFullVideosInner
		return ret
	}
	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetVideosOk() ([]NodeFullVideosInner, bool) {
	if o == nil || IsNil(o.Videos) {
		return nil, false
	}
	return o.Videos, true
}

// HasVideos returns a boolean if a field has been set.
func (o *NodeFull) HasVideos() bool {
	if o != nil && !IsNil(o.Videos) {
		return true
	}

	return false
}

// SetVideos gets a reference to the given []NodeFullVideosInner and assigns it to the Videos field.
func (o *NodeFull) SetVideos(v []NodeFullVideosInner) {
	o.Videos = v
}

// GetVirtualMachines returns the VirtualMachines field value if set, zero value otherwise.
func (o *NodeFull) GetVirtualMachines() []NodeFullVirtualMachinesInner {
	if o == nil || IsNil(o.VirtualMachines) {
		var ret []NodeFullVirtualMachinesInner
		return ret
	}
	return o.VirtualMachines
}

// GetVirtualMachinesOk returns a tuple with the VirtualMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFull) GetVirtualMachinesOk() ([]NodeFullVirtualMachinesInner, bool) {
	if o == nil || IsNil(o.VirtualMachines) {
		return nil, false
	}
	return o.VirtualMachines, true
}

// HasVirtualMachines returns a boolean if a field has been set.
func (o *NodeFull) HasVirtualMachines() bool {
	if o != nil && !IsNil(o.VirtualMachines) {
		return true
	}

	return false
}

// SetVirtualMachines gets a reference to the given []NodeFullVirtualMachinesInner and assigns it to the VirtualMachines field.
func (o *NodeFull) SetVirtualMachines(v []NodeFullVirtualMachinesInner) {
	o.VirtualMachines = v
}

func (o NodeFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["hostname"] = o.Hostname
	toSerialize["status"] = o.Status
	if !IsNil(o.ArchitectureDescription) {
		toSerialize["architectureDescription"] = o.ArchitectureDescription
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["ipAddresses"] = o.IpAddresses
	if !IsNil(o.LastRunDate) {
		toSerialize["lastRunDate"] = o.LastRunDate
	}
	if !IsNil(o.LastInventoryDate) {
		toSerialize["lastInventoryDate"] = o.LastInventoryDate
	}
	if !IsNil(o.Machine) {
		toSerialize["machine"] = o.Machine
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	toSerialize["managementTechnology"] = o.ManagementTechnology
	toSerialize["policyServerId"] = o.PolicyServerId
	toSerialize["properties"] = o.Properties
	if !IsNil(o.PolicyMode) {
		toSerialize["policyMode"] = o.PolicyMode
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Bios) {
		toSerialize["bios"] = o.Bios
	}
	if !IsNil(o.Controllers) {
		toSerialize["controllers"] = o.Controllers
	}
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	if !IsNil(o.FileSystems) {
		toSerialize["fileSystems"] = o.FileSystems
	}
	if !IsNil(o.ManagementTechnologyDetails) {
		toSerialize["managementTechnologyDetails"] = o.ManagementTechnologyDetails
	}
	if !IsNil(o.Memories) {
		toSerialize["memories"] = o.Memories
	}
	if !IsNil(o.NetworkInterfaces) {
		toSerialize["networkInterfaces"] = o.NetworkInterfaces
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Processes) {
		toSerialize["processes"] = o.Processes
	}
	if !IsNil(o.Processors) {
		toSerialize["processors"] = o.Processors
	}
	if !IsNil(o.Slots) {
		toSerialize["slots"] = o.Slots
	}
	if !IsNil(o.Software) {
		toSerialize["software"] = o.Software
	}
	if !IsNil(o.SoftwareUpdate) {
		toSerialize["softwareUpdate"] = o.SoftwareUpdate
	}
	if !IsNil(o.Sound) {
		toSerialize["sound"] = o.Sound
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !IsNil(o.Videos) {
		toSerialize["videos"] = o.Videos
	}
	if !IsNil(o.VirtualMachines) {
		toSerialize["virtualMachines"] = o.VirtualMachines
	}
	return toSerialize, nil
}

func (o *NodeFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"hostname",
		"status",
		"ipAddresses",
		"managementTechnology",
		"policyServerId",
		"properties",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNodeFull := _NodeFull{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNodeFull)

	if err != nil {
		return err
	}

	*o = NodeFull(varNodeFull)

	return err
}

type NullableNodeFull struct {
	value *NodeFull
	isSet bool
}

func (v NullableNodeFull) Get() *NodeFull {
	return v.value
}

func (v *NullableNodeFull) Set(val *NodeFull) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeFull) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeFull(val *NodeFull) *NullableNodeFull {
	return &NullableNodeFull{value: val, isSet: true}
}

func (v NullableNodeFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


