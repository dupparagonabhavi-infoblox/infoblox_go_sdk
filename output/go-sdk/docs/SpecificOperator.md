# SpecificOperator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The value to filter by. | [optional] 
**Op** | Pointer to **string** | The operator to use for filtering. | [optional] 

## Methods

### NewSpecificOperator

`func NewSpecificOperator() *SpecificOperator`

NewSpecificOperator instantiates a new SpecificOperator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecificOperatorWithDefaults

`func NewSpecificOperatorWithDefaults() *SpecificOperator`

NewSpecificOperatorWithDefaults instantiates a new SpecificOperator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SpecificOperator) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SpecificOperator) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SpecificOperator) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SpecificOperator) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetOp

`func (o *SpecificOperator) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SpecificOperator) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SpecificOperator) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *SpecificOperator) HasOp() bool`

HasOp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


