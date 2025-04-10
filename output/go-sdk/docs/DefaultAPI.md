# \DefaultAPI

All URIs are relative to *https://172.28.82.7/wapi/v2.10.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordAGet**](DefaultAPI.md#RecordAGet) | **Get** /record:a | Get A Records
[**RecordAObjRefDelete**](DefaultAPI.md#RecordAObjRefDelete) | **Delete** /record:a/{obj_ref} | king
[**RecordAObjRefGet**](DefaultAPI.md#RecordAObjRefGet) | **Get** /record:a/{obj_ref} | hello
[**RecordAObjRefPost**](DefaultAPI.md#RecordAObjRefPost) | **Post** /record:a/{obj_ref} | king
[**RecordAObjRefPut**](DefaultAPI.md#RecordAObjRefPut) | **Put** /record:a/{obj_ref} | king
[**RecordAPost**](DefaultAPI.md#RecordAPost) | **Post** /record:a | to create  record
[**RequestPost**](DefaultAPI.md#RequestPost) | **Post** /request | Execute a series of network management requests



## RecordAGet

> RecordAGet(ctx).ReturnFields(returnFields).MaxResults(maxResults).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Name(name).Comment(comment).Execute()

Get A Records

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	returnFields := "name,ipv4addr" // string | Comma-separated fields to return (e.g., name, ipv4addr) (optional)
	maxResults := int32(56) // int32 | Give positive value to indicate records to be truncated (optional)
	returnFields2 := "comments" // string | Returns basic fields as default as well as non-basic fields when mentioned explicitly (optional)
	returnAsObject := int32(56) // int32 | If set to 1, returns result as object (optional)
	name := openapiclient._record_a_get_name_parameter{SpecificOperator: openapiclient.NewSpecificOperator()} // RecordAGetNameParameter | A filter object for regex searches. If you want to use a basic operator (e.g., '='), enter the value as a string.  For specific operators, use the object format with 'value' and 'op' fields.  (optional)
	comment := openapiclient._record_a_get_name_parameter{SpecificOperator: openapiclient.NewSpecificOperator()} // RecordAGetNameParameter | A filter object for regex searches. If you want to use a basic operator (e.g., '='), enter the value as a string.  For specific operators, use the object format with 'value' and 'op' fields.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RecordAGet(context.Background()).ReturnFields(returnFields).MaxResults(maxResults).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Name(name).Comment(comment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecordAGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordAGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnFields** | **string** | Comma-separated fields to return (e.g., name, ipv4addr) | 
 **maxResults** | **int32** | Give positive value to indicate records to be truncated | 
 **returnFields2** | **string** | Returns basic fields as default as well as non-basic fields when mentioned explicitly | 
 **returnAsObject** | **int32** | If set to 1, returns result as object | 
 **name** | [**RecordAGetNameParameter**](RecordAGetNameParameter.md) | A filter object for regex searches. If you want to use a basic operator (e.g., &#39;&#x3D;&#39;), enter the value as a string.  For specific operators, use the object format with &#39;value&#39; and &#39;op&#39; fields.  | 
 **comment** | [**RecordAGetNameParameter**](RecordAGetNameParameter.md) | A filter object for regex searches. If you want to use a basic operator (e.g., &#39;&#x3D;&#39;), enter the value as a string.  For specific operators, use the object format with &#39;value&#39; and &#39;op&#39; fields.  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordAObjRefDelete

> RecordAObjRefDelete(ctx, objRef).Execute()

king



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objRef := "objRef_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RecordAObjRefDelete(context.Background(), objRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecordAObjRefDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objRef** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordAObjRefDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordAObjRefGet

> RecordAObjRefGet(ctx, objRef).ReturnFields(returnFields).MaxResults(maxResults).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

hello



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objRef := "objRef_example" // string | 
	returnFields := "name,ipv4addr" // string | Comma-separated fields to return (e.g., name, ipv4addr) (optional)
	maxResults := int32(56) // int32 | Give positive value to indicate records to be truncated (optional)
	returnFields2 := "comments" // string | Returns basic fields as default as well as non-basic fields when mentioned explicitly (optional)
	returnAsObject := int32(56) // int32 | If set to 1, returns result as object (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RecordAObjRefGet(context.Background(), objRef).ReturnFields(returnFields).MaxResults(maxResults).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecordAObjRefGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objRef** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordAObjRefGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnFields** | **string** | Comma-separated fields to return (e.g., name, ipv4addr) | 
 **maxResults** | **int32** | Give positive value to indicate records to be truncated | 
 **returnFields2** | **string** | Returns basic fields as default as well as non-basic fields when mentioned explicitly | 
 **returnAsObject** | **int32** | If set to 1, returns result as object | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordAObjRefPost

> RecordAObjRefPost(ctx, objRef).ReturnFields(returnFields).Method(method).RecordACreateRequest(recordACreateRequest).Execute()

king



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objRef := "objRef_example" // string | 
	returnFields := "name,ipv4addr" // string | Comma-separated fields to return (e.g., name, ipv4addr) (optional)
	method := "method_example" // string | it can be used to override a methods property (optional)
	recordACreateRequest := *openapiclient.NewRecordACreateRequest() // RecordACreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RecordAObjRefPost(context.Background(), objRef).ReturnFields(returnFields).Method(method).RecordACreateRequest(recordACreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecordAObjRefPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objRef** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordAObjRefPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnFields** | **string** | Comma-separated fields to return (e.g., name, ipv4addr) | 
 **method** | **string** | it can be used to override a methods property | 
 **recordACreateRequest** | [**RecordACreateRequest**](RecordACreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordAObjRefPut

> RecordAObjRefPut(ctx, objRef).ReturnFields(returnFields).RecordACreateRequest(recordACreateRequest).Execute()

king



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objRef := "objRef_example" // string | 
	returnFields := "name,ipv4addr" // string | Comma-separated fields to return (e.g., name, ipv4addr) (optional)
	recordACreateRequest := *openapiclient.NewRecordACreateRequest() // RecordACreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RecordAObjRefPut(context.Background(), objRef).ReturnFields(returnFields).RecordACreateRequest(recordACreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecordAObjRefPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objRef** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordAObjRefPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnFields** | **string** | Comma-separated fields to return (e.g., name, ipv4addr) | 
 **recordACreateRequest** | [**RecordACreateRequest**](RecordACreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordAPost

> RecordAPost(ctx).Method(method).RecordACreateRequest(recordACreateRequest).Execute()

to create  record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	method := "method_example" // string | it can be used to override a methods property (optional)
	recordACreateRequest := *openapiclient.NewRecordACreateRequest() // RecordACreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RecordAPost(context.Background()).Method(method).RecordACreateRequest(recordACreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RecordAPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecordAPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **method** | **string** | it can be used to override a methods property | 
 **recordACreateRequest** | [**RecordACreateRequest**](RecordACreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestPost

> RequestPost200Response RequestPost(ctx).RequestPostRequestInner(requestPostRequestInner).Execute()

Execute a series of network management requests

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	requestPostRequestInner := []openapiclient.RequestPostRequestInner{*openapiclient.NewRequestPostRequestInner()} // []RequestPostRequestInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RequestPost(context.Background()).RequestPostRequestInner(requestPostRequestInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RequestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestPost`: RequestPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RequestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestPostRequestInner** | [**[]RequestPostRequestInner**](RequestPostRequestInner.md) |  | 

### Return type

[**RequestPost200Response**](RequestPost200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

