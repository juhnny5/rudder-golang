# BrandingConf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayBar** | **bool** | Whether header bar is displayed or not | 
**DisplayLabel** | **bool** | Whether header bar&#39;s label is displayed or not | 
**LabelText** | **string** | The header bar&#39;s label title | 
**BarColor** | [**Color**](Color.md) |  | 
**LabelColor** | [**Color**](Color.md) |  | 
**WideLogo** | [**Logo**](Logo.md) |  | 
**SmallLogo** | [**Logo**](Logo.md) |  | 
**DisplayBarLogin** | **bool** | Whether header bar is displayed in login page or not | 
**DisplayMotd** | **bool** | Whether the message of the day is displayed in login page or not | 
**Motd** | **string** | Message of the day in login page | 

## Methods

### NewBrandingConf

`func NewBrandingConf(displayBar bool, displayLabel bool, labelText string, barColor Color, labelColor Color, wideLogo Logo, smallLogo Logo, displayBarLogin bool, displayMotd bool, motd string, ) *BrandingConf`

NewBrandingConf instantiates a new BrandingConf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingConfWithDefaults

`func NewBrandingConfWithDefaults() *BrandingConf`

NewBrandingConfWithDefaults instantiates a new BrandingConf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayBar

`func (o *BrandingConf) GetDisplayBar() bool`

GetDisplayBar returns the DisplayBar field if non-nil, zero value otherwise.

### GetDisplayBarOk

`func (o *BrandingConf) GetDisplayBarOk() (*bool, bool)`

GetDisplayBarOk returns a tuple with the DisplayBar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayBar

`func (o *BrandingConf) SetDisplayBar(v bool)`

SetDisplayBar sets DisplayBar field to given value.


### GetDisplayLabel

`func (o *BrandingConf) GetDisplayLabel() bool`

GetDisplayLabel returns the DisplayLabel field if non-nil, zero value otherwise.

### GetDisplayLabelOk

`func (o *BrandingConf) GetDisplayLabelOk() (*bool, bool)`

GetDisplayLabelOk returns a tuple with the DisplayLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLabel

`func (o *BrandingConf) SetDisplayLabel(v bool)`

SetDisplayLabel sets DisplayLabel field to given value.


### GetLabelText

`func (o *BrandingConf) GetLabelText() string`

GetLabelText returns the LabelText field if non-nil, zero value otherwise.

### GetLabelTextOk

`func (o *BrandingConf) GetLabelTextOk() (*string, bool)`

GetLabelTextOk returns a tuple with the LabelText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelText

`func (o *BrandingConf) SetLabelText(v string)`

SetLabelText sets LabelText field to given value.


### GetBarColor

`func (o *BrandingConf) GetBarColor() Color`

GetBarColor returns the BarColor field if non-nil, zero value otherwise.

### GetBarColorOk

`func (o *BrandingConf) GetBarColorOk() (*Color, bool)`

GetBarColorOk returns a tuple with the BarColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarColor

`func (o *BrandingConf) SetBarColor(v Color)`

SetBarColor sets BarColor field to given value.


### GetLabelColor

`func (o *BrandingConf) GetLabelColor() Color`

GetLabelColor returns the LabelColor field if non-nil, zero value otherwise.

### GetLabelColorOk

`func (o *BrandingConf) GetLabelColorOk() (*Color, bool)`

GetLabelColorOk returns a tuple with the LabelColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelColor

`func (o *BrandingConf) SetLabelColor(v Color)`

SetLabelColor sets LabelColor field to given value.


### GetWideLogo

`func (o *BrandingConf) GetWideLogo() Logo`

GetWideLogo returns the WideLogo field if non-nil, zero value otherwise.

### GetWideLogoOk

`func (o *BrandingConf) GetWideLogoOk() (*Logo, bool)`

GetWideLogoOk returns a tuple with the WideLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWideLogo

`func (o *BrandingConf) SetWideLogo(v Logo)`

SetWideLogo sets WideLogo field to given value.


### GetSmallLogo

`func (o *BrandingConf) GetSmallLogo() Logo`

GetSmallLogo returns the SmallLogo field if non-nil, zero value otherwise.

### GetSmallLogoOk

`func (o *BrandingConf) GetSmallLogoOk() (*Logo, bool)`

GetSmallLogoOk returns a tuple with the SmallLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallLogo

`func (o *BrandingConf) SetSmallLogo(v Logo)`

SetSmallLogo sets SmallLogo field to given value.


### GetDisplayBarLogin

`func (o *BrandingConf) GetDisplayBarLogin() bool`

GetDisplayBarLogin returns the DisplayBarLogin field if non-nil, zero value otherwise.

### GetDisplayBarLoginOk

`func (o *BrandingConf) GetDisplayBarLoginOk() (*bool, bool)`

GetDisplayBarLoginOk returns a tuple with the DisplayBarLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayBarLogin

`func (o *BrandingConf) SetDisplayBarLogin(v bool)`

SetDisplayBarLogin sets DisplayBarLogin field to given value.


### GetDisplayMotd

`func (o *BrandingConf) GetDisplayMotd() bool`

GetDisplayMotd returns the DisplayMotd field if non-nil, zero value otherwise.

### GetDisplayMotdOk

`func (o *BrandingConf) GetDisplayMotdOk() (*bool, bool)`

GetDisplayMotdOk returns a tuple with the DisplayMotd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMotd

`func (o *BrandingConf) SetDisplayMotd(v bool)`

SetDisplayMotd sets DisplayMotd field to given value.


### GetMotd

`func (o *BrandingConf) GetMotd() string`

GetMotd returns the Motd field if non-nil, zero value otherwise.

### GetMotdOk

`func (o *BrandingConf) GetMotdOk() (*string, bool)`

GetMotdOk returns a tuple with the Motd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotd

`func (o *BrandingConf) SetMotd(v string)`

SetMotd sets Motd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


