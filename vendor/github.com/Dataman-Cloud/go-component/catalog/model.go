package catalog

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Catalog struct {
	ID          uint64 `json:"Id"`
	Name        string `json:"Name"`
	Bundle      string `json:"Bundle" gorm:"size:65532"`
	Readme      string `json:"Readme" gorm:"size:65532"`
	Description string `json:"Description" gorm:"size:65532"`
	IconData    string `json:"IconData" gorm:"size:65532"`
	UserId      uint64 `json:"UserId"`
	Type        uint8  `json:"Type"`
}

const (
	ICON_SIZE    = 1024 * 1024 * 1024
	ICON_DEFAULT = "img/default.png"
)

const (
	CATALOG_SYSTEM_DEFAULT = 0
)

type Size interface {
	Size() int64
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
	if fs, err := os.Stat(filepath.Join(path, filepath.Base(path)+".png")); err != nil {
		return nil, err
	} else if fs.Size() > ICON_SIZE {
		return nil, errors.New("icon size too big")
	}

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

func ImageHandle(request *http.Request) (string, error) {
	var buf []byte
	var err error

	icon, _, err := request.FormFile("icon")
	if err != nil {
		if buf, err = ioutil.ReadFile(ICON_DEFAULT); err != nil {
			return "", err
		}
		return "", nil
	} else {
		if fileSize, ok := icon.(Size); !ok || fileSize.Size() > ICON_SIZE {
			return "", errors.New("invalid image")
		}

		buf, err = ioutil.ReadAll(icon)
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("data:%s;base64,%s",
		http.DetectContentType(buf),
		base64.StdEncoding.EncodeToString(buf)), nil
}
