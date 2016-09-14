package catalog

import (
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
)

func (catalogApi *CatalogApi) Save(catalog *Catalog) error {
	if err := catalogApi.DbClient.Save(catalog).Error; err != nil {
		return cranerror.NewError(CodeCatalogGetCatalogError, err.Error())
	}

	return nil
}

func (catalogApi *CatalogApi) List() ([]Catalog, error) {
	var catalogs []Catalog
	if err := catalogApi.DbClient.Find(&catalogs).Error; err != nil {
		return catalogs, cranerror.NewError(CodeCatalogListCatalogError, err.Error())
	}

	return catalogs, nil
}

func (catalogApi *CatalogApi) Get(catalogId uint64) (Catalog, error) {
	var catalog Catalog
	if err := catalogApi.DbClient.Where("id = ?", catalogId).First(&catalog).Error; err != nil {
		return catalog, cranerror.NewError(CodeCatalogGetCatalogError, err.Error())
	}

	return catalog, nil
}

func (catalogApi *CatalogApi) Delete(catalogId uint64) error {
	if err := catalogApi.DbClient.Delete(&Catalog{ID: catalogId}).Error; err != nil {
		return cranerror.NewError(CodeCatalogInvalidCatalogId, err.Error())
	}

	return nil
}

func (catalogApi *CatalogApi) Update(catalog *Catalog) error {
	if err := catalogApi.DbClient.Save(catalog).Error; err != nil {
		return cranerror.NewError(CodeCatalogInvalidParam, err.Error())
	}

	return nil
}
