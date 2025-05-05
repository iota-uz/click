# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Item/service name with units at the end | 
**Barcode** | Pointer to **string** | Barcode | [optional] 
**Labels** | Pointer to **[]string** | Array of marking codes | [optional] 
**SPIC** | **string** | SPIC code | 
**Units** | **int64** | Unit code | 
**PackageCode** | Pointer to **string** | Package Code | [optional] 
**GoodPrice** | **int64** | Price of one product/service unit | 
**Price** | Pointer to **int64** | The total amount of the item, including quantity, excluding discounts | [optional] 
**Amount** | Pointer to **int64** | Quantity | [optional] 
**VAT** | Pointer to **int64** | Amount of VAT in tiyins | [optional] 
**VATPercent** | Pointer to **int32** | VAT percentage | [optional] 
**Discount** | Pointer to **int64** | Discount | [optional] 
**Other** | Pointer to **int64** | Other discounts | [optional] 
**CommissionInfo** | Pointer to [**CommissionInfo**](CommissionInfo.md) |  | [optional] 

## Methods

### NewItem

`func NewItem(name string, sPIC string, units int64, goodPrice int64, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Item) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Item) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Item) SetName(v string)`

SetName sets Name field to given value.


### GetBarcode

`func (o *Item) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *Item) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *Item) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *Item) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetLabels

`func (o *Item) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Item) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Item) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Item) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSPIC

`func (o *Item) GetSPIC() string`

GetSPIC returns the SPIC field if non-nil, zero value otherwise.

### GetSPICOk

`func (o *Item) GetSPICOk() (*string, bool)`

GetSPICOk returns a tuple with the SPIC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPIC

`func (o *Item) SetSPIC(v string)`

SetSPIC sets SPIC field to given value.


### GetUnits

`func (o *Item) GetUnits() int64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Item) GetUnitsOk() (*int64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Item) SetUnits(v int64)`

SetUnits sets Units field to given value.


### GetPackageCode

`func (o *Item) GetPackageCode() string`

GetPackageCode returns the PackageCode field if non-nil, zero value otherwise.

### GetPackageCodeOk

`func (o *Item) GetPackageCodeOk() (*string, bool)`

GetPackageCodeOk returns a tuple with the PackageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageCode

`func (o *Item) SetPackageCode(v string)`

SetPackageCode sets PackageCode field to given value.

### HasPackageCode

`func (o *Item) HasPackageCode() bool`

HasPackageCode returns a boolean if a field has been set.

### GetGoodPrice

`func (o *Item) GetGoodPrice() int64`

GetGoodPrice returns the GoodPrice field if non-nil, zero value otherwise.

### GetGoodPriceOk

`func (o *Item) GetGoodPriceOk() (*int64, bool)`

GetGoodPriceOk returns a tuple with the GoodPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodPrice

`func (o *Item) SetGoodPrice(v int64)`

SetGoodPrice sets GoodPrice field to given value.


### GetPrice

`func (o *Item) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Item) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Item) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Item) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAmount

`func (o *Item) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Item) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Item) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Item) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetVAT

`func (o *Item) GetVAT() int64`

GetVAT returns the VAT field if non-nil, zero value otherwise.

### GetVATOk

`func (o *Item) GetVATOk() (*int64, bool)`

GetVATOk returns a tuple with the VAT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAT

`func (o *Item) SetVAT(v int64)`

SetVAT sets VAT field to given value.

### HasVAT

`func (o *Item) HasVAT() bool`

HasVAT returns a boolean if a field has been set.

### GetVATPercent

`func (o *Item) GetVATPercent() int32`

GetVATPercent returns the VATPercent field if non-nil, zero value otherwise.

### GetVATPercentOk

`func (o *Item) GetVATPercentOk() (*int32, bool)`

GetVATPercentOk returns a tuple with the VATPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVATPercent

`func (o *Item) SetVATPercent(v int32)`

SetVATPercent sets VATPercent field to given value.

### HasVATPercent

`func (o *Item) HasVATPercent() bool`

HasVATPercent returns a boolean if a field has been set.

### GetDiscount

`func (o *Item) GetDiscount() int64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Item) GetDiscountOk() (*int64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Item) SetDiscount(v int64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Item) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetOther

`func (o *Item) GetOther() int64`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *Item) GetOtherOk() (*int64, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *Item) SetOther(v int64)`

SetOther sets Other field to given value.

### HasOther

`func (o *Item) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetCommissionInfo

`func (o *Item) GetCommissionInfo() CommissionInfo`

GetCommissionInfo returns the CommissionInfo field if non-nil, zero value otherwise.

### GetCommissionInfoOk

`func (o *Item) GetCommissionInfoOk() (*CommissionInfo, bool)`

GetCommissionInfoOk returns a tuple with the CommissionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionInfo

`func (o *Item) SetCommissionInfo(v CommissionInfo)`

SetCommissionInfo sets CommissionInfo field to given value.

### HasCommissionInfo

`func (o *Item) HasCommissionInfo() bool`

HasCommissionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


