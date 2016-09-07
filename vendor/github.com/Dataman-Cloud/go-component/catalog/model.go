package catalog

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type Catalog struct {
	ID          uint64 `json:"Id"`
	Name        string `json:"Name"`
	Bundle      string `json:"Bundle" gorm:"size:65532"`
	Readme      string `json:"Readme" gorm:"size:65532"`
	Description string `json:"Description" gorm:"size:65532"`
	IconData    string `json:"IconData" gorm:"size:65532"`
	IconType    string `json:"IconType"`
	UserId      uint64 `json:"UserId"`
	Type        uint8  `json:"Type"`
}

func CatalogFromPath(path string) (*Catalog, error) {
	bundle, err := ioutil.ReadFile(filepath.Join(path, "bundle.json"))
	if err != nil {
		return nil, err
	}

	readme, err := ioutil.ReadFile(filepath.Join(path, "readme.md"))
	if err != nil {
		return nil, err
	}

	description, err := ioutil.ReadFile(filepath.Join(path, "description"))
	if err != nil {
		return nil, err
	}

	f, err := ioutil.ReadFile(filepath.Join(path, filepath.Base(path)+".png"))
	if err != nil {
		return nil, err
	}

	catalog := &Catalog{
		Name:        filepath.Base(path),
		Bundle:      string(bundle),
		Readme:      string(readme),
		Description: string(description),
		IconData:    fmt.Sprintf("data:%s;base64,%s", http.DetectContentType(f), base64.StdEncoding.EncodeToString(f)),
	}
	return catalog, nil
}

func AllCatalogFromPath(path string) ([]*Catalog, error) {
	catalogs := make([]*Catalog, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return catalogs, err
	}

	for _, file := range files {
		if file.IsDir() {
			catalog, err := CatalogFromPath(filepath.Join(path, file.Name()))
			if err != nil {
				return catalogs, err
			}
			catalogs = append(catalogs, catalog)
		}
	}

	return catalogs, nil
}
