# RequestPostRequestInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | The HTTP method to use for the request | [optional] 
**Object** | Pointer to **string** | The object type to interact with | [optional] 
**Data** | Pointer to [**RequestPostRequestInnerData**](RequestPostRequestInnerData.md) |  | [optional] 
**AssignState** | Pointer to [**RequestPostRequestInnerAssignState**](RequestPostRequestInnerAssignState.md) |  | [optional] 
**Discard** | Pointer to **bool** | Whether to discard the response | [optional] 
**Args** | Pointer to [**RequestPostRequestInnerArgs**](RequestPostRequestInnerArgs.md) |  | [optional] 
**EnableSubstitution** | Pointer to **bool** | Whether to enable substitution in the request | [optional] 

## Methods

### NewRequestPostRequestInner

`func NewRequestPostRequestInner() *RequestPostRequestInner`

NewRequestPostRequestInner instantiates a new RequestPostRequestInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPostRequestInnerWithDefaults

`func NewRequestPostRequestInnerWithDefaults() *RequestPostRequestInner`

NewRequestPostRequestInnerWithDefaults instantiates a new RequestPostRequestInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *RequestPostRequestInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RequestPostRequestInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RequestPostRequestInner) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *RequestPostRequestInner) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetObject

`func (o *RequestPostRequestInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RequestPostRequestInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RequestPostRequestInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *RequestPostRequestInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *RequestPostRequestInner) GetData() RequestPostRequestInnerData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestPostRequestInner) GetDataOk() (*RequestPostRequestInnerData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestPostRequestInner) SetData(v RequestPostRequestInnerData)`

SetData sets Data field to given value.

### HasData

`func (o *RequestPostRequestInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetAssignState

`func (o *RequestPostRequestInner) GetAssignState() RequestPostRequestInnerAssignState`

GetAssignState returns the AssignState field if non-nil, zero value otherwise.

### GetAssignStateOk

`func (o *RequestPostRequestInner) GetAssignStateOk() (*RequestPostRequestInnerAssignState, bool)`

GetAssignStateOk returns a tuple with the AssignState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignState

`func (o *RequestPostRequestInner) SetAssignState(v RequestPostRequestInnerAssignState)`

SetAssignState sets AssignState field to given value.

### HasAssignState

`func (o *RequestPostRequestInner) HasAssignState() bool`

HasAssignState returns a boolean if a field has been set.

### GetDiscard

`func (o *RequestPostRequestInner) GetDiscard() bool`

GetDiscard returns the Discard field if non-nil, zero value otherwise.

### GetDiscardOk

`func (o *RequestPostRequestInner) GetDiscardOk() (*bool, bool)`

GetDiscardOk returns a tuple with the Discard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscard

`func (o *RequestPostRequestInner) SetDiscard(v bool)`

SetDiscard sets Discard field to given value.

### HasDiscard

`func (o *RequestPostRequestInner) HasDiscard() bool`

HasDiscard returns a boolean if a field has been set.

### GetArgs

`func (o *RequestPostRequestInner) GetArgs() RequestPostRequestInnerArgs`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *RequestPostRequestInner) GetArgsOk() (*RequestPostRequestInnerArgs, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *RequestPostRequestInner) SetArgs(v RequestPostRequestInnerArgs)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *RequestPostRequestInner) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetEnableSubstitution

`func (o *RequestPostRequestInner) GetEnableSubstitution() bool`

GetEnableSubstitution returns the EnableSubstitution field if non-nil, zero value otherwise.

### GetEnableSubstitutionOk

`func (o *RequestPostRequestInner) GetEnableSubstitutionOk() (*bool, bool)`

GetEnableSubstitutionOk returns a tuple with the EnableSubstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubstitution

`func (o *RequestPostRequestInner) SetEnableSubstitution(v bool)`

SetEnableSubstitution sets EnableSubstitution field to given value.

### HasEnableSubstitution

`func (o *RequestPostRequestInner) HasEnableSubstitution() bool`

HasEnableSubstitution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


