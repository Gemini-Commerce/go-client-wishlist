# GeminiCommerce\Wishlist\WishlistAPI

All URIs are relative to *https://wishlist.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WishlistAddItemToWishlist**](WishlistAPI.md#WishlistAddItemToWishlist) | **Post** /wishlist.Wishlist/AddItemToWishlist | 
[**WishlistAreItemsInWishlists**](WishlistAPI.md#WishlistAreItemsInWishlists) | **Post** /wishlist.Wishlist/AreItemsInWishlists | 
[**WishlistBulkCreateSharing**](WishlistAPI.md#WishlistBulkCreateSharing) | **Post** /wishlist.Wishlist/BulkCreateSharing | Sharing endpoints
[**WishlistBulkRemoveItemsFromWishlists**](WishlistAPI.md#WishlistBulkRemoveItemsFromWishlists) | **Post** /wishlist.Wishlist/BulkRemoveItemsFromWishlists | BulkRemoveItemsFromWishlists removes items from wishlists.
[**WishlistBulkRevokeSharing**](WishlistAPI.md#WishlistBulkRevokeSharing) | **Post** /wishlist.Wishlist/BulkRevokeSharing | 
[**WishlistCreateWishlist**](WishlistAPI.md#WishlistCreateWishlist) | **Post** /wishlist.Wishlist/CreateWishlist | 
[**WishlistDeleteWishlist**](WishlistAPI.md#WishlistDeleteWishlist) | **Post** /wishlist.Wishlist/DeleteWishlist | 
[**WishlistGetItemFromWishlist**](WishlistAPI.md#WishlistGetItemFromWishlist) | **Post** /wishlist.Wishlist/GetItemFromWishlist | 
[**WishlistGetWishlistById**](WishlistAPI.md#WishlistGetWishlistById) | **Post** /wishlist.Wishlist/GetWishlistById | 
[**WishlistGetWishlistBySharedCode**](WishlistAPI.md#WishlistGetWishlistBySharedCode) | **Post** /wishlist.Wishlist/GetWishlistBySharedCode | 
[**WishlistListWishlistItems**](WishlistAPI.md#WishlistListWishlistItems) | **Post** /wishlist.Wishlist/ListWishlistItems | 
[**WishlistListWishlists**](WishlistAPI.md#WishlistListWishlists) | **Post** /wishlist.Wishlist/ListWishlists | 
[**WishlistMergeWishlists**](WishlistAPI.md#WishlistMergeWishlists) | **Post** /wishlist.Wishlist/MergeWishlists | 
[**WishlistRemoveItemFromWishlist**](WishlistAPI.md#WishlistRemoveItemFromWishlist) | **Post** /wishlist.Wishlist/RemoveItemFromWishlist | 
[**WishlistUpdateItemInWishlist**](WishlistAPI.md#WishlistUpdateItemInWishlist) | **Post** /wishlist.Wishlist/UpdateItemInWishlist | 
[**WishlistUpdateWishlist**](WishlistAPI.md#WishlistUpdateWishlist) | **Post** /wishlist.Wishlist/UpdateWishlist | 



## WishlistAddItemToWishlist

> WishlistWishlistItemResponse WishlistAddItemToWishlist(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistAddItemToWishlistRequest("TenantId_example", "WishlistId_example", "ItemGrn_example") // WishlistAddItemToWishlistRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistAddItemToWishlist(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistAddItemToWishlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistAddItemToWishlist`: WishlistWishlistItemResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistAddItemToWishlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistAddItemToWishlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistAddItemToWishlistRequest**](WishlistAddItemToWishlistRequest.md) |  | 

### Return type

[**WishlistWishlistItemResponse**](WishlistWishlistItemResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistAreItemsInWishlists

> WishlistAreItemsInWishlistsResponse WishlistAreItemsInWishlists(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistAreItemsInWishlistsRequest("TenantId_example", []string{"ItemGrns_example"}) // WishlistAreItemsInWishlistsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistAreItemsInWishlists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistAreItemsInWishlists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistAreItemsInWishlists`: WishlistAreItemsInWishlistsResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistAreItemsInWishlists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistAreItemsInWishlistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistAreItemsInWishlistsRequest**](WishlistAreItemsInWishlistsRequest.md) |  | 

### Return type

[**WishlistAreItemsInWishlistsResponse**](WishlistAreItemsInWishlistsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistBulkCreateSharing

> WishlistBulkCreateSharingResponse WishlistBulkCreateSharing(ctx).Body(body).Execute()

Sharing endpoints

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistBulkCreateSharingRequest() // WishlistBulkCreateSharingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistBulkCreateSharing(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistBulkCreateSharing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistBulkCreateSharing`: WishlistBulkCreateSharingResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistBulkCreateSharing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistBulkCreateSharingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistBulkCreateSharingRequest**](WishlistBulkCreateSharingRequest.md) |  | 

### Return type

[**WishlistBulkCreateSharingResponse**](WishlistBulkCreateSharingResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistBulkRemoveItemsFromWishlists

> map[string]interface{} WishlistBulkRemoveItemsFromWishlists(ctx).Body(body).Execute()

BulkRemoveItemsFromWishlists removes items from wishlists.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistBulkRemoveItemsFromWishlistsRequest("TenantId_example", []string{"ItemGrns_example"}) // WishlistBulkRemoveItemsFromWishlistsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistBulkRemoveItemsFromWishlists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistBulkRemoveItemsFromWishlists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistBulkRemoveItemsFromWishlists`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistBulkRemoveItemsFromWishlists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistBulkRemoveItemsFromWishlistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistBulkRemoveItemsFromWishlistsRequest**](WishlistBulkRemoveItemsFromWishlistsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistBulkRevokeSharing

> map[string]interface{} WishlistBulkRevokeSharing(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistBulkRevokeSharingRequest() // WishlistBulkRevokeSharingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistBulkRevokeSharing(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistBulkRevokeSharing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistBulkRevokeSharing`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistBulkRevokeSharing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistBulkRevokeSharingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistBulkRevokeSharingRequest**](WishlistBulkRevokeSharingRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistCreateWishlist

> WishlistWishlistResponse WishlistCreateWishlist(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistCreateWishlistRequest("TenantId_example", openapiclient.wishlistPrivacy("PRIVACY_UNKNOWN")) // WishlistCreateWishlistRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistCreateWishlist(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistCreateWishlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistCreateWishlist`: WishlistWishlistResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistCreateWishlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistCreateWishlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistCreateWishlistRequest**](WishlistCreateWishlistRequest.md) |  | 

### Return type

[**WishlistWishlistResponse**](WishlistWishlistResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistDeleteWishlist

> map[string]interface{} WishlistDeleteWishlist(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistDeleteWishlistRequest("TenantId_example", "Id_example") // WishlistDeleteWishlistRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistDeleteWishlist(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistDeleteWishlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistDeleteWishlist`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistDeleteWishlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistDeleteWishlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistDeleteWishlistRequest**](WishlistDeleteWishlistRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistGetItemFromWishlist

> WishlistWishlistItemResponse WishlistGetItemFromWishlist(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistGetItemFromWishlistRequest("TenantId_example", "Id_example") // WishlistGetItemFromWishlistRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistGetItemFromWishlist(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistGetItemFromWishlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistGetItemFromWishlist`: WishlistWishlistItemResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistGetItemFromWishlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistGetItemFromWishlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistGetItemFromWishlistRequest**](WishlistGetItemFromWishlistRequest.md) |  | 

### Return type

[**WishlistWishlistItemResponse**](WishlistWishlistItemResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistGetWishlistById

> WishlistWishlistResponse WishlistGetWishlistById(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistGetWishlistByIdRequest("TenantId_example", "Id_example") // WishlistGetWishlistByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistGetWishlistById(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistGetWishlistById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistGetWishlistById`: WishlistWishlistResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistGetWishlistById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistGetWishlistByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistGetWishlistByIdRequest**](WishlistGetWishlistByIdRequest.md) |  | 

### Return type

[**WishlistWishlistResponse**](WishlistWishlistResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistGetWishlistBySharedCode

> WishlistWishlistResponse WishlistGetWishlistBySharedCode(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistGetWishlistBySharedCodeRequest("TenantId_example", "SharedCode_example") // WishlistGetWishlistBySharedCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistGetWishlistBySharedCode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistGetWishlistBySharedCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistGetWishlistBySharedCode`: WishlistWishlistResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistGetWishlistBySharedCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistGetWishlistBySharedCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistGetWishlistBySharedCodeRequest**](WishlistGetWishlistBySharedCodeRequest.md) |  | 

### Return type

[**WishlistWishlistResponse**](WishlistWishlistResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistListWishlistItems

> WishlistListWishlistItemsResponse WishlistListWishlistItems(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistListWishlistItemsRequest("TenantId_example", "WishlistId_example") // WishlistListWishlistItemsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistListWishlistItems(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistListWishlistItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistListWishlistItems`: WishlistListWishlistItemsResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistListWishlistItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistListWishlistItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistListWishlistItemsRequest**](WishlistListWishlistItemsRequest.md) |  | 

### Return type

[**WishlistListWishlistItemsResponse**](WishlistListWishlistItemsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistListWishlists

> WishlistListWishlistsResponse WishlistListWishlists(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistListWishlistsRequest("TenantId_example") // WishlistListWishlistsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistListWishlists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistListWishlists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistListWishlists`: WishlistListWishlistsResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistListWishlists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistListWishlistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistListWishlistsRequest**](WishlistListWishlistsRequest.md) |  | 

### Return type

[**WishlistListWishlistsResponse**](WishlistListWishlistsResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistMergeWishlists

> WishlistWishlistResponse WishlistMergeWishlists(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistMergeWishlistsRequest("TenantId_example", "StartingWishlistId_example", "CustomerGrn_example") // WishlistMergeWishlistsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistMergeWishlists(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistMergeWishlists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistMergeWishlists`: WishlistWishlistResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistMergeWishlists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistMergeWishlistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistMergeWishlistsRequest**](WishlistMergeWishlistsRequest.md) |  | 

### Return type

[**WishlistWishlistResponse**](WishlistWishlistResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistRemoveItemFromWishlist

> map[string]interface{} WishlistRemoveItemFromWishlist(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistRemoveItemFromWishlistRequest("TenantId_example", "Id_example") // WishlistRemoveItemFromWishlistRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistRemoveItemFromWishlist(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistRemoveItemFromWishlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistRemoveItemFromWishlist`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistRemoveItemFromWishlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistRemoveItemFromWishlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistRemoveItemFromWishlistRequest**](WishlistRemoveItemFromWishlistRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistUpdateItemInWishlist

> WishlistWishlistItemResponse WishlistUpdateItemInWishlist(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistUpdateItemInWishlistRequest("TenantId_example", "Id_example", *openapiclient.NewWishlistUpdateItemInWishlistRequestPayload(), "PayloadMask_example") // WishlistUpdateItemInWishlistRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistUpdateItemInWishlist(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistUpdateItemInWishlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistUpdateItemInWishlist`: WishlistWishlistItemResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistUpdateItemInWishlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistUpdateItemInWishlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistUpdateItemInWishlistRequest**](WishlistUpdateItemInWishlistRequest.md) |  | 

### Return type

[**WishlistWishlistItemResponse**](WishlistWishlistItemResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WishlistUpdateWishlist

> WishlistWishlistResponse WishlistUpdateWishlist(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-wishlist"
)

func main() {
	body := *openapiclient.NewWishlistUpdateWishlistRequest("TenantId_example", "Id_example", *openapiclient.NewWishlistUpdateWishlistRequestPayload(), "PayloadMask_example") // WishlistUpdateWishlistRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WishlistAPI.WishlistUpdateWishlist(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WishlistAPI.WishlistUpdateWishlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WishlistUpdateWishlist`: WishlistWishlistResponse
	fmt.Fprintf(os.Stdout, "Response from `WishlistAPI.WishlistUpdateWishlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWishlistUpdateWishlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WishlistUpdateWishlistRequest**](WishlistUpdateWishlistRequest.md) |  | 

### Return type

[**WishlistWishlistResponse**](WishlistWishlistResponse.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

