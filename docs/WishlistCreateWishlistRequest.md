# # WishlistCreateWishlistRequest


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId**| **string** |   |
**Privacy**| [**WishlistPrivacy**](WishlistPrivacy.md) |  for more information please, see Model/WishlistPrivacy.php  | [default to UNKNOWN]
**Label**| [**WishlistLocalizedText**](WishlistLocalizedText.md) |   | [optional]
**Description**| [**WishlistLocalizedText**](WishlistLocalizedText.md) |   | [optional]
**CustomerGrn**| **string** | If the customer GRN is set on JWT, it will be used as default. Otherwise, it will be used the customer_grn field.  | [optional]
**IsDefault**| **bool** |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

