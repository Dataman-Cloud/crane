package catalog

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

type FakeDB struct {
	FilePath string
	Db       *gorm.DB
}

func SetupDB() *FakeDB {
	sqliteDbPath := filepath.Join(os.TempDir(), "test.db")
	db, _ := gorm.Open("sqlite3", sqliteDbPath)
	db.AutoMigrate(&Catalog{})
	return &FakeDB{FilePath: sqliteDbPath, Db: db}
}

func TearDownDB(fakeDb *FakeDB) {
	fakeDb.Db.Close()
	os.Remove(fakeDb.FilePath)
}

func MockCatalog() *Catalog {
	return &Catalog{
		ID:          1,
		Name:        "catalog-name",
		Bundle:      "bundle",
		Description: "Description",
		IconData:    "IconData",
		AccountId:   1,
		Type:        1,
	}
}

func TestCatalogSave(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	catalogApi := CatalogApi{DbClient: fakeDb.Db}
	err := fakeDb.Db.Create(MockCatalog()).Error
	if err != nil {
		panic(err)
	}

	assert.Nil(t, catalogApi.Save(MockCatalog()))
}

func TestCatalogList(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	catalogApi := CatalogApi{DbClient: fakeDb.Db}
	err := fakeDb.Db.Create(MockCatalog()).Error
	if err != nil {
		panic(err)
	}

	catalogs, err := catalogApi.List()
	assert.Nil(t, err)
	assert.Equal(t, len(catalogs), 1)
}

func TestCatalogGet(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	catalogApi := CatalogApi{DbClient: fakeDb.Db}
	err := fakeDb.Db.Create(MockCatalog()).Error
	if err != nil {
		panic(err)
	}

	var c Catalog
	fakeDb.Db.Last(&c)
	catalogGet, err := catalogApi.Get(c.ID)
	assert.Nil(t, err)
	assert.NotNil(t, catalogGet)
	assert.Equal(t, c.ID, catalogGet.ID)
}

func TestCatalogDelete(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	catalogApi := CatalogApi{DbClient: fakeDb.Db}
	err := fakeDb.Db.Create(MockCatalog()).Error
	if err != nil {
		panic(err)
	}

	var c Catalog
	fakeDb.Db.Last(&c)

	err = catalogApi.Delete(c.ID)
	assert.Nil(t, err)
}

func TestCatalogUpdate(t *testing.T) {
	fakeDb := SetupDB()
	defer TearDownDB(fakeDb)
	catalogApi := CatalogApi{DbClient: fakeDb.Db}
	err := fakeDb.Db.Create(MockCatalog()).Error
	if err != nil {
		panic(err)
	}

	assert.Nil(t, catalogApi.Update(MockCatalog()))
}
