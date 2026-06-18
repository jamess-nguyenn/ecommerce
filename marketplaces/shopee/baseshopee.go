package main

import (
	"ecommerce/helpers"
	"ecommerce/lib"
	"fmt"
	"os"
)

const (
	//common
	spcCdsVer = "2"
	baseUrl   = "https://banhang.shopee.vn/api/"

	//storefront
	urlGetStorefrontInfo = "selleraccount/shop_info/"

	//category
	urlGetCategoryTree = "v3/category/get_category_tree"

	//product
	urlGetProductList = "v3/opt/mpsku/list/v2/get_product_list"
)

func main() {
	GetStorefront()
	//GetCategory()
	//GetProduct(1, 12)
}

func getCookie() string {
	data, err := os.ReadFile("marketplaces/shopee/cookie.txt")
	if err != nil {
		helpers.LogPanic("Could not read the file: " + err.Error())
		return ""
	}

	return string(data)
}

func GetStorefront() {
	queryParams := map[string]string{
		"SPC_CDS":     helpers.GenerateUUID(),
		"SPC_CDS_VER": spcCdsVer,
	}
	headers := map[string]string{
		"accept": "application/json",
		"cookie": getCookie(),
	}

	res := lib.NewHTTPRequest().
		SetQueryParams(queryParams).
		SetHeaders(headers).
		Get(baseUrl + urlGetStorefrontInfo)

	storefront, ok := res.Data.(map[string]any)
	if !ok {
		helpers.LogError("The data is wrong format")
		return
	}

	fmt.Println(string(helpers.MarshalValue(storefront)))
}

func GetCategory() {
	includeLevel := "3"
	version := "1.0.1"

	queryParams := map[string]string{
		"SPC_CDS":       helpers.GenerateUUID(),
		"SPC_CDS_VER":   spcCdsVer,
		"include_level": includeLevel,
		"version":       version,
	}
	headers := map[string]string{
		"accept": "application/json",
		"cookie": getCookie(),
	}

	res := lib.NewHTTPRequest().
		SetQueryParams(queryParams).
		SetHeaders(headers).
		Get(baseUrl + urlGetCategoryTree)

	dataMap, ok := res.Data.(map[string]any)
	if !ok {
		helpers.LogError("The data is wrong format")
		return
	}

	categories, ok := dataMap["list"].([]any)
	if !ok {
		helpers.LogError("Could not parse the category data")
		return
	}

	fmt.Println(string(helpers.MarshalValue(categories)))
}

func GetProduct(page, limit int) {
	listType := "live_all"
	needAds := "true"

	queryParams := map[string]string{
		"SPC_CDS":     os.Getenv("SHOPEE_SPC_CDS"),
		"SPC_CDS_VER": spcCdsVer,
		"list_type":   listType,
		"need_ads":    needAds,
		"page_number": helpers.IntToString(page),
		"page_size":   helpers.IntToString(limit),
	}

	headers := map[string]string{
		"accept": "application/json",
		"cookie": getCookie(),
	}

	res := lib.NewHTTPRequest().
		SetQueryParams(queryParams).
		SetHeaders(headers).
		Get(baseUrl + urlGetProductList)

	dataMap, ok := res.Data.(map[string]any)
	if !ok {
		helpers.LogError("The data is wrong format")
		return
	}

	//pageInfoData, ok := dataMap["page_info"].(map[string]any)
	//if !ok {
	//	helpers.LogError("Could not parse page info")
	//	return
	//}
	//
	//pageNumber, _ := pageInfoData["page_number"].(float64)
	//pageSize, _ := pageInfoData["page_size"].(float64)
	//total, _ := pageInfoData["total"].(float64)
	//
	//log.Printf("Page Number: %v, Page Size: %v, Total: %v", int(pageNumber), int(pageSize), int(total))

	products, ok := dataMap["products"].([]any)
	if !ok {
		helpers.LogError("Could not parse the product data")
		return
	}

	fmt.Println(string(helpers.MarshalValue(products)))
}
