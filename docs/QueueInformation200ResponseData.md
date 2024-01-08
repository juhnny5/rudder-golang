# QueueInformation200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueMaxSize** | **int32** |  | 
**QueueSaturated** | **bool** | Is the inventory queue full | 

## Methods

### NewQueueInformation200ResponseData

`func NewQueueInformation200ResponseData(queueMaxSize int32, queueSaturated bool, ) *QueueInformation200ResponseData`

NewQueueInformation200ResponseData instantiates a new QueueInformation200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueInformation200ResponseDataWithDefaults

`func NewQueueInformation200ResponseDataWithDefaults() *QueueInformation200ResponseData`

NewQueueInformation200ResponseDataWithDefaults instantiates a new QueueInformation200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueMaxSize

`func (o *QueueInformation200ResponseData) GetQueueMaxSize() int32`

GetQueueMaxSize returns the QueueMaxSize field if non-nil, zero value otherwise.

### GetQueueMaxSizeOk

`func (o *QueueInformation200ResponseData) GetQueueMaxSizeOk() (*int32, bool)`

GetQueueMaxSizeOk returns a tuple with the QueueMaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMaxSize

`func (o *QueueInformation200ResponseData) SetQueueMaxSize(v int32)`

SetQueueMaxSize sets QueueMaxSize field to given value.


### GetQueueSaturated

`func (o *QueueInformation200ResponseData) GetQueueSaturated() bool`

GetQueueSaturated returns the QueueSaturated field if non-nil, zero value otherwise.

### GetQueueSaturatedOk

`func (o *QueueInformation200ResponseData) GetQueueSaturatedOk() (*bool, bool)`

GetQueueSaturatedOk returns a tuple with the QueueSaturated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSaturated

`func (o *QueueInformation200ResponseData) SetQueueSaturated(v bool)`

SetQueueSaturated sets QueueSaturated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


