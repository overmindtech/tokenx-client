# JetStreamLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consumer** | Pointer to **int64** |  | [optional] 
**DiskStorage** | Pointer to **int64** |  | [optional] 
**MaxBytesRequired** | Pointer to **bool** |  | [optional] 
**MemStorage** | Pointer to **int64** |  | [optional] 
**Streams** | Pointer to **int64** |  | [optional] 

## Methods

### NewJetStreamLimits

`func NewJetStreamLimits() *JetStreamLimits`

NewJetStreamLimits instantiates a new JetStreamLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJetStreamLimitsWithDefaults

`func NewJetStreamLimitsWithDefaults() *JetStreamLimits`

NewJetStreamLimitsWithDefaults instantiates a new JetStreamLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumer

`func (o *JetStreamLimits) GetConsumer() int64`

GetConsumer returns the Consumer field if non-nil, zero value otherwise.

### GetConsumerOk

`func (o *JetStreamLimits) GetConsumerOk() (*int64, bool)`

GetConsumerOk returns a tuple with the Consumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumer

`func (o *JetStreamLimits) SetConsumer(v int64)`

SetConsumer sets Consumer field to given value.

### HasConsumer

`func (o *JetStreamLimits) HasConsumer() bool`

HasConsumer returns a boolean if a field has been set.

### GetDiskStorage

`func (o *JetStreamLimits) GetDiskStorage() int64`

GetDiskStorage returns the DiskStorage field if non-nil, zero value otherwise.

### GetDiskStorageOk

`func (o *JetStreamLimits) GetDiskStorageOk() (*int64, bool)`

GetDiskStorageOk returns a tuple with the DiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorage

`func (o *JetStreamLimits) SetDiskStorage(v int64)`

SetDiskStorage sets DiskStorage field to given value.

### HasDiskStorage

`func (o *JetStreamLimits) HasDiskStorage() bool`

HasDiskStorage returns a boolean if a field has been set.

### GetMaxBytesRequired

`func (o *JetStreamLimits) GetMaxBytesRequired() bool`

GetMaxBytesRequired returns the MaxBytesRequired field if non-nil, zero value otherwise.

### GetMaxBytesRequiredOk

`func (o *JetStreamLimits) GetMaxBytesRequiredOk() (*bool, bool)`

GetMaxBytesRequiredOk returns a tuple with the MaxBytesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytesRequired

`func (o *JetStreamLimits) SetMaxBytesRequired(v bool)`

SetMaxBytesRequired sets MaxBytesRequired field to given value.

### HasMaxBytesRequired

`func (o *JetStreamLimits) HasMaxBytesRequired() bool`

HasMaxBytesRequired returns a boolean if a field has been set.

### GetMemStorage

`func (o *JetStreamLimits) GetMemStorage() int64`

GetMemStorage returns the MemStorage field if non-nil, zero value otherwise.

### GetMemStorageOk

`func (o *JetStreamLimits) GetMemStorageOk() (*int64, bool)`

GetMemStorageOk returns a tuple with the MemStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemStorage

`func (o *JetStreamLimits) SetMemStorage(v int64)`

SetMemStorage sets MemStorage field to given value.

### HasMemStorage

`func (o *JetStreamLimits) HasMemStorage() bool`

HasMemStorage returns a boolean if a field has been set.

### GetStreams

`func (o *JetStreamLimits) GetStreams() int64`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *JetStreamLimits) GetStreamsOk() (*int64, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *JetStreamLimits) SetStreams(v int64)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *JetStreamLimits) HasStreams() bool`

HasStreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


