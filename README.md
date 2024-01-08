# Go API client for rudder-golang

Download OpenAPI specification: [openapi.yml](openapi.yml)

**Other documentation sources**:

* [Main documentation](https://docs.rudder.io)
* [Internal relay API](https://docs.rudder.io/api/relay/)

# Introduction

Rudder exposes a REST API, enabling the user to interact with Rudder without using the webapp, for example, in scripts or cron jobs.

## Authentication

The Rudder REST API uses simple API keys for authentication.
All requests must be authenticated (except from a generic status API).
The tokens are 32-character strings, passed in a `X-API-Token` header, like in: 

```bash
curl --header \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules
```

The tokens are the API equivalent of a password, and must
be secured just like a password would be.

### API accounts

The accounts are managed in the Web interface. There are two possible types of accounts:

* **Global API accounts**: they are not linked to a Rudder user, and are managed by Rudder administrators in the _Administration -> API accounts_ page. You should define an expiration date whenever possible.

![General API tokens settings](assets/api-tokens.png \"General API tokens settings\")

* **User tokens**: they are linked to a Rudder user, and give the same rights the user has.
There can be only one token by user. This feature is provided by the `api-authorizatons` plugin.

![User API token](assets/api-user.png \"User API token\")

When an action produces a change of configuration on the server, the API account that made it will
be recorded in the event log, like for a Web interaction.

### Authorization

When using Rudder without the `api-authorizatons` plugin, only global accounts are available, with
two possible privilege levels, read-only or write.
With the `api-authorizatons` plugin, you also get access to:

* User tokens, which have the same permissions as the user, using the Rudder roles and permissions feature.
* Custom ACLs on global API accounts. They provide fine-grained permissions on every endpoint:

![Custom API ACL](assets/custom-acl.png \"Custom API ACL\")

As a general principle,
you should create dedicated tokens with the least privilege level for each different interaction you have with the
API.
This limits the risks of exploitation if a token is stolen, and allows tracking the activity
of each token separately. Token renewal is also easier when they are only used for a limited purpose.

## Versioning

Each time the API is extended with new features (new functions, new parameters, new responses, ...), it will be assigned a new version number. This will allow you
to keep your existing scripts (based on previous behavior). Versions will always be integers (no 2.1 or 3.3, just 2, 3, 4, ...) or `latest`.

You can change the version of the API used by setting it either within the url or in a header:

* the URL: each URL is prefixed by its version id, like `/api/version/function`.

```bash
# Version 10
curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/10/rules
# Latest
curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules
# Wrong (not an integer) => 404 not found
curl -X GET -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/3.14/rules
```

* the HTTP headers. You can add the **X-API-Version** header to your request. The value needs to be an integer or `latest`.

```bash
# Version 10
curl -X GET -H \"X-API-Token: yourToken\" -H \"X-API-Version: 10\" https://rudder.example.com/rudder/api/rules
# Wrong => Error response indicating which versions are available
curl -X GET -H \"X-API-Token: yourToken\" -H \"X-API-Version: 3.14\" https://rudder.example.com/rudder/api/rules
```

In the future, we may declare some versions as deprecated, in order to remove them in a later version of Rudder, but we will never remove any versions without warning, or without a safe
period of time to allow migration from previous versions.


<h4>Existing versions</h4>
<table>
  <thead>
    <tr>
      <th style=\"width: 20%\">Version</th>
      <th style=\"width: 20%\">Rudder versions it appeared in</th>
      <th style=\"width: 70%\">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td class=\"code\">1</td>
      <td class=\"code\">Never released (for internal use only)</td>
      <td>Experimental version</td>
    </tr>
    <tr>
      <td class=\"code\">2 to 10 (deprecated)</td>
      <td class=\"code\">4.3 and before</td>
      <td>These versions provided the core set of API features for rules, directives, nodes global parameters, change requests and compliance, rudder settings, and system API</td>
    </tr>
    <tr>
      <td class=\"code\">11</td>
      <td class=\"code\">5.0</td>
      <td>New system API (replacing old localhost v1 api): status, maintenance operations and server behavior</td>
    </tr>
    <tr>
      <td class=\"code\">12</td>
      <td class=\"code\">6.0 and 6.1</td>
      <td>Node key management</td>
    </tr>
    <tr>
      <td class=\"code\">13</td>
      <td class=\"code\">6.2</td>
      <td><ul>
        <li>Node status endpoint</li>
        <li>System health check</li>
        <li>System maintenance job to purge software [that endpoint was back-ported in 6.1]</li>
      </ul></td>
    </tr>
    <tr>
      <td class=\"code\">14</td>
      <td class=\"code\">7.0</td>
      <td><ul>
        <li>Secret management</li>
        <li>Directive tree</li>
        <li>Improve techniques management</li>
        <li>Demote a relay</li>
      </ul></td>
    </tr>
    <tr>
      <td class=\"code\">15</td>
      <td class=\"code\">7.1</td>
      <td><ul>
        <li>Package updates in nodes</li>
      </ul></td>
    </tr>
    <tr>
      <td class=\"code\">16</td>
      <td class=\"code\">7.2</td>
      <td><ul>
        <li>Create node API included from plugin</li>
        <li>Configuration archive import/export</li>
      </ul></td>
    </tr>
    <tr>
      <td class=\"code\">17</td>
      <td class=\"code\">7.3</td>
      <td><ul>
        <li>Compliance by directive</li>
        <li>Path campaigns API included</li>
      </ul></td>
    </tr>
    <tr>
      <td class=\"code\">18</td>
      <td class=\"code\">8.0</td>
      <td><ul>
        <li>Allowed network </li>
        <li>Improve the structure of `/settings/allowed_networks` output</li>
      </ul></td>
    </tr>
  </tbody>
</table>


## Response format

All responses from the API are in the JSON format.

```json
{
  \"action\": \"The name of the called function\",
  \"id\": \"The ID of the element you want, if relevant\",
  \"result\": \"The result of your action: success or error\",
  \"data\": \"Only present if this is a success and depends on the function, it's usually a JSON object\",
  \"errorDetails\": \"Only present if this is an error, it contains the error message\"
}
```


* __Success__ responses are sent with the 200 HTTP (Success) code

* __Error__ responses are sent with a HTTP error code (mostly 5xx...)


## HTTP method

Rudder's REST API is based on the usage of [HTTP methods](http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html). We use them to indicate what action will be done by the request. Currently, we use four of them:


* **GET**: search or retrieve information (get rule details, get a group, ...)

* **PUT**: add new objects (create a directive, clone a Rule, ...)

* **DELETE**: remove objects (delete a node, delete a parameter, ...)

* **POST**: update existing objects (update a directive, reload a group, ...)


## Parameters

### General parameters

Some parameters are available for almost all API functions. They will be described in this section.
They must be part of the query and can't be submitted in a JSON form.

#### Available for all requests

<table>
  <thead>
    <tr>
      <th style=\"width: 30%\">Field</th>
      <th style=\"width: 10%\">Type</th>
      <th style=\"width: 70%\">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td class=\"code\">prettify</td>
      <td><b>boolean</b><br><i>optional</i></td>
      <td>
        Determine if the answer should be prettified (human friendly) or not. We recommend using this for debugging purposes, but not for general script usage as this does add some unnecessary load on the server side.
        <p class=\"default-value\">Default value: <code>false</code></p>
      </td>
    </tr>
  </tbody>
</table>


#### Available for modification requests (PUT/POST/DELETE)

<table>
  <thead>
    <tr>
      <th style=\"width: 25%\">Field</th>
      <th style=\"width: 12%\">Type</th>
      <th style=\"width: 70%\">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td class=\"code\">reason</td>
      <td><b>string</b><br><i>optional</i> or <i>required</i></td>
      <td>
        Set a message to explain the change. If you set the reason messages to be mandatory in the web interface, failing to supply this value will lead to an error.
        <p class=\"default-value\">Default value: <code>\"\"</code></p>
      </td>
    </tr>
    <tr>
      <td class=\"code\">changeRequestName</td>
      <td><b>string</b><br><i>optional</i></td>
      <td>
        Set the change request name, is used only if workflows are enabled. The default value depends on the function called
        <p class=\"default-value\">Default value: <code>A default string for each function</code></p>
      </td>
    </tr>
    <tr>
      <td class=\"code\">changeRequestDescription</td>
      <td><b>string</b><br><i>optional</i></td>
      <td>
        Set the change request description, is used only if workflows are enabled.
        <p class=\"default-value\">Default value: <code>\"\"</code></p>
      </td>
    </tr>
  </tbody>
</table>


### Passing parameters

Parameters to the API can be sent:

* As part of the URL for resource identification

* As data for POST/PUT requests

  * Directly in JSON format

  * As request arguments

#### As part of the URL for resource identification

Parameters in URLs are used to indicate which resource you want to interact with. The function will not work if this resource is missing.

```bash
# Get the Rule of ID \"id\"
curl -H \"X-API-Token: yourToken\" https://rudder.example.com/rudder/api/latest/rules/id
```

CAUTION: To avoid surprising behavior, do not put a '/' at the end of a URL: it would be interpreted as '/[empty string parameter]' and redirected to '/index', likely not what you wanted to do.


#### Sending data for POST/PUT requests

##### Directly in JSON format

JSON format is the preferred way to interact with Rudder API for creating or updating resources.
You'll also have to set the *Content-Type* header to **application/json** (without it the JSON content would be ignored).
In a `curl` `POST` request, that header can be provided with the `-H` parameter:

```bash
curl -X POST -H \"Content-Type: application/json\" ...
```

The supplied file must contain a valid JSON: strings need quotes, booleans and integers don't, etc.

The (human-readable) format is:

```json
{
  \"key1\": \"value1\",
  \"key2\": false,
  \"key3\": 42
}
```

Here is an example with inlined data:

```bash
# Update the Rule 'id' with a new name, disabled, and setting it one directive
curl -X POST -H \"X-API-Token: yourToken\" -H  \"Content-Type: application/json\"
https://rudder.example.com/rudder/api/rules/latest/{id}
  -d '{ \"displayName\": \"new name\", \"enabled\": false, \"directives\": \"directiveId\"}'
```

You can also pass a supply the JSON in a file:

```bash
# Update the Rule 'id' with a new name, disabled, and setting it one directive
curl -X POST -H \"X-API-Token: yourToken\" -H \"Content-Type: application/json\" https://rudder.example.com/rudder/api/rules/latest/{id} -d @jsonParam
```

Note that the general parameters view in the previous chapter cannot be passed in a JSON, and you will need to pass them a URL parameters if you want them to be taken into account (you can't mix JSON and request parameters):

```bash
# Update the Rule 'id' with a new name, disabled, and setting it one directive with reason message \"Reason used\"
curl -X POST -H \"X-API-Token: yourToken\" -H \"Content-Type: application/json\" \"https://rudder.example.com/rudder/api/rules/latest/{id}?reason=Reason used\" -d @jsonParam -d \"reason=Reason ignored\"
```

##### Request parameters

In some cases, when you have little, simple data to update, JSON can feel bloated. In such cases, you can use
request parameters. You will need to pass one parameter for each data you want to change.

Parameters follow the following schema:

```
key=value
```

You can pass parameters by two means:

* As query parameters: At the end of your url, put a **?** then your first parameter and then a **&** before next parameters. In that case, parameters need to be https://en.wikipedia.org/wiki/Percent-encoding[URL encoded]

```bash
# Update the Rule 'id' with a new name, disabled, and setting it one directive
curl -X POST -H \"X-API-Token: yourToken\"  https://rudder.example.com/rudder/api/rules/latest/{id}?\"displayName=my new name\"&\"enabled=false\"&\"directives=aDirectiveId\"
```

* As request data: You can pass those parameters in the request data, they won't figure in the URL, making it lighter to read, You can pass a file that contains data.

```bash
# Update the Rule 'id' with a new name, disabled, and setting it one directive (in file directive-info.json)
curl -X POST -H \"X-API-Token: yourToken\"
https://rudder.example.com/rudder/api/rules/latest/{id} -d \"displayName=my new name\" -d \"enabled=false\" -d @directive-info.json
```


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 18
- Package version: 0.0.1
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.rudder.io](https://www.rudder.io)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import rudder-golang "github.com/juhnny5/rudder-golang"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `rudder-golang.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), rudder-golang.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `rudder-golang.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), rudder-golang.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `rudder-golang.ContextOperationServerIndices` and `rudder-golang.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), rudder-golang.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), rudder-golang.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://rudder.example.local/rudder/api/latest*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*APIInfoAPI* | [**ApiGeneralInformations**](docs/APIInfoAPI.md#apigeneralinformations) | **Get** /info | List all endpoints
*APIInfoAPI* | [**ApiInformations**](docs/APIInfoAPI.md#apiinformations) | **Get** /info/details/{endpointName} | Get information about one API endpoint
*APIInfoAPI* | [**ApiSubInformations**](docs/APIInfoAPI.md#apisubinformations) | **Get** /info/{sectionId} | Get information on endpoint in a section
*ArchivesAPI* | [**CallImport**](docs/ArchivesAPI.md#callimport) | **Post** /archives/import | Import a ZIP archive of policies into Rudder
*ArchivesAPI* | [**Export**](docs/ArchivesAPI.md#export) | **Get** /archives/export | Get a ZIP archive of the requested items and their dependencies
*BrandingAPI* | [**GetBrandingConf**](docs/BrandingAPI.md#getbrandingconf) | **Get** /branding | Get branding configuration
*BrandingAPI* | [**ReloadBrandingConf**](docs/BrandingAPI.md#reloadbrandingconf) | **Post** /branding/reload | Reload branding file
*BrandingAPI* | [**UpdateBRandingConf**](docs/BrandingAPI.md#updatebrandingconf) | **Post** /branding | Update web interface customization
*CVEAPI* | [**CheckCVE**](docs/CVEAPI.md#checkcve) | **Post** /cve/check | Trigger a CVE check
*CVEAPI* | [**GetAllCve**](docs/CVEAPI.md#getallcve) | **Get** /cve | Get all CVE details
*CVEAPI* | [**GetCVECheckConfiguration**](docs/CVEAPI.md#getcvecheckconfiguration) | **Get** /cve/check/config | Get CVE check config
*CVEAPI* | [**GetCVEList**](docs/CVEAPI.md#getcvelist) | **Post** /cve/list | Get a list of CVE details
*CVEAPI* | [**GetCve**](docs/CVEAPI.md#getcve) | **Get** /cve/{cveId} | Get a CVE details
*CVEAPI* | [**GetLastCVECheck**](docs/CVEAPI.md#getlastcvecheck) | **Get** /cve/check/last | Get last CVE check result
*CVEAPI* | [**ReadCVEfromFS**](docs/CVEAPI.md#readcvefromfs) | **Post** /cve/update/fs | Update CVE database from file system
*CVEAPI* | [**UpdateCVE**](docs/CVEAPI.md#updatecve) | **Post** /cve/update | Update CVE database from remote source
*CVEAPI* | [**UpdateCVECheckConfiguration**](docs/CVEAPI.md#updatecvecheckconfiguration) | **Post** /cve/check/config | Update cve check config
*CampaignsAPI* | [**AllCampaigns**](docs/CampaignsAPI.md#allcampaigns) | **Get** /campaigns | Get all campaigns details
*CampaignsAPI* | [**DeleteCampaign**](docs/CampaignsAPI.md#deletecampaign) | **Delete** /campaigns/{id} | Delete a campaign
*CampaignsAPI* | [**DeleteCampaignEvent**](docs/CampaignsAPI.md#deletecampaignevent) | **Delete** /campaigns/events/{id} | Delete a campaign event details
*CampaignsAPI* | [**GetAllCampaignEvents**](docs/CampaignsAPI.md#getallcampaignevents) | **Get** /campaigns/events | Get all campaign events
*CampaignsAPI* | [**GetCampaign**](docs/CampaignsAPI.md#getcampaign) | **Get** /campaigns/{id} | Get a campaign details
*CampaignsAPI* | [**GetCampaignEvent**](docs/CampaignsAPI.md#getcampaignevent) | **Get** /campaigns/events/{id} | Get a campaign event details
*CampaignsAPI* | [**GetEventsCampaign**](docs/CampaignsAPI.md#geteventscampaign) | **Get** /campaigns/{id}/events | Get campaign events for a campaign
*CampaignsAPI* | [**SaveCampaign**](docs/CampaignsAPI.md#savecampaign) | **Post** /campaigns | Save a campaign
*CampaignsAPI* | [**SaveCampaignEvent**](docs/CampaignsAPI.md#savecampaignevent) | **Post** /campaigns/events/{id} | Update an existing event
*CampaignsAPI* | [**ScheduleCampaign**](docs/CampaignsAPI.md#schedulecampaign) | **Post** /campaigns/{id}/schedule | Schedule a campaign event for a campaign
*ChangeRequestsAPI* | [**AcceptChangeRequest**](docs/ChangeRequestsAPI.md#acceptchangerequest) | **Post** /changeRequests/{changeRequestId}/accept | Accept a request details
*ChangeRequestsAPI* | [**ChangeRequestDetails**](docs/ChangeRequestsAPI.md#changerequestdetails) | **Get** /changeRequests/{changeRequestId} | Get a change request details
*ChangeRequestsAPI* | [**DeclineChangeRequest**](docs/ChangeRequestsAPI.md#declinechangerequest) | **Delete** /changeRequests/{changeRequestId} | Decline a request details
*ChangeRequestsAPI* | [**ListChangeRequests**](docs/ChangeRequestsAPI.md#listchangerequests) | **Get** /api/changeRequests | List all change requests
*ChangeRequestsAPI* | [**ListUsers**](docs/ChangeRequestsAPI.md#listusers) | **Get** /users | List user
*ChangeRequestsAPI* | [**RemoveValidatedUser**](docs/ChangeRequestsAPI.md#removevalidateduser) | **Delete** /validatedUsers/{username} | Remove an user from validated user list
*ChangeRequestsAPI* | [**SaveWorkflowUser**](docs/ChangeRequestsAPI.md#saveworkflowuser) | **Post** /validatedUsers | Update validated user list
*ChangeRequestsAPI* | [**UpdateChangeRequest**](docs/ChangeRequestsAPI.md#updatechangerequest) | **Post** /changeRequests/{changeRequestId} | Update a request details
*ComplianceAPI* | [**GetDirectiveComplianceId**](docs/ComplianceAPI.md#getdirectivecomplianceid) | **Get** /compliance/directives/{directiveId} | Compliance details by directive
*ComplianceAPI* | [**GetDirectivesCompliance**](docs/ComplianceAPI.md#getdirectivescompliance) | **Get** /compliance/directives | Compliance details for all directives
*ComplianceAPI* | [**GetGlobalCompliance**](docs/ComplianceAPI.md#getglobalcompliance) | **Get** /compliance | Global compliance
*ComplianceAPI* | [**GetNodeCompliance**](docs/ComplianceAPI.md#getnodecompliance) | **Get** /compliance/nodes/{nodeId} | Compliance details by node
*ComplianceAPI* | [**GetNodesCompliance**](docs/ComplianceAPI.md#getnodescompliance) | **Get** /compliance/nodes | Compliance details for all nodes
*ComplianceAPI* | [**GetRuleCompliance**](docs/ComplianceAPI.md#getrulecompliance) | **Get** /compliance/rules/{ruleId} | Compliance details by rule
*ComplianceAPI* | [**GetRulesCompliance**](docs/ComplianceAPI.md#getrulescompliance) | **Get** /compliance/rules | Compliance details for all rules
*DataSourcesAPI* | [**CreateDataSource**](docs/DataSourcesAPI.md#createdatasource) | **Put** /datasources | Create a data source
*DataSourcesAPI* | [**DeleteDataSource**](docs/DataSourcesAPI.md#deletedatasource) | **Delete** /datasources/{datasourceId} | Delete a data source
*DataSourcesAPI* | [**GetAllDataSources**](docs/DataSourcesAPI.md#getalldatasources) | **Get** /datasources | List all data sources
*DataSourcesAPI* | [**GetDataSource**](docs/DataSourcesAPI.md#getdatasource) | **Get** /datasources/{datasourceId} | Get data source configuration
*DataSourcesAPI* | [**ReloadAllDatasourcesAllNodes**](docs/DataSourcesAPI.md#reloadalldatasourcesallnodes) | **Post** /datasources/reload | Update properties from data sources
*DataSourcesAPI* | [**ReloadAllDatasourcesOneNode**](docs/DataSourcesAPI.md#reloadalldatasourcesonenode) | **Post** /nodes/{nodeId}/fetchData | Update properties for one node from all data sources
*DataSourcesAPI* | [**ReloadOneDatasourceAllNodes**](docs/DataSourcesAPI.md#reloadonedatasourceallnodes) | **Post** /datasources/reload/{datasourceId} | Update properties from data sources
*DataSourcesAPI* | [**ReloadOneDatasourceOneNode**](docs/DataSourcesAPI.md#reloadonedatasourceonenode) | **Post** /nodes/{nodeId}/fetchData/{datasourceId} | Update properties for one node from a data source
*DataSourcesAPI* | [**UpdateDataSource**](docs/DataSourcesAPI.md#updatedatasource) | **Post** /datasources/{datasourceId} | Update a data source configuration
*DirectivesAPI* | [**CheckDirective**](docs/DirectivesAPI.md#checkdirective) | **Post** /directives/{directiveId}/check | Check that update on a directive is valid
*DirectivesAPI* | [**CreateDirective**](docs/DirectivesAPI.md#createdirective) | **Put** /directives | Create a directive
*DirectivesAPI* | [**DeleteDirective**](docs/DirectivesAPI.md#deletedirective) | **Delete** /directives/{directiveId} | Delete a directive
*DirectivesAPI* | [**DirectiveDetails**](docs/DirectivesAPI.md#directivedetails) | **Get** /directives/{directiveId} | Get directive details
*DirectivesAPI* | [**ListDirectives**](docs/DirectivesAPI.md#listdirectives) | **Get** /directives | List all directives
*DirectivesAPI* | [**UpdateDirective**](docs/DirectivesAPI.md#updatedirective) | **Post** /directives/{directiveId} | Update a directive details
*GroupsAPI* | [**CreateGroup**](docs/GroupsAPI.md#creategroup) | **Put** /groups | Create a group
*GroupsAPI* | [**CreateGroupCategory**](docs/GroupsAPI.md#creategroupcategory) | **Put** /groups/categories | Create a group category
*GroupsAPI* | [**DeleteGroup**](docs/GroupsAPI.md#deletegroup) | **Delete** /groups/{groupId} | Delete a group
*GroupsAPI* | [**DeleteGroupCategory**](docs/GroupsAPI.md#deletegroupcategory) | **Delete** /groups/categories/{groupCategoryId} | Delete group category
*GroupsAPI* | [**GetGroupCategoryDetails**](docs/GroupsAPI.md#getgroupcategorydetails) | **Get** /groups/categories/{groupCategoryId} | Get group category details
*GroupsAPI* | [**GetGroupTree**](docs/GroupsAPI.md#getgrouptree) | **Get** /groups/tree | Get groups tree
*GroupsAPI* | [**GroupDetails**](docs/GroupsAPI.md#groupdetails) | **Get** /groups/{groupId} | Get group details
*GroupsAPI* | [**ListGroups**](docs/GroupsAPI.md#listgroups) | **Get** /groups | List all groups
*GroupsAPI* | [**ReloadGroup**](docs/GroupsAPI.md#reloadgroup) | **Post** /groups/{groupId}/reload | Reload a group
*GroupsAPI* | [**UpdateGroup**](docs/GroupsAPI.md#updategroup) | **Post** /groups/{groupId} | Update group details
*GroupsAPI* | [**UpdateGroupCategory**](docs/GroupsAPI.md#updategroupcategory) | **Post** /groups/categories/{groupCategoryId} | Update group category details
*InventoriesAPI* | [**FileWatcherRestart**](docs/InventoriesAPI.md#filewatcherrestart) | **Post** /inventories/watcher/restart | Restart inventory watcher
*InventoriesAPI* | [**FileWatcherStart**](docs/InventoriesAPI.md#filewatcherstart) | **Post** /inventories/watcher/start | Start inventory watcher
*InventoriesAPI* | [**FileWatcherStop**](docs/InventoriesAPI.md#filewatcherstop) | **Post** /inventories/watcher/stop | Stop inventory watcher
*InventoriesAPI* | [**QueueInformation**](docs/InventoriesAPI.md#queueinformation) | **Get** /inventories/info | Get information about inventory processing queue
*InventoriesAPI* | [**UploadInventory**](docs/InventoriesAPI.md#uploadinventory) | **Post** /inventories/upload | Upload an inventory for processing
*NodesAPI* | [**ApplyPolicy**](docs/NodesAPI.md#applypolicy) | **Post** /nodes/{nodeId}/applyPolicy | Trigger an agent run
*NodesAPI* | [**ApplyPolicyAllNodes**](docs/NodesAPI.md#applypolicyallnodes) | **Post** /nodes/applyPolicy | Trigger an agent run on all nodes
*NodesAPI* | [**ChangePendingNodeStatus**](docs/NodesAPI.md#changependingnodestatus) | **Post** /nodes/pending/{nodeId} | Update pending Node status
*NodesAPI* | [**CreateNodes**](docs/NodesAPI.md#createnodes) | **Put** /nodes | Create one or several new nodes
*NodesAPI* | [**DeleteNode**](docs/NodesAPI.md#deletenode) | **Delete** /nodes/{nodeId} | Delete a node
*NodesAPI* | [**GetNodesStatus**](docs/NodesAPI.md#getnodesstatus) | **Get** /nodes/status | Get nodes acceptation status
*NodesAPI* | [**ListAcceptedNodes**](docs/NodesAPI.md#listacceptednodes) | **Get** /nodes | List managed nodes
*NodesAPI* | [**ListPendingNodes**](docs/NodesAPI.md#listpendingnodes) | **Get** /nodes/pending | List pending nodes
*NodesAPI* | [**NodeDetails**](docs/NodesAPI.md#nodedetails) | **Get** /nodes/{nodeId} | Get information about a node
*NodesAPI* | [**NodeInheritedProperties**](docs/NodesAPI.md#nodeinheritedproperties) | **Get** /nodes/{nodeId}/inheritedProperties | Get inherited node properties for a node
*NodesAPI* | [**UpdateNode**](docs/NodesAPI.md#updatenode) | **Post** /nodes/{nodeId} | Update node settings and properties
*OpenSCAPAPI* | [**OpenscapReport**](docs/OpenSCAPAPI.md#openscapreport) | **Get** /openscap/report/{nodeId} | Get an OpenSCAP report
*ParametersAPI* | [**CreateParameter**](docs/ParametersAPI.md#createparameter) | **Put** /parameters | Create a new property
*ParametersAPI* | [**DeleteParameter**](docs/ParametersAPI.md#deleteparameter) | **Delete** /parameters/{parameterId} | Delete a global parameter
*ParametersAPI* | [**ListParameters**](docs/ParametersAPI.md#listparameters) | **Get** /parameters | List all global properties
*ParametersAPI* | [**ParameterDetails**](docs/ParametersAPI.md#parameterdetails) | **Get** /parameters/{parameterId} | Get the value of a global property
*ParametersAPI* | [**UpdateParameter**](docs/ParametersAPI.md#updateparameter) | **Post** /parameters/{parameterId} | Update a global property&#39;s value
*RulesAPI* | [**CreateRule**](docs/RulesAPI.md#createrule) | **Put** /rules | Create a rule
*RulesAPI* | [**CreateRuleCategory**](docs/RulesAPI.md#createrulecategory) | **Put** /rules/categories | Create a rule category
*RulesAPI* | [**DeleteRule**](docs/RulesAPI.md#deleterule) | **Delete** /rules/{ruleId} | Delete a rule
*RulesAPI* | [**DeleteRuleCategory**](docs/RulesAPI.md#deleterulecategory) | **Delete** /rules/categories/{ruleCategoryId} | Delete group category
*RulesAPI* | [**GetRuleCategoryDetails**](docs/RulesAPI.md#getrulecategorydetails) | **Get** /rules/categories/{ruleCategoryId} | Get rule category details
*RulesAPI* | [**GetRuleTree**](docs/RulesAPI.md#getruletree) | **Get** /rules/tree | Get rules tree
*RulesAPI* | [**ListRules**](docs/RulesAPI.md#listrules) | **Get** /rules | List all rules
*RulesAPI* | [**RuleDetails**](docs/RulesAPI.md#ruledetails) | **Get** /rules/{ruleId} | Get a rule details
*RulesAPI* | [**UpdateRule**](docs/RulesAPI.md#updaterule) | **Post** /rules/{ruleId} | Update a rule details
*RulesAPI* | [**UpdateRuleCategory**](docs/RulesAPI.md#updaterulecategory) | **Post** /rules/categories/{ruleCategoryId} | Update rule category details
*ScaleOutRelayAPI* | [**DemoteToNode**](docs/ScaleOutRelayAPI.md#demotetonode) | **Post** /scaleoutrelay/demote/{nodeId} | Demote a relay to simple node
*ScaleOutRelayAPI* | [**PromoteToRelay**](docs/ScaleOutRelayAPI.md#promotetorelay) | **Post** /scaleoutrelay/promote/{nodeId} | Promote a node to relay
*SecretManagementAPI* | [**AddSecret**](docs/SecretManagementAPI.md#addsecret) | **Put** /secret | Create a secret
*SecretManagementAPI* | [**DeleteSecret**](docs/SecretManagementAPI.md#deletesecret) | **Delete** /secret/{name} | Delete a secret
*SecretManagementAPI* | [**GetAllSecrets**](docs/SecretManagementAPI.md#getallsecrets) | **Get** /secret | List all secrets
*SecretManagementAPI* | [**GetSecret**](docs/SecretManagementAPI.md#getsecret) | **Get** /secret/{name} | Get one secret
*SecretManagementAPI* | [**UpdateSecret**](docs/SecretManagementAPI.md#updatesecret) | **Post** /secret | Update a secret
*SettingsAPI* | [**GetAllSettings**](docs/SettingsAPI.md#getallsettings) | **Get** /settings | List all settings
*SettingsAPI* | [**GetAllowedNetworks**](docs/SettingsAPI.md#getallowednetworks) | **Get** /settings/allowed_networks/{nodeId} | Get allowed networks for a policy server
*SettingsAPI* | [**GetSetting**](docs/SettingsAPI.md#getsetting) | **Get** /settings/{settingId} | Get the value of a setting
*SettingsAPI* | [**ModifyAllowedNetworks**](docs/SettingsAPI.md#modifyallowednetworks) | **Post** /settings/allowed_networks/{nodeId}/diff | Modify allowed networks for a policy server
*SettingsAPI* | [**ModifySetting**](docs/SettingsAPI.md#modifysetting) | **Post** /settings/{settingId} | Set the value of a setting
*SettingsAPI* | [**SetAllowedNetworks**](docs/SettingsAPI.md#setallowednetworks) | **Post** /settings/allowed_networks/{nodeId} | Set allowed networks for a policy server
*StatusAPI* | [**None**](docs/StatusAPI.md#none) | **Get** /status | Check if Rudder is alive
*SystemAPI* | [**CreateArchive**](docs/SystemAPI.md#createarchive) | **Post** /system/archives/{archiveKind} | Create an archive
*SystemAPI* | [**GetHealthcheckResult**](docs/SystemAPI.md#gethealthcheckresult) | **Get** /system/healthcheck | Get healthcheck
*SystemAPI* | [**GetStatus**](docs/SystemAPI.md#getstatus) | **Get** /system/status | Get server status
*SystemAPI* | [**GetSystemInfo**](docs/SystemAPI.md#getsysteminfo) | **Get** /system/info | Get server information
*SystemAPI* | [**GetZipArchive**](docs/SystemAPI.md#getziparchive) | **Get** /system/archives/{archiveKind}/zip/{commitId} | Get an archive as a ZIP
*SystemAPI* | [**ListArchives**](docs/SystemAPI.md#listarchives) | **Get** /system/archives/{archiveKind} | List archives
*SystemAPI* | [**PurgeSoftware**](docs/SystemAPI.md#purgesoftware) | **Post** /system/maintenance/purgeSoftware | Trigger batch for cleaning unreferenced software
*SystemAPI* | [**RegeneratePolicies**](docs/SystemAPI.md#regeneratepolicies) | **Post** /system/regenerate/policies | Trigger a new policy generation
*SystemAPI* | [**ReloadAll**](docs/SystemAPI.md#reloadall) | **Post** /system/reload | Reload both techniques and dynamic groups
*SystemAPI* | [**ReloadGroups**](docs/SystemAPI.md#reloadgroups) | **Post** /system/reload/groups | Reload dynamic groups
*SystemAPI* | [**ReloadTechniques**](docs/SystemAPI.md#reloadtechniques) | **Post** /system/reload/techniques | Reload techniques
*SystemAPI* | [**RestoreArchive**](docs/SystemAPI.md#restorearchive) | **Post** /system/archives/{archiveKind}/restore/{archiveRestoreKind} | Restore an archive
*SystemAPI* | [**UpdatePolicies**](docs/SystemAPI.md#updatepolicies) | **Post** /system/update/policies | Trigger update of policies
*SystemUpdateCampaignsAPI* | [**GetCampaignEventResult**](docs/SystemUpdateCampaignsAPI.md#getcampaigneventresult) | **Get** /systemUpdate/events/{id}/result | Get a campaign event result
*SystemUpdateCampaignsAPI* | [**GetCampaignResults**](docs/SystemUpdateCampaignsAPI.md#getcampaignresults) | **Get** /systemUpdate/campaign/{id}/history | Get a campaign result history
*SystemUpdateCampaignsAPI* | [**GetSystemUpdateResultForNode**](docs/SystemUpdateCampaignsAPI.md#getsystemupdateresultfornode) | **Get** /systemUpdate/events/{id}/result/{nodeId} | Get detailed campaign event result for a Node
*TechniquesAPI* | [**CreateTechnique**](docs/TechniquesAPI.md#createtechnique) | **Put** /techniques | Create technique
*TechniquesAPI* | [**DeleteTechnique**](docs/TechniquesAPI.md#deletetechnique) | **Delete** /techniques/{techniqueId}/{techniqueVersion} | Delete technique
*TechniquesAPI* | [**GetTechniqueAllVersion**](docs/TechniquesAPI.md#gettechniqueallversion) | **Get** /techniques/{techniqueId} | Technique metadata by ID
*TechniquesAPI* | [**GetTechniqueAllVersionId**](docs/TechniquesAPI.md#gettechniqueallversionid) | **Get** /techniques/{techniqueId}/{techniqueVersion} | Technique metadata by version and ID
*TechniquesAPI* | [**GetTechniquesResources**](docs/TechniquesAPI.md#gettechniquesresources) | **Get** /techniques/{techniqueId}/{techniqueVersion}/resources | Technique&#39;s resources
*TechniquesAPI* | [**ListTechniqueVersionDirectives**](docs/TechniquesAPI.md#listtechniqueversiondirectives) | **Get** /techniques/{techniqueId}/{techniqueVersion}/directives | List all directives based on a version of a technique
*TechniquesAPI* | [**ListTechniques**](docs/TechniquesAPI.md#listtechniques) | **Get** /techniques | List all techniques
*TechniquesAPI* | [**ListTechniquesDirectives**](docs/TechniquesAPI.md#listtechniquesdirectives) | **Get** /techniques/{techniqueId}/directives | List all directives based on a technique
*TechniquesAPI* | [**ListTechniquesVersions**](docs/TechniquesAPI.md#listtechniquesversions) | **Get** /techniques/versions | List versions
*TechniquesAPI* | [**Methods**](docs/TechniquesAPI.md#methods) | **Get** /methods | List methods
*TechniquesAPI* | [**ReloadMethods**](docs/TechniquesAPI.md#reloadmethods) | **Post** /methods/reload | Reload methods
*TechniquesAPI* | [**TechniqueCategories**](docs/TechniquesAPI.md#techniquecategories) | **Get** /techniques/categories | List categories
*TechniquesAPI* | [**TechniqueRevisions**](docs/TechniquesAPI.md#techniquerevisions) | **Get** /techniques/{techniqueId}/{techniqueVersion}/revisions | Technique&#39;s revisions
*TechniquesAPI* | [**Techniques**](docs/TechniquesAPI.md#techniques) | **Post** /techniques/reload | Reload techniques
*TechniquesAPI* | [**UpdateTechnique**](docs/TechniquesAPI.md#updatetechnique) | **Post** /techniques/{techniqueId}/{techniqueVersion} | Update technique
*UserManagementAPI* | [**AddUser**](docs/UserManagementAPI.md#adduser) | **Post** /usermanagement | Add user
*UserManagementAPI* | [**DeleteUser**](docs/UserManagementAPI.md#deleteuser) | **Delete** /usermanagement/{username} | Delete an user
*UserManagementAPI* | [**GetRole**](docs/UserManagementAPI.md#getrole) | **Get** /usermanagement/roles | List all roles
*UserManagementAPI* | [**GetUserInfo**](docs/UserManagementAPI.md#getuserinfo) | **Get** /usermanagement/users | List all users
*UserManagementAPI* | [**ReloadUserConf**](docs/UserManagementAPI.md#reloaduserconf) | **Get** /usermanagement/users/reload | Reload user
*UserManagementAPI* | [**UpdateUser**](docs/UserManagementAPI.md#updateuser) | **Post** /usermanagement/update/{username} | Update user&#39;s infos


## Documentation For Models

 - [AcceptChangeRequest200Response](docs/AcceptChangeRequest200Response.md)
 - [AcceptChangeRequestRequest](docs/AcceptChangeRequestRequest.md)
 - [AddSecret200Response](docs/AddSecret200Response.md)
 - [AddUser200Response](docs/AddUser200Response.md)
 - [AddUser200ResponseData](docs/AddUser200ResponseData.md)
 - [AddUser200ResponseDataAddedUser](docs/AddUser200ResponseDataAddedUser.md)
 - [AgentKey](docs/AgentKey.md)
 - [AllCampaigns200Response](docs/AllCampaigns200Response.md)
 - [AllCampaigns200ResponseData](docs/AllCampaigns200ResponseData.md)
 - [ApiEndpointsInner](docs/ApiEndpointsInner.md)
 - [ApiGeneralInformations200Response](docs/ApiGeneralInformations200Response.md)
 - [ApiGeneralInformations200ResponseData](docs/ApiGeneralInformations200ResponseData.md)
 - [ApiInformations200Response](docs/ApiInformations200Response.md)
 - [ApiInformations200ResponseData](docs/ApiInformations200ResponseData.md)
 - [ApiInformations200ResponseDataEndpointsInner](docs/ApiInformations200ResponseDataEndpointsInner.md)
 - [ApiSubInformations200Response](docs/ApiSubInformations200Response.md)
 - [ApiVersion](docs/ApiVersion.md)
 - [ApiVersionAllInner](docs/ApiVersionAllInner.md)
 - [ApiVersions](docs/ApiVersions.md)
 - [ApplyPolicyAllNodes200Response](docs/ApplyPolicyAllNodes200Response.md)
 - [ApplyPolicyAllNodes200ResponseDataInner](docs/ApplyPolicyAllNodes200ResponseDataInner.md)
 - [BrandingConf](docs/BrandingConf.md)
 - [CampaignDetails](docs/CampaignDetails.md)
 - [CampaignDetailsDetails](docs/CampaignDetailsDetails.md)
 - [CampaignDetailsInfo](docs/CampaignDetailsInfo.md)
 - [CampaignDetailsInfoSchedule](docs/CampaignDetailsInfoSchedule.md)
 - [CampaignDetailsInfoStatus](docs/CampaignDetailsInfoStatus.md)
 - [CampaignEventDetails](docs/CampaignEventDetails.md)
 - [CampaignEventDetailsStatus](docs/CampaignEventDetailsStatus.md)
 - [CampaignEventNodeResult](docs/CampaignEventNodeResult.md)
 - [CampaignEventNodeResultNodesInner](docs/CampaignEventNodeResultNodesInner.md)
 - [CampaignEventNodeResultNodesInnerResult](docs/CampaignEventNodeResultNodesInnerResult.md)
 - [CampaignEventNodeResultNodesInnerResultSoftwareUpdatedInner](docs/CampaignEventNodeResultNodesInnerResultSoftwareUpdatedInner.md)
 - [CampaignEventResult](docs/CampaignEventResult.md)
 - [CampaignEventResultNodesInner](docs/CampaignEventResultNodesInner.md)
 - [CampaignScheduleMonthly](docs/CampaignScheduleMonthly.md)
 - [CampaignScheduleMonthlyEnd](docs/CampaignScheduleMonthlyEnd.md)
 - [CampaignScheduleMonthlyStart](docs/CampaignScheduleMonthlyStart.md)
 - [CampaignScheduleOneshot](docs/CampaignScheduleOneshot.md)
 - [CampaignScheduleWeekly](docs/CampaignScheduleWeekly.md)
 - [CampaignScheduleWeeklyEnd](docs/CampaignScheduleWeeklyEnd.md)
 - [CampaignScheduleWeeklyStart](docs/CampaignScheduleWeeklyStart.md)
 - [CampaignSoftwareUpdate](docs/CampaignSoftwareUpdate.md)
 - [CampaignSoftwareUpdateSoftwareUpdateInner](docs/CampaignSoftwareUpdateSoftwareUpdateInner.md)
 - [CampaignStatusArchived](docs/CampaignStatusArchived.md)
 - [CampaignStatusDisabled](docs/CampaignStatusDisabled.md)
 - [CampaignStatusEnabled](docs/CampaignStatusEnabled.md)
 - [CampaignSystemUpdate](docs/CampaignSystemUpdate.md)
 - [CategoriesTree](docs/CategoriesTree.md)
 - [ChangePendingNodeStatus200Response](docs/ChangePendingNodeStatus200Response.md)
 - [ChangePendingNodeStatus200ResponseData](docs/ChangePendingNodeStatus200ResponseData.md)
 - [ChangePendingNodeStatusRequest](docs/ChangePendingNodeStatusRequest.md)
 - [ChangeRequest](docs/ChangeRequest.md)
 - [ChangeRequestChanges](docs/ChangeRequestChanges.md)
 - [ChangeRequestChangesRulesInner](docs/ChangeRequestChangesRulesInner.md)
 - [ChangeRequestDetails200Response](docs/ChangeRequestDetails200Response.md)
 - [Check](docs/Check.md)
 - [CheckCVE200Response](docs/CheckCVE200Response.md)
 - [CheckCVE200ResponseData](docs/CheckCVE200ResponseData.md)
 - [CheckDirective200Response](docs/CheckDirective200Response.md)
 - [Color](docs/Color.md)
 - [ComplianceDirectiveId](docs/ComplianceDirectiveId.md)
 - [ComplianceDirectiveIdData](docs/ComplianceDirectiveIdData.md)
 - [CreateArchive200Response](docs/CreateArchive200Response.md)
 - [CreateArchive200ResponseData](docs/CreateArchive200ResponseData.md)
 - [CreateDataSource200Response](docs/CreateDataSource200Response.md)
 - [CreateDataSource200ResponseData](docs/CreateDataSource200ResponseData.md)
 - [CreateDirective200Response](docs/CreateDirective200Response.md)
 - [CreateGroup200Response](docs/CreateGroup200Response.md)
 - [CreateGroupCategory200Response](docs/CreateGroupCategory200Response.md)
 - [CreateGroupCategory200ResponseData](docs/CreateGroupCategory200ResponseData.md)
 - [CreateNodes200Response](docs/CreateNodes200Response.md)
 - [CreateNodes200ResponseData](docs/CreateNodes200ResponseData.md)
 - [CreateParameter200Response](docs/CreateParameter200Response.md)
 - [CreateRule200Response](docs/CreateRule200Response.md)
 - [CreateRuleCategory200Response](docs/CreateRuleCategory200Response.md)
 - [CreateRuleCategory200ResponseData](docs/CreateRuleCategory200ResponseData.md)
 - [CreateTechnique200Response](docs/CreateTechnique200Response.md)
 - [CreateTechnique200ResponseData](docs/CreateTechnique200ResponseData.md)
 - [CreateTechnique200ResponseDataTechniques](docs/CreateTechnique200ResponseDataTechniques.md)
 - [CveCheck](docs/CveCheck.md)
 - [CveCheckPackagesInner](docs/CveCheckPackagesInner.md)
 - [CveCheckScore](docs/CveCheckScore.md)
 - [CveDetails](docs/CveDetails.md)
 - [CveDetailsCvssv2](docs/CveDetailsCvssv2.md)
 - [CveDetailsCvssv3](docs/CveDetailsCvssv3.md)
 - [Datasource](docs/Datasource.md)
 - [DatasourceRunParameters](docs/DatasourceRunParameters.md)
 - [DatasourceRunParametersSchedule](docs/DatasourceRunParametersSchedule.md)
 - [DatasourceType](docs/DatasourceType.md)
 - [DatasourceTypeParameters](docs/DatasourceTypeParameters.md)
 - [DatasourceTypeParametersHeadersInner](docs/DatasourceTypeParametersHeadersInner.md)
 - [DatasourceTypeParametersRequestMode](docs/DatasourceTypeParametersRequestMode.md)
 - [DeclineChangeRequest200Response](docs/DeclineChangeRequest200Response.md)
 - [DeleteCampaign200Response](docs/DeleteCampaign200Response.md)
 - [DeleteCampaignEvent200Response](docs/DeleteCampaignEvent200Response.md)
 - [DeleteDataSource200Response](docs/DeleteDataSource200Response.md)
 - [DeleteDirective200Response](docs/DeleteDirective200Response.md)
 - [DeleteGroup200Response](docs/DeleteGroup200Response.md)
 - [DeleteGroupCategory200Response](docs/DeleteGroupCategory200Response.md)
 - [DeleteNode200Response](docs/DeleteNode200Response.md)
 - [DeleteParameter200Response](docs/DeleteParameter200Response.md)
 - [DeleteParameter500Response](docs/DeleteParameter500Response.md)
 - [DeleteRule200Response](docs/DeleteRule200Response.md)
 - [DeleteRuleCategory200Response](docs/DeleteRuleCategory200Response.md)
 - [DeleteRuleCategory200ResponseData](docs/DeleteRuleCategory200ResponseData.md)
 - [DeleteSecret200Response](docs/DeleteSecret200Response.md)
 - [DeleteTechnique200Response](docs/DeleteTechnique200Response.md)
 - [DeleteTechnique200ResponseData](docs/DeleteTechnique200ResponseData.md)
 - [DeleteTechnique200ResponseDataTechniques](docs/DeleteTechnique200ResponseDataTechniques.md)
 - [DeleteUser200Response](docs/DeleteUser200Response.md)
 - [DeleteUser200ResponseData](docs/DeleteUser200ResponseData.md)
 - [DeleteUser200ResponseDataDeletedUser](docs/DeleteUser200ResponseDataDeletedUser.md)
 - [DemoteToNode200Response](docs/DemoteToNode200Response.md)
 - [Directive](docs/Directive.md)
 - [DirectiveDetails200Response](docs/DirectiveDetails200Response.md)
 - [DirectiveNew](docs/DirectiveNew.md)
 - [DirectiveNodeCompliance](docs/DirectiveNodeCompliance.md)
 - [DirectiveRuleCompliance](docs/DirectiveRuleCompliance.md)
 - [DirectiveRuleComplianceComponentsInner](docs/DirectiveRuleComplianceComponentsInner.md)
 - [DirectiveTagsInner](docs/DirectiveTagsInner.md)
 - [EditorTechnique](docs/EditorTechnique.md)
 - [FileWatcherRestart200Response](docs/FileWatcherRestart200Response.md)
 - [FileWatcherStart200Response](docs/FileWatcherStart200Response.md)
 - [FileWatcherStop200Response](docs/FileWatcherStop200Response.md)
 - [GetAllCampaignEvents200Response](docs/GetAllCampaignEvents200Response.md)
 - [GetAllCampaignEvents200ResponseData](docs/GetAllCampaignEvents200ResponseData.md)
 - [GetAllCve200Response](docs/GetAllCve200Response.md)
 - [GetAllCve200ResponseData](docs/GetAllCve200ResponseData.md)
 - [GetAllDataSources200Response](docs/GetAllDataSources200Response.md)
 - [GetAllDataSources200ResponseData](docs/GetAllDataSources200ResponseData.md)
 - [GetAllSecrets200Response](docs/GetAllSecrets200Response.md)
 - [GetAllSecrets200ResponseData](docs/GetAllSecrets200ResponseData.md)
 - [GetAllSecrets200ResponseDataSecretsInner](docs/GetAllSecrets200ResponseDataSecretsInner.md)
 - [GetAllSettings200Response](docs/GetAllSettings200Response.md)
 - [GetAllSettings200ResponseData](docs/GetAllSettings200ResponseData.md)
 - [GetAllSettings200ResponseDataSettings](docs/GetAllSettings200ResponseDataSettings.md)
 - [GetAllSettings200ResponseDataSettingsAllowedNetworksInner](docs/GetAllSettings200ResponseDataSettingsAllowedNetworksInner.md)
 - [GetAllowedNetworks200Response](docs/GetAllowedNetworks200Response.md)
 - [GetAllowedNetworks200ResponseData](docs/GetAllowedNetworks200ResponseData.md)
 - [GetBrandingConf200Response](docs/GetBrandingConf200Response.md)
 - [GetBrandingConf200ResponseData](docs/GetBrandingConf200ResponseData.md)
 - [GetCVECheckConfiguration200Response](docs/GetCVECheckConfiguration200Response.md)
 - [GetCVECheckConfiguration200ResponseData](docs/GetCVECheckConfiguration200ResponseData.md)
 - [GetCVEList200Response](docs/GetCVEList200Response.md)
 - [GetCVEListRequest](docs/GetCVEListRequest.md)
 - [GetCampaign200Response](docs/GetCampaign200Response.md)
 - [GetCampaignEventResult200Response](docs/GetCampaignEventResult200Response.md)
 - [GetCampaignEventResult200ResponseData](docs/GetCampaignEventResult200ResponseData.md)
 - [GetCampaignResults200Response](docs/GetCampaignResults200Response.md)
 - [GetCampaignResults200ResponseData](docs/GetCampaignResults200ResponseData.md)
 - [GetCve200Response](docs/GetCve200Response.md)
 - [GetDataSource200Response](docs/GetDataSource200Response.md)
 - [GetDirectiveComplianceId200Response](docs/GetDirectiveComplianceId200Response.md)
 - [GetDirectivesCompliance200Response](docs/GetDirectivesCompliance200Response.md)
 - [GetDirectivesCompliance200ResponseData](docs/GetDirectivesCompliance200ResponseData.md)
 - [GetDirectivesCompliance200ResponseDataDirectivesCompliance](docs/GetDirectivesCompliance200ResponseDataDirectivesCompliance.md)
 - [GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails](docs/GetDirectivesCompliance200ResponseDataDirectivesComplianceComplianceDetails.md)
 - [GetEventsCampaign200Response](docs/GetEventsCampaign200Response.md)
 - [GetGlobalCompliance200Response](docs/GetGlobalCompliance200Response.md)
 - [GetGlobalCompliance200ResponseData](docs/GetGlobalCompliance200ResponseData.md)
 - [GetGlobalCompliance200ResponseDataGlobalCompliance](docs/GetGlobalCompliance200ResponseDataGlobalCompliance.md)
 - [GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails](docs/GetGlobalCompliance200ResponseDataGlobalComplianceComplianceDetails.md)
 - [GetGroupCategoryDetails200Response](docs/GetGroupCategoryDetails200Response.md)
 - [GetGroupTree200Response](docs/GetGroupTree200Response.md)
 - [GetGroupTree200ResponseData](docs/GetGroupTree200ResponseData.md)
 - [GetHealthcheckResult200Response](docs/GetHealthcheckResult200Response.md)
 - [GetLastCVECheck200Response](docs/GetLastCVECheck200Response.md)
 - [GetLastCVECheck200ResponseData](docs/GetLastCVECheck200ResponseData.md)
 - [GetNodeCompliance200Response](docs/GetNodeCompliance200Response.md)
 - [GetNodesCompliance200Response](docs/GetNodesCompliance200Response.md)
 - [GetNodesCompliance200ResponseData](docs/GetNodesCompliance200ResponseData.md)
 - [GetNodesCompliance200ResponseDataNodesInner](docs/GetNodesCompliance200ResponseDataNodesInner.md)
 - [GetNodesStatus200Response](docs/GetNodesStatus200Response.md)
 - [GetNodesStatus200ResponseData](docs/GetNodesStatus200ResponseData.md)
 - [GetNodesStatus200ResponseDataNodesInner](docs/GetNodesStatus200ResponseDataNodesInner.md)
 - [GetRole200Response](docs/GetRole200Response.md)
 - [GetRole200ResponseDataInner](docs/GetRole200ResponseDataInner.md)
 - [GetRuleCategoryDetails200Response](docs/GetRuleCategoryDetails200Response.md)
 - [GetRuleCategoryDetails200ResponseData](docs/GetRuleCategoryDetails200ResponseData.md)
 - [GetRuleCompliance200Response](docs/GetRuleCompliance200Response.md)
 - [GetRuleTree200Response](docs/GetRuleTree200Response.md)
 - [GetRuleTree200ResponseData](docs/GetRuleTree200ResponseData.md)
 - [GetRulesCompliance200Response](docs/GetRulesCompliance200Response.md)
 - [GetRulesCompliance200ResponseData](docs/GetRulesCompliance200ResponseData.md)
 - [GetRulesCompliance200ResponseDataRulesInner](docs/GetRulesCompliance200ResponseDataRulesInner.md)
 - [GetSecret200Response](docs/GetSecret200Response.md)
 - [GetSetting200Response](docs/GetSetting200Response.md)
 - [GetSetting200ResponseData](docs/GetSetting200ResponseData.md)
 - [GetStatus200Response](docs/GetStatus200Response.md)
 - [GetStatus200ResponseData](docs/GetStatus200ResponseData.md)
 - [GetSystemInfo200Response](docs/GetSystemInfo200Response.md)
 - [GetSystemInfo200ResponseData](docs/GetSystemInfo200ResponseData.md)
 - [GetSystemInfo200ResponseDataRudder](docs/GetSystemInfo200ResponseDataRudder.md)
 - [GetSystemUpdateResultForNode200Response](docs/GetSystemUpdateResultForNode200Response.md)
 - [GetSystemUpdateResultForNode200ResponseData](docs/GetSystemUpdateResultForNode200ResponseData.md)
 - [GetTechniqueAllVersion200Response](docs/GetTechniqueAllVersion200Response.md)
 - [GetTechniqueAllVersion200ResponseData](docs/GetTechniqueAllVersion200ResponseData.md)
 - [GetTechniqueAllVersion200ResponseDataTechniquesInner](docs/GetTechniqueAllVersion200ResponseDataTechniquesInner.md)
 - [GetTechniquesResources200Response](docs/GetTechniquesResources200Response.md)
 - [GetTechniquesResources200ResponseData](docs/GetTechniquesResources200ResponseData.md)
 - [GetUserInfo200Response](docs/GetUserInfo200Response.md)
 - [GetUserInfo200ResponseData](docs/GetUserInfo200ResponseData.md)
 - [Group](docs/Group.md)
 - [GroupCategory](docs/GroupCategory.md)
 - [GroupCategoryUpdate](docs/GroupCategoryUpdate.md)
 - [GroupDetails200Response](docs/GroupDetails200Response.md)
 - [GroupNew](docs/GroupNew.md)
 - [GroupNewQuery](docs/GroupNewQuery.md)
 - [GroupPropertiesInner](docs/GroupPropertiesInner.md)
 - [GroupQuery](docs/GroupQuery.md)
 - [GroupQueryWhereInner](docs/GroupQueryWhereInner.md)
 - [GroupUpdate](docs/GroupUpdate.md)
 - [Import200Response](docs/Import200Response.md)
 - [Import200ResponseData](docs/Import200ResponseData.md)
 - [ListAcceptedNodes200Response](docs/ListAcceptedNodes200Response.md)
 - [ListAcceptedNodes200ResponseData](docs/ListAcceptedNodes200ResponseData.md)
 - [ListArchives200Response](docs/ListArchives200Response.md)
 - [ListArchives200ResponseData](docs/ListArchives200ResponseData.md)
 - [ListArchives200ResponseDataFullInner](docs/ListArchives200ResponseDataFullInner.md)
 - [ListChangeRequests200Response](docs/ListChangeRequests200Response.md)
 - [ListChangeRequests200ResponseData](docs/ListChangeRequests200ResponseData.md)
 - [ListDirectives200Response](docs/ListDirectives200Response.md)
 - [ListDirectives200ResponseData](docs/ListDirectives200ResponseData.md)
 - [ListGroups200Response](docs/ListGroups200Response.md)
 - [ListGroups200ResponseData](docs/ListGroups200ResponseData.md)
 - [ListParameters200Response](docs/ListParameters200Response.md)
 - [ListParameters200ResponseData](docs/ListParameters200ResponseData.md)
 - [ListPendingNodes200Response](docs/ListPendingNodes200Response.md)
 - [ListRules200Response](docs/ListRules200Response.md)
 - [ListRules200ResponseData](docs/ListRules200ResponseData.md)
 - [ListTechniqueVersionDirectives200Response](docs/ListTechniqueVersionDirectives200Response.md)
 - [ListTechniques200Response](docs/ListTechniques200Response.md)
 - [ListTechniques200ResponseData](docs/ListTechniques200ResponseData.md)
 - [ListTechniquesDirectives200Response](docs/ListTechniquesDirectives200Response.md)
 - [ListTechniquesVersions200Response](docs/ListTechniquesVersions200Response.md)
 - [ListTechniquesVersions200ResponseData](docs/ListTechniquesVersions200ResponseData.md)
 - [ListUsers200Response](docs/ListUsers200Response.md)
 - [Logo](docs/Logo.md)
 - [MethodParameter](docs/MethodParameter.md)
 - [MethodParameterConstraints](docs/MethodParameterConstraints.md)
 - [Methods](docs/Methods.md)
 - [Methods200Response](docs/Methods200Response.md)
 - [Methods200ResponseData](docs/Methods200ResponseData.md)
 - [MethodsCondition](docs/MethodsCondition.md)
 - [MethodsDeprecated](docs/MethodsDeprecated.md)
 - [ModifyAllowedNetworks200Response](docs/ModifyAllowedNetworks200Response.md)
 - [ModifyAllowedNetworks200ResponseData](docs/ModifyAllowedNetworks200ResponseData.md)
 - [ModifyAllowedNetworksRequest](docs/ModifyAllowedNetworksRequest.md)
 - [ModifyAllowedNetworksRequestAllowedNetworks](docs/ModifyAllowedNetworksRequestAllowedNetworks.md)
 - [ModifySetting200Response](docs/ModifySetting200Response.md)
 - [ModifySettingRequest](docs/ModifySettingRequest.md)
 - [NodeAddInner](docs/NodeAddInner.md)
 - [NodeComplianceComponent](docs/NodeComplianceComponent.md)
 - [NodeComplianceComponentComplianceDetails](docs/NodeComplianceComponentComplianceDetails.md)
 - [NodeComplianceComponentValuesInner](docs/NodeComplianceComponentValuesInner.md)
 - [NodeComplianceComponentValuesInnerReportsInner](docs/NodeComplianceComponentValuesInnerReportsInner.md)
 - [NodeDetails200Response](docs/NodeDetails200Response.md)
 - [NodeFull](docs/NodeFull.md)
 - [NodeFullBios](docs/NodeFullBios.md)
 - [NodeFullControllersInner](docs/NodeFullControllersInner.md)
 - [NodeFullEnvironmentVariablesInner](docs/NodeFullEnvironmentVariablesInner.md)
 - [NodeFullFileSystemsInner](docs/NodeFullFileSystemsInner.md)
 - [NodeFullMachine](docs/NodeFullMachine.md)
 - [NodeFullManagementTechnologyDetails](docs/NodeFullManagementTechnologyDetails.md)
 - [NodeFullManagementTechnologyInner](docs/NodeFullManagementTechnologyInner.md)
 - [NodeFullMemoriesInner](docs/NodeFullMemoriesInner.md)
 - [NodeFullNetworkInterfacesInner](docs/NodeFullNetworkInterfacesInner.md)
 - [NodeFullOs](docs/NodeFullOs.md)
 - [NodeFullPortsInner](docs/NodeFullPortsInner.md)
 - [NodeFullProcessesInner](docs/NodeFullProcessesInner.md)
 - [NodeFullProcessorsInner](docs/NodeFullProcessorsInner.md)
 - [NodeFullPropertiesInner](docs/NodeFullPropertiesInner.md)
 - [NodeFullSlotsInner](docs/NodeFullSlotsInner.md)
 - [NodeFullSoftwareInner](docs/NodeFullSoftwareInner.md)
 - [NodeFullSoftwareInnerLicense](docs/NodeFullSoftwareInnerLicense.md)
 - [NodeFullSoftwareUpdateInner](docs/NodeFullSoftwareUpdateInner.md)
 - [NodeFullSoundInner](docs/NodeFullSoundInner.md)
 - [NodeFullStorageInner](docs/NodeFullStorageInner.md)
 - [NodeFullTimezone](docs/NodeFullTimezone.md)
 - [NodeFullVideosInner](docs/NodeFullVideosInner.md)
 - [NodeFullVirtualMachinesInner](docs/NodeFullVirtualMachinesInner.md)
 - [NodeInheritedProperties](docs/NodeInheritedProperties.md)
 - [NodeInheritedProperties200Response](docs/NodeInheritedProperties200Response.md)
 - [NodeInheritedPropertiesPropertiesInner](docs/NodeInheritedPropertiesPropertiesInner.md)
 - [NodeInheritedPropertiesPropertiesInnerHierarchyInner](docs/NodeInheritedPropertiesPropertiesInnerHierarchyInner.md)
 - [NodeSettings](docs/NodeSettings.md)
 - [Os](docs/Os.md)
 - [Parameter](docs/Parameter.md)
 - [ParameterDetails200Response](docs/ParameterDetails200Response.md)
 - [PromoteToRelay200Response](docs/PromoteToRelay200Response.md)
 - [PurgeSoftware200Response](docs/PurgeSoftware200Response.md)
 - [QueueInformation200Response](docs/QueueInformation200Response.md)
 - [QueueInformation200ResponseData](docs/QueueInformation200ResponseData.md)
 - [ReadCVEfromFS200Response](docs/ReadCVEfromFS200Response.md)
 - [RegeneratePolicies200Response](docs/RegeneratePolicies200Response.md)
 - [RegeneratePolicies200ResponseData](docs/RegeneratePolicies200ResponseData.md)
 - [ReloadAll200Response](docs/ReloadAll200Response.md)
 - [ReloadAll200ResponseData](docs/ReloadAll200ResponseData.md)
 - [ReloadAllDatasourcesAllNodes200Response](docs/ReloadAllDatasourcesAllNodes200Response.md)
 - [ReloadAllDatasourcesOneNode200Response](docs/ReloadAllDatasourcesOneNode200Response.md)
 - [ReloadGroup200Response](docs/ReloadGroup200Response.md)
 - [ReloadGroups200Response](docs/ReloadGroups200Response.md)
 - [ReloadGroups200ResponseData](docs/ReloadGroups200ResponseData.md)
 - [ReloadOneDatasourceAllNodes200Response](docs/ReloadOneDatasourceAllNodes200Response.md)
 - [ReloadOneDatasourceOneNode200Response](docs/ReloadOneDatasourceOneNode200Response.md)
 - [ReloadTechniques200Response](docs/ReloadTechniques200Response.md)
 - [ReloadTechniques200ResponseData](docs/ReloadTechniques200ResponseData.md)
 - [ReloadUserConf200Response](docs/ReloadUserConf200Response.md)
 - [ReloadUserConf200ResponseData](docs/ReloadUserConf200ResponseData.md)
 - [ReloadUserConf200ResponseDataReload](docs/ReloadUserConf200ResponseDataReload.md)
 - [RemoveValidatedUser200Response](docs/RemoveValidatedUser200Response.md)
 - [RestoreArchive200Response](docs/RestoreArchive200Response.md)
 - [RestoreArchive200ResponseData](docs/RestoreArchive200ResponseData.md)
 - [Rule](docs/Rule.md)
 - [RuleCategory](docs/RuleCategory.md)
 - [RuleCategoryUpdate](docs/RuleCategoryUpdate.md)
 - [RuleComplianceComponent](docs/RuleComplianceComponent.md)
 - [RuleComplianceComponentComplianceDetails](docs/RuleComplianceComponentComplianceDetails.md)
 - [RuleComplianceComponentComponentsInner](docs/RuleComplianceComponentComponentsInner.md)
 - [RuleComplianceComponentComponentsInnerValuesInner](docs/RuleComplianceComponentComponentsInnerValuesInner.md)
 - [RuleComplianceComponentComponentsInnerValuesInnerReportsInner](docs/RuleComplianceComponentComponentsInnerValuesInnerReportsInner.md)
 - [RuleDetails200Response](docs/RuleDetails200Response.md)
 - [RuleNew](docs/RuleNew.md)
 - [RuleStatus](docs/RuleStatus.md)
 - [RuleTarget](docs/RuleTarget.md)
 - [RuleTargetsInner](docs/RuleTargetsInner.md)
 - [RuleTargetsInnerExclude](docs/RuleTargetsInnerExclude.md)
 - [RuleTargetsInnerInclude](docs/RuleTargetsInnerInclude.md)
 - [RuleWithCategory](docs/RuleWithCategory.md)
 - [SaveCampaign200Response](docs/SaveCampaign200Response.md)
 - [SaveCampaignEvent200Response](docs/SaveCampaignEvent200Response.md)
 - [SaveWorkflowUser200Response](docs/SaveWorkflowUser200Response.md)
 - [SaveWorkflowUserRequest](docs/SaveWorkflowUserRequest.md)
 - [ScheduleCampaign200Response](docs/ScheduleCampaign200Response.md)
 - [Secrets](docs/Secrets.md)
 - [SetAllowedNetworks200Response](docs/SetAllowedNetworks200Response.md)
 - [SetAllowedNetworks200ResponseData](docs/SetAllowedNetworks200ResponseData.md)
 - [SetAllowedNetworksRequest](docs/SetAllowedNetworksRequest.md)
 - [TechniqueBlock](docs/TechniqueBlock.md)
 - [TechniqueBlockCallsInner](docs/TechniqueBlockCallsInner.md)
 - [TechniqueBlockReportingLogic](docs/TechniqueBlockReportingLogic.md)
 - [TechniqueCategories200Response](docs/TechniqueCategories200Response.md)
 - [TechniqueCategories200ResponseData](docs/TechniqueCategories200ResponseData.md)
 - [TechniqueCategory](docs/TechniqueCategory.md)
 - [TechniqueMethodCall](docs/TechniqueMethodCall.md)
 - [TechniqueMethodCallParametersInner](docs/TechniqueMethodCallParametersInner.md)
 - [TechniqueParameter](docs/TechniqueParameter.md)
 - [TechniqueResource](docs/TechniqueResource.md)
 - [TechniqueRevisions200Response](docs/TechniqueRevisions200Response.md)
 - [TechniqueRevisions200ResponseData](docs/TechniqueRevisions200ResponseData.md)
 - [TechniquesResourcesInner](docs/TechniquesResourcesInner.md)
 - [TechniquesRevisionsInner](docs/TechniquesRevisionsInner.md)
 - [TechniquesVersionsInner](docs/TechniquesVersionsInner.md)
 - [Timezone](docs/Timezone.md)
 - [UpdateBRandingConf200Response](docs/UpdateBRandingConf200Response.md)
 - [UpdateCVE200Response](docs/UpdateCVE200Response.md)
 - [UpdateCVE200ResponseData](docs/UpdateCVE200ResponseData.md)
 - [UpdateCVECheckConfiguration200Response](docs/UpdateCVECheckConfiguration200Response.md)
 - [UpdateCVECheckConfigurationRequest](docs/UpdateCVECheckConfigurationRequest.md)
 - [UpdateCVERequest](docs/UpdateCVERequest.md)
 - [UpdateChangeRequest200Response](docs/UpdateChangeRequest200Response.md)
 - [UpdateChangeRequestRequest](docs/UpdateChangeRequestRequest.md)
 - [UpdateDataSource200Response](docs/UpdateDataSource200Response.md)
 - [UpdateDirective200Response](docs/UpdateDirective200Response.md)
 - [UpdateGroup200Response](docs/UpdateGroup200Response.md)
 - [UpdateGroupCategory200Response](docs/UpdateGroupCategory200Response.md)
 - [UpdateNode200Response](docs/UpdateNode200Response.md)
 - [UpdateParameter200Response](docs/UpdateParameter200Response.md)
 - [UpdatePolicies200Response](docs/UpdatePolicies200Response.md)
 - [UpdateRule200Response](docs/UpdateRule200Response.md)
 - [UpdateRule200ResponseData](docs/UpdateRule200ResponseData.md)
 - [UpdateRuleCategory200Response](docs/UpdateRuleCategory200Response.md)
 - [UpdateSecret200Response](docs/UpdateSecret200Response.md)
 - [UpdateUser200Response](docs/UpdateUser200Response.md)
 - [UpdateUser200ResponseData](docs/UpdateUser200ResponseData.md)
 - [UpdateUser200ResponseDataUpdatedUser](docs/UpdateUser200ResponseDataUpdatedUser.md)
 - [UploadInventory200Response](docs/UploadInventory200Response.md)
 - [Users](docs/Users.md)
 - [ValidatedUser](docs/ValidatedUser.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### API-Tokens

- **Type**: API key
- **API key parameter name**: X-API-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-API-Token and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		rudder-golang.ContextAPIKeys,
		map[string]rudder-golang.APIKey{
			"X-API-Token": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

dev@rudder.io

