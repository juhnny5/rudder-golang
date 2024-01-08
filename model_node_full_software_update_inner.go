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
)

// checks if the NodeFullSoftwareUpdateInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeFullSoftwareUpdateInner{}

// NodeFullSoftwareUpdateInner struct for NodeFullSoftwareUpdateInner
type NodeFullSoftwareUpdateInner struct {
	// name of software that can be updated
	Name *string `json:"name,omitempty"`
	// available version for software
	Version *string `json:"version,omitempty"`
	// CPU architecture of the update
	Arch *string `json:"arch,omitempty"`
	// tool that discovered that update
	From *string `json:"from,omitempty"`
	// if available, kind of patch provided by that update, else none
	Kind *string `json:"kind,omitempty"`
	// information about the source providing that update
	Source *string `json:"source,omitempty"`
	// details about the content of the update, if available
	Description *string `json:"description,omitempty"`
	// if available, the severity of the update
	Severity *string `json:"severity,omitempty"`
	Ids []string `json:"ids,omitempty"`
}

// NewNodeFullSoftwareUpdateInner instantiates a new NodeFullSoftwareUpdateInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeFullSoftwareUpdateInner() *NodeFullSoftwareUpdateInner {
	this := NodeFullSoftwareUpdateInner{}
	return &this
}

// NewNodeFullSoftwareUpdateInnerWithDefaults instantiates a new NodeFullSoftwareUpdateInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeFullSoftwareUpdateInnerWithDefaults() *NodeFullSoftwareUpdateInner {
	this := NodeFullSoftwareUpdateInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NodeFullSoftwareUpdateInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFullSoftwareUpdateInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NodeFullSoftwareUpdateInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NodeFullSoftwareUpdateInner) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NodeFullSoftwareUpdateInner) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFullSoftwareUpdateInner) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NodeFullSoftwareUpdateInner) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NodeFullSoftwareUpdateInner) SetVersion(v string) {
	o.Version = &v
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *NodeFullSoftwareUpdateInner) GetArch() string {
	if o == nil || IsNil(o.Arch) {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFullSoftwareUpdateInner) GetArchOk() (*string, bool) {
	if o == nil || IsNil(o.Arch) {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *NodeFullSoftwareUpdateInner) HasArch() bool {
	if o != nil && !IsNil(o.Arch) {
		return true
	}

	return false
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *NodeFullSoftwareUpdateInner) SetArch(v string) {
	o.Arch = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *NodeFullSoftwareUpdateInner) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFullSoftwareUpdateInner) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *NodeFullSoftwareUpdateInner) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *NodeFullSoftwareUpdateInner) SetFrom(v string) {
	o.From = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *NodeFullSoftwareUpdateInner) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFullSoftwareUpdateInner) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *NodeFullSoftwareUpdateInner) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *NodeFullSoftwareUpdateInner) SetKind(v string) {
	o.Kind = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *NodeFullSoftwareUpdateInner) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFullSoftwareUpdateInner) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *NodeFullSoftwareUpdateInner) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *NodeFullSoftwareUpdateInner) SetSource(v string) {
	o.Source = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NodeFullSoftwareUpdateInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFullSoftwareUpdateInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NodeFullSoftwareUpdateInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NodeFullSoftwareUpdateInner) SetDescription(v string) {
	o.Description = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *NodeFullSoftwareUpdateInner) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFullSoftwareUpdateInner) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *NodeFullSoftwareUpdateInner) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *NodeFullSoftwareUpdateInner) SetSeverity(v string) {
	o.Severity = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *NodeFullSoftwareUpdateInner) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeFullSoftwareUpdateInner) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *NodeFullSoftwareUpdateInner) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *NodeFullSoftwareUpdateInner) SetIds(v []string) {
	o.Ids = v
}

func (o NodeFullSoftwareUpdateInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeFullSoftwareUpdateInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Arch) {
		toSerialize["arch"] = o.Arch
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableNodeFullSoftwareUpdateInner struct {
	value *NodeFullSoftwareUpdateInner
	isSet bool
}

func (v NullableNodeFullSoftwareUpdateInner) Get() *NodeFullSoftwareUpdateInner {
	return v.value
}

func (v *NullableNodeFullSoftwareUpdateInner) Set(val *NodeFullSoftwareUpdateInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeFullSoftwareUpdateInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeFullSoftwareUpdateInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeFullSoftwareUpdateInner(val *NodeFullSoftwareUpdateInner) *NullableNodeFullSoftwareUpdateInner {
	return &NullableNodeFullSoftwareUpdateInner{value: val, isSet: true}
}

func (v NullableNodeFullSoftwareUpdateInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeFullSoftwareUpdateInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


