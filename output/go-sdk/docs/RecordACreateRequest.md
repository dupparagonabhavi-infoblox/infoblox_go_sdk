# RecordACreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Ipv4addr** | Pointer to [**IPv4Addr**](IPv4Addr.md) |  | [optional] 
**DdnsPrincipal** | Pointer to **bool** |  | [optional] 
**DdnsProtected** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Creator** | Pointer to **string** |  | [optional] 
**Disable** | Pointer to **bool** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] [readonly] 
**View** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**UseTtl** | Pointer to **bool** |  | [optional] 
**CloudInfo** | Pointer to [**RecordACreateRequestCloudInfo**](RecordACreateRequestCloudInfo.md) |  | [optional] 
**Extattrs** | Pointer to [**map[string]ExtensibleAttributesValue**](ExtensibleAttributesValue.md) |  | [optional] 

## Methods

### NewRecordACreateRequest

`func NewRecordACreateRequest() *RecordACreateRequest`

NewRecordACreateRequest instantiates a new RecordACreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordACreateRequestWithDefaults

`func NewRecordACreateRequestWithDefaults() *RecordACreateRequest`

NewRecordACreateRequestWithDefaults instantiates a new RecordACreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RecordACreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordACreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordACreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordACreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIpv4addr

`func (o *RecordACreateRequest) GetIpv4addr() IPv4Addr`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *RecordACreateRequest) GetIpv4addrOk() (*IPv4Addr, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *RecordACreateRequest) SetIpv4addr(v IPv4Addr)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *RecordACreateRequest) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *RecordACreateRequest) GetDdnsPrincipal() bool`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *RecordACreateRequest) GetDdnsPrincipalOk() (*bool, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *RecordACreateRequest) SetDdnsPrincipal(v bool)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *RecordACreateRequest) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *RecordACreateRequest) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *RecordACreateRequest) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *RecordACreateRequest) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *RecordACreateRequest) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetComment

`func (o *RecordACreateRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordACreateRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordACreateRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordACreateRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordACreateRequest) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordACreateRequest) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordACreateRequest) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordACreateRequest) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordACreateRequest) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordACreateRequest) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordACreateRequest) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordACreateRequest) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDisable

`func (o *RecordACreateRequest) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordACreateRequest) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordACreateRequest) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordACreateRequest) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetZone

`func (o *RecordACreateRequest) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordACreateRequest) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordACreateRequest) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordACreateRequest) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetView

`func (o *RecordACreateRequest) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordACreateRequest) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordACreateRequest) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordACreateRequest) HasView() bool`

HasView returns a boolean if a field has been set.

### GetTtl

`func (o *RecordACreateRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordACreateRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordACreateRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordACreateRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordACreateRequest) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordACreateRequest) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordACreateRequest) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordACreateRequest) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordACreateRequest) GetCloudInfo() RecordACreateRequestCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordACreateRequest) GetCloudInfoOk() (*RecordACreateRequestCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordACreateRequest) SetCloudInfo(v RecordACreateRequestCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordACreateRequest) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetExtattrs

`func (o *RecordACreateRequest) GetExtattrs() map[string]ExtensibleAttributesValue`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *RecordACreateRequest) GetExtattrsOk() (*map[string]ExtensibleAttributesValue, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *RecordACreateRequest) SetExtattrs(v map[string]ExtensibleAttributesValue)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *RecordACreateRequest) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


