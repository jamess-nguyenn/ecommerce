package mysql

import "github.com/shopspring/decimal"

const (
	tableName = "storefronts"
)

type CommonStorefront struct {
	StorefrontName          string              `json:"storefront_name" gorm:"column:storefront_name"`
	MarketplaceStorefrontId string              `json:"marketplace_storefront_id" gorm:"column:marketplace_storefront_id"`
	IsActive                bool                `json:"is_active" gorm:"column:is_active"`
	Balance                 decimal.NullDecimal `json:"balance" gorm:"column:balance"`
	CountryCode             string              `json:"country_code" gorm:"column:country_code"`
	Cookie                  *string             `json:"cookie" gorm:"column:cookie"`
}

type Storefront struct {
	CommonColumn
	CommonStorefront
}

func (s *Storefront) TableName() string {
	return tableName
}

type StorefrontCreation struct {
	IDColumn
	CommonStorefront
}

func (s *StorefrontCreation) TableName() string {
	return tableName
}

type StorefrontUpdate struct {
	StorefrontName *string             `json:"storefront_name" gorm:"column:storefront_name"`
	IsActive       *bool               `json:"is_active" gorm:"column:is_active"`
	Balance        decimal.NullDecimal `json:"balance" gorm:"column:balance"`
	CountryCode    *string             `json:"country_code" gorm:"column:country_code"`
	Cookie         *string             `json:"cookie" gorm:"column:cookie"`
}

func (s *StorefrontUpdate) TableName() string {
	return tableName
}

func (s *StorefrontUpdate) IsEmpty() bool {
	return s.StorefrontName == nil &&
		s.IsActive == nil &&
		!s.Balance.Valid &&
		s.CountryCode == nil &&
		s.Cookie == nil
}

// filter storefront

//type Filter struct {
//	Status string `json:"status" form:"status"`
//}
