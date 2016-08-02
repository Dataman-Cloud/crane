package registry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Dataman-Cloud/rolex/src/plugins/auth"
	"github.com/Dataman-Cloud/rolex/src/plugins/auth/authenticators"
	"github.com/Dataman-Cloud/rolex/src/util/config"
	"github.com/Dataman-Cloud/rolex/src/util/db"

	"github.com/docker/distribution/manifest/schema1"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

const manifestPattern = `^application/vnd.docker.distribution.manifest.v\d`

type Registry struct {
	Config   *config.Config
	DbClient *gorm.DB
}

func NewRegistry(Config *config.Config) *Registry {
	registry := &Registry{Config: Config, DbClient: db.DB()}
	return registry
}

func (registry *Registry) MigriateTable() {
	registry.DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Tag{})
	registry.DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&ImageAccess{})
}

func (registry *Registry) Token(ctx *gin.Context) {
	username, password, _ := ctx.Request.BasicAuth()
	authenticated := registry.Authenticate(username, password)

	service := ctx.Query("service")
	scope := ctx.Query("scope")

	if len(scope) == 0 && !authenticated {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	accesses := ParseResourceActions(scope)
	for _, access := range accesses {
		FilterAccess(username, authenticated, access)
	}

	//create token
	rawToken, err := MakeToken(registry.Config, username, service, accesses)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": rawToken})
}

func (registry *Registry) Authenticate(principal, password string) bool {
	var authenticator auth.Authenticator
	if registry.Config.AccountAuthenticator == "db" {
		authenticator = authenticators.NewDBAuthenticator()
	} else if registry.Config.AccountAuthenticator == "ldap" {
	} else {
		authenticator = authenticators.NewDefaultAuthenticator()
	}

	_, err := authenticator.Login(&auth.Account{Email: principal, Password: password})
	if err != nil {
		return false
	}
	return true
}

func (registry *Registry) Notifications(ctx *gin.Context) {
	notification := &Notification{}
	if err := ctx.BindJSON(&notification); err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			fmt.Printf("Notification JSON syntax error at byte %v: %s", jsonErr.Offset, jsonErr.Error())
		case *json.UnmarshalTypeError:
			fmt.Printf("Unexpected type at by type %v. Expected %s but received %s.",
				jsonErr.Offset, jsonErr.Type, jsonErr.Value)
		}
	}

	for _, e := range notification.Events {
		matched, _ := regexp.MatchString(manifestPattern, e.Target.MediaType)
		if matched && strings.HasPrefix(ctx.Request.UserAgent(), "Go-http-client") {
			registry.HandleNotification(e)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (registry *Registry) TagList(ctx *gin.Context) {
	var tags []Tag
	err := registry.DbClient.Where("namespace = ? AND image = ?", ctx.Param("namespace"), ctx.Param("image")).Find(&tags).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"code": 1, "data": "fail"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": tags})
}

func (registry *Registry) GetManifests(ctx *gin.Context) {
	account_, found := ctx.Get("account")
	if !found {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	account := account_.(auth.Account)

	resp, _, err := registry.RegistryAPIGet(fmt.Sprintf("%s/%s/manifests/%s", ctx.Param("namespace"), ctx.Param("image"), ctx.Param("reference")), account.Email)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	var manifest schema1.Manifest
	err = json.Unmarshal(resp, &manifest)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": manifest})
}

func (registry *Registry) DeleteManifests(ctx *gin.Context) {
	account_, found := ctx.Get("account")
	if !found {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	account := account_.(auth.Account)

	_, _, err := registry.RegistryAPIDeleteSchemaV2(fmt.Sprintf("%s/%s/manifests/%s", ctx.Param("namespace"), ctx.Param("image"), ctx.Param("reference")), account.Email)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
}

func (registry *Registry) MineCatalog(ctx *gin.Context) {
	account_, found := ctx.Get("account")
	if !found {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	account := account_.(auth.Account)

	var tags []Tag
	registry.DbClient.Exec("set sql_mode=''")
	err := registry.DbClient.Where("namespace = ?", account.ID).Order("created_at DESC").Group("image").Find(&tags).Error
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"code": 1, "data": "fail"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": tags})
}

func (registry *Registry) PublicCatalog(ctx *gin.Context) {
	var tags []Tag
	registry.DbClient.Exec("set sql_mode=''")
	err := registry.DbClient.Where("Publicity = 1").Order("created_at DESC").Group("namespace,image").Find(&tags).Error
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"code": 1, "data": "fail"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": tags})
}

func (registry *Registry) TagPublicity(ctx *gin.Context) {
	var param struct {
		Publicity uint8 `json:"Publicity"`
	}

	if err := ctx.BindJSON(&param); err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			fmt.Printf("Notification JSON syntax error at byte %v: %s", jsonErr.Offset, jsonErr.Error())
		case *json.UnmarshalTypeError:
			fmt.Printf("Unexpected type at by type %v. Expected %s but received %s.",
				jsonErr.Offset, jsonErr.Type, jsonErr.Value)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "data": "fail"})
		return
	}

	tag := Tag{}
	registry.DbClient.Where("namespace = ? AND image = ? AND tag = ?", ctx.Param("namespace"), ctx.Param("image"), ctx.Param("reference")).Find(&tag)
	err := registry.DbClient.Model(&tag).UpdateColumn("Publicity", param.Publicity).Error
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"code": 1, "data": "fail"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": "success"})
}

func (registry *Registry) HandleNotification(n Event) {
	if n.Action == "push" {
		if len(strings.Split(n.Target.Repository, "/")) != 2 {
			return
		}
		namespace := strings.Split(n.Target.Repository, "/")[0]
		image := strings.Split(n.Target.Repository, "/")[1]

		resp, _, err := registry.RegistryAPIGet(fmt.Sprintf("%s/%s/tags/list", namespace, image), n.Actor.Name)
		if err != nil {
			return
		}
		var respBody struct {
			Name string   `json:"name"`
			Tags []string `json:"tags"`
		}
		err = json.Unmarshal(resp, &respBody)
		if err != nil {
			return
		}

		for _, t := range respBody.Tags {
			tag := &Tag{}
			err = registry.DbClient.Where("namespace = ? AND image = ? AND tag = ? ", namespace, image, t).Find(tag).Error
			if err != nil && strings.Contains(err.Error(), "not found") {
				tag.Namespace = namespace
				tag.Image = image
				tag.Tag = t
				tag.Publicity = 0
				if namespace == "library" {
					tag.Publicity = 1
				}
				registry.DbClient.Save(tag)
			}

			tag.Size, tag.Digest, err = registry.SizeAndReferenceForTag(fmt.Sprintf("%s/%s/manifests/%s", namespace, image, t), n.Actor.Name)
			if tag.ID != 0 {
				registry.DbClient.Model(tag).Updates(Tag{Size: tag.Size, Digest: tag.Digest})
			} else {
				registry.DbClient.Save(tag)
			}
		}
	}
	registry.LogImageAccess(n)
}

func (registry *Registry) LogImageAccess(n Event) {
	ia := &ImageAccess{}
	if len(strings.Split(n.Target.Repository, "/")) == 2 {
		ia.Namespace = strings.Split(n.Target.Repository, "/")[0]
		ia.Image = strings.Split(n.Target.Repository, "/")[1]
		ia.Digest = n.Target.Digest
		ia.AccountEmail = n.Actor.Name
		ia.Action = n.Action
		registry.DbClient.Save(ia)
	}
}

func (registry *Registry) SizeAndReferenceForTag(url string, account string) (uint64, string, error) {
	var size uint64
	content, digest, err := registry.RegistryAPIGetSchemaV2(url, account)
	if err != nil {
		return 0, "", err
	}

	var v2response V2RegistryResponse
	err = json.Unmarshal(content, &v2response)
	if err != nil {
		return 0, "", err
	}

	size = v2response.Config.Size
	for _, layer := range v2response.Layers {
		size = size + layer.Size
	}

	return size, digest, nil
}
