package services

import (
	"ecommerce/helpers"
	"ecommerce/models/mysql"
)

type StorefrontRepository interface {
	Create(record *mysql.Storefront) (*mysql.Storefront, error)
	GetByPrimaryKey(id uint64, columns []string) (*mysql.Storefront, error)
	GetByUniqueKey(conditions map[string]any, columns []string) (*mysql.Storefront, error)
	UpdateById(id uint64, record *mysql.Storefront) error
}

type storefrontService struct {
	repository StorefrontRepository
}

func NewStorefrontService(repository StorefrontRepository) *storefrontService {
	return &storefrontService{repository: repository}
}

func (service *storefrontService) validateCreate(request *mysql.StorefrontCreation) error {
	storefrontName := helpers.TrimSpace(request.StorefrontName)
	marketplaceStorefrontId := helpers.TrimSpace(request.MarketplaceStorefrontId)
	countryCode := helpers.TrimSpace(request.CountryCode)
	cookie := helpers.TrimSpace(*request.Cookie)

	if err := helpers.ValidateCannotBeEmptyAndExceedCharacter("Storefront name", storefrontName, 255); err != nil {
		return err
	}

	if err := helpers.ValidateCannotBeEmptyAndExceedCharacter("Marketplace storefront id", marketplaceStorefrontId, 100); err != nil {
		return err
	}

	if err := helpers.ValidateCannotBeEmptyAndExceedCharacter("Country code", countryCode, 100); err != nil {
		return err
	}

	if err := helpers.ValidateCannotBeEmpty("Cookie", cookie); err != nil {
		return err
	}

	conditions := map[string]any{"marketplace_storefront_id": marketplaceStorefrontId}
	columns := []string{"id"}
	_, errDb := service.repository.GetByUniqueKey(conditions, columns)

	if err := helpers.ValidateDuplicatedKey(errDb); err != nil {
		return err
	}

	return nil
}

func (service *storefrontService) Create(request *mysql.StorefrontCreation) (*mysql.Storefront, error) {
	if err := service.validateCreate(request); err != nil {
		return nil, err
	}

	cookie := helpers.TrimSpace(*request.Cookie)

	data := &mysql.Storefront{
		CommonStorefront: mysql.CommonStorefront{
			StorefrontName:          helpers.TrimSpace(request.StorefrontName),
			MarketplaceStorefrontId: helpers.TrimSpace(request.MarketplaceStorefrontId),
			IsActive:                request.IsActive,
			Balance:                 request.Balance,
			CountryCode:             helpers.TrimSpace(request.CountryCode),
			Cookie:                  &cookie,
		},
	}

	record, err := service.repository.Create(data)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (service *storefrontService) GetByPrimaryKey(id uint64) (*mysql.Storefront, error) {
	record, err := service.repository.GetByPrimaryKey(id, nil)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (service *storefrontService) validateUpdate(id uint64, request *mysql.StorefrontUpdate) error {
	if request.IsEmpty() {
		return mysql.ErrNoFieldsToUpdate
	}

	if request.StorefrontName != nil {
		storefrontName := helpers.TrimSpace(*request.StorefrontName)
		if err := helpers.ValidateCannotBeEmptyAndExceedCharacter("Storefront name", storefrontName, 255); err != nil {
			return err
		}
	}

	if request.CountryCode != nil {
		countryCode := helpers.TrimSpace(*request.CountryCode)
		if err := helpers.ValidateCannotBeEmptyAndExceedCharacter("Country code", countryCode, 100); err != nil {
			return err
		}
	}

	if request.Cookie != nil {
		cookie := helpers.TrimSpace(*request.Cookie)
		if err := helpers.ValidateCannotBeEmpty("Cookie", cookie); err != nil {
			return err
		}
	}

	if _, err := service.repository.GetByPrimaryKey(id, nil); err != nil {
		return err
	}

	return nil
}

func (service *storefrontService) buildUpdatedData(request *mysql.StorefrontUpdate) *mysql.Storefront {
	storefront := mysql.CommonStorefront{}

	if request.StorefrontName != nil {
		storefront.StorefrontName = helpers.TrimSpace(*request.StorefrontName)
	}

	if request.IsActive != nil {
		storefront.IsActive = *request.IsActive
	}

	if request.Balance.Valid {
		storefront.Balance = request.Balance
	}

	if request.CountryCode != nil {
		storefront.CountryCode = helpers.TrimSpace(*request.CountryCode)
	}

	if request.Cookie != nil {
		cookie := helpers.TrimSpace(*request.Cookie)
		storefront.Cookie = &cookie
	}

	return &mysql.Storefront{
		CommonStorefront: storefront,
	}
}

func (service *storefrontService) UpdateById(id uint64, request *mysql.StorefrontUpdate) error {
	if err := service.validateUpdate(id, request); err != nil {
		return err
	}

	if err := service.repository.UpdateById(id, service.buildUpdatedData(request)); err != nil {
		return err
	}

	return nil
}
