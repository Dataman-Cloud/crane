package catalog

import (
	"io/ioutil"
	"path/filepath"
)

type Catalog struct {
	Name        string `json:"Name"`
	Bundle      string `json:"Bundle"`
	Icon        string `json:"Icon"`
	Readme      string `json:"Readme"`
	Description string `json:"Description"`
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

	catalog := &Catalog{
		Name:        filepath.Base(path),
		Bundle:      string(bundle),
		Readme:      string(readme),
		Description: string(description),
		Icon:        "/icons" + "/" + filepath.Base(path) + "/" + filepath.Base(path) + ".png",
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
