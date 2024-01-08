# Color

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Red** | **float32** | percentage of red component | 
**Blue** | **float32** | percentage of blue component | 
**Green** | **float32** | percentage of green component | 
**Alpha** | **float32** | percentage of opacity | 

## Methods

### NewColor

`func NewColor(red float32, blue float32, green float32, alpha float32, ) *Color`

NewColor instantiates a new Color object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColorWithDefaults

`func NewColorWithDefaults() *Color`

NewColorWithDefaults instantiates a new Color object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRed

`func (o *Color) GetRed() float32`

GetRed returns the Red field if non-nil, zero value otherwise.

### GetRedOk

`func (o *Color) GetRedOk() (*float32, bool)`

GetRedOk returns a tuple with the Red field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRed

`func (o *Color) SetRed(v float32)`

SetRed sets Red field to given value.


### GetBlue

`func (o *Color) GetBlue() float32`

GetBlue returns the Blue field if non-nil, zero value otherwise.

### GetBlueOk

`func (o *Color) GetBlueOk() (*float32, bool)`

GetBlueOk returns a tuple with the Blue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlue

`func (o *Color) SetBlue(v float32)`

SetBlue sets Blue field to given value.


### GetGreen

`func (o *Color) GetGreen() float32`

GetGreen returns the Green field if non-nil, zero value otherwise.

### GetGreenOk

`func (o *Color) GetGreenOk() (*float32, bool)`

GetGreenOk returns a tuple with the Green field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreen

`func (o *Color) SetGreen(v float32)`

SetGreen sets Green field to given value.


### GetAlpha

`func (o *Color) GetAlpha() float32`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *Color) GetAlphaOk() (*float32, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *Color) SetAlpha(v float32)`

SetAlpha sets Alpha field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


