package catalog

import (
	"errors"

	log "github.com/Sirupsen/logrus"
)

func (catalogApi *CatalogApi) Save(catalog *Catalog) error {
	var catalogs []Catalog
	if err := catalogApi.DbClient.Where("name = ?", catalog.Name).Find(&catalogs).Error; err != nil {
		log.Errorf("get catalog error: %v", err)
		return err
	}

	if len(catalogs) > 0 {
		log.Warnf("catlog: %s already exist", catalog.Name)
		return errors.New("already exist")
	}

	if err := catalogApi.DbClient.Save(catalog).Error; err != nil {
		log.Errorf("save catalog error: %v", err)
		return err
	}

	return nil
}

func (catalogApi *CatalogApi) List() ([]Catalog, error) {
	var catalogs []Catalog
	err := catalogApi.DbClient.Find(&catalogs).Error
	return catalogs, err
}

func (catalogApi *CatalogApi) Get(catalogId uint64) (Catalog, error) {
	var catalog Catalog
	err := catalogApi.DbClient.Where("id = ?", catalogId).First(&catalog).Error
	return catalog, err
}
