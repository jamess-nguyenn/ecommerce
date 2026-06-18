package main

import (
	"ecommerce/lib"
)

func main() {
	CallCategoryNews()
}

func CallCategoryNews() {
	queryParams := map[string]string{
		"site_id":       "1000000",
		"category_id":   "1001005",
		"type":          "1",
		"limit":         "15",
		"data_select":   "article_id,article_type,title,share_url,thumbnail_url,publish_time,lead,privacy,original_cate,article_category",
		"thumb_size":    "300x180",
		"thumb_quality": "100",
		"thumb_dpr":     "1,2",
		"thumb_fit":     "crop",
		"exclude_id":    "4895811",
	}

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	lib.NewHTTPRequest().
		SetQueryParams(queryParams).
		SetHeaders(headers).
		Get("https://gw.vnexpress.net/mv")

}
