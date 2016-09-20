package registry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/Dataman-Cloud/crane/src/plugins/auth"
	"github.com/Dataman-Cloud/crane/src/plugins/auth/authenticators"
	"github.com/Dataman-Cloud/crane/src/utils/cranerror"
	"github.com/Dataman-Cloud/crane/src/utils/db"
	"github.com/Dataman-Cloud/crane/src/utils/httpresponse"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/distribution/manifest/schema1"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

const (
	//Registry
	CodeRegistryGetManifestError          = "503-14001"
	CodeRegistryManifestParseError        = "503-14002"
	CodeRegistryManifestDeleteError       = "503-14003"
	CodeRegistryImagePublicityParamError  = "503-14004"
	CodeRegistryImagePublicityUpdateError = "503-14005"
	CodeRegistryCatalogListError          = "503-14006"
	CodeRegistryTagsListError             = "503-14007"
	CodeRegistryGetBasicAuthFaild         = "400-14008"
	CodeRegistryUnauthorized              = "401-14009"
	CodeRegistryMakeTokenFaild            = "503-14010"
)

const manifestPattern = `^application/vnd.docker.distribution.manifest.v\d`

type Registry struct {
	DbClient      *gorm.DB
	Authenticator auth.Authenticator

	AccountAuthenticator string
	PrivateKeyPath       string
	RegistryAddr         string
}

func NewRegistry(AccountAuthenticator string, PrivateKeyPath string, RegistryAddr string) *Registry {
	registry := &Registry{AccountAuthenticator: AccountAuthenticator,
		PrivateKeyPath: PrivateKeyPath,
		RegistryAddr:   RegistryAddr,
		DbClient:       db.DB()}

	if registry.AccountAuthenticator == "db" {
		registry.Authenticator = authenticators.NewDBAuthenticator()
	} else if registry.AccountAuthenticator == "ldap" {
	} else {
		registry.Authenticator = authenticators.NewDefaultAuthenticator()
	}

	registry.MigriateTable()
	return registry
}

func (registry *Registry) MigriateTable() {
	registry.DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Image{})
	registry.DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Tag{})
	registry.DbClient.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&ImageAccess{})
}

func (registry *Registry) Token(ctx *gin.Context) {
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		log.Error("registry get token error: can't get basicauth from request")
		craneerr := cranerror.NewError(CodeRegistryGetBasicAuthFaild, "registry get token error: can't get basicauth from request")
		httpresponse.Error(ctx, craneerr)
		return
	}
	authenticated := registry.Authenticate(username, password)

	service := ctx.Query("service")
	scope := ctx.Query("scope")

	if len(scope) == 0 && !authenticated {
		craneerr := cranerror.NewError(CodeRegistryUnauthorized, "registry get token unauthorized")
		httpresponse.Error(ctx, craneerr)
		return
	}

	accesses := registry.ParseResourceActions(scope)
	for _, access := range accesses {
		registry.FilterAccess(username, authenticated, access)
	}

	//create token
	rawToken, err := registry.MakeToken(registry.PrivateKeyPath, username, service, accesses)
	if err != nil {
		log.Errorf("get registry token error: %v", err)
		craneerr := cranerror.NewError(CodeRegistryMakeTokenFaild, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	//fixed format
	ctx.JSON(http.StatusOK, gin.H{"token": rawToken})
}

func (registry *Registry) Authenticate(principal, password string) bool {
	_, err := registry.Authenticator.Login(&auth.Account{Email: principal, Password: password})
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
			log.Errorf("Notification JSON syntax error at byte %v: %s", jsonErr.Offset, jsonErr.Error())
		case *json.UnmarshalTypeError:
			log.Errorf("Unexpected type at by type %v. Expected %s but received %s.",
				jsonErr.Offset, jsonErr.Type, jsonErr.Value)
		}
	}

	for _, e := range notification.Events {
		matched, _ := regexp.MatchString(manifestPattern, e.Target.MediaType)
		if matched && strings.HasPrefix(ctx.Request.UserAgent(), "Go-http-client") {
			registry.HandleNotification(e)
		}
	}

	httpresponse.Ok(ctx, gin.H{})
}

func (registry *Registry) TagList(ctx *gin.Context) {
	var tags []*Tag
	err := registry.DbClient.Where("namespace = ? AND image = ?", ctx.Param("namespace"), ctx.Param("image")).Find(&tags).Error
	if err != nil {
		craneerr := cranerror.NewError(CodeRegistryTagsListError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	for _, tag := range tags {
		registry.DbClient.Model(&ImageAccess{}).Where("namespace = ? AND image = ? AND digest = ? AND action='pull'", tag.Namespace, tag.Image, tag.Digest).Count(&tag.PullCount)
		registry.DbClient.Model(&ImageAccess{}).Where("namespace = ? AND image = ? AND digest = ? AND action='push'", tag.Namespace, tag.Image, tag.Digest).Count(&tag.PushCount)
	}

	httpresponse.Ok(ctx, tags)
}

func (registry *Registry) GetManifests(ctx *gin.Context) {
	account_, found := ctx.Get("account")
	if !found {
		craneerr := cranerror.NewError(CodeRegistryUnauthorized, "invalid user")
		httpresponse.Error(ctx, craneerr)
		return
	}
	account := account_.(auth.Account)

	resp, _, err := registry.RegistryAPIGet(fmt.Sprintf("%s/%s/manifests/%s", ctx.Param("namespace"), ctx.Param("image"), ctx.Param("reference")), account.Email)
	if err != nil {
		craneerr := cranerror.NewError(CodeRegistryGetManifestError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	var manifest schema1.Manifest
	err = json.Unmarshal(resp, &manifest)
	if err != nil {
		craneerr := cranerror.NewError(CodeRegistryManifestParseError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	httpresponse.Ok(ctx, manifest)
}

func (registry *Registry) DeleteManifests(ctx *gin.Context) {
	account_, found := ctx.Get("account")
	if !found {
		craneerr := cranerror.NewError(CodeRegistryUnauthorized, "invalid user")
		httpresponse.Error(ctx, craneerr)
		return
	}
	account := account_.(auth.Account)

	_, _, err := registry.RegistryAPIDeleteSchemaV2(fmt.Sprintf("%s/%s/manifests/%s", ctx.Param("namespace"), ctx.Param("image"), ctx.Param("reference")), account.Email)
	if err != nil {
		err := cranerror.NewError(CodeRegistryManifestDeleteError, err.Error())
		httpresponse.Error(ctx, err)
		return
	}

	httpresponse.Ok(ctx, "success")
}

func (registry *Registry) MineRepositories(ctx *gin.Context) {
	keywords := ctx.Query("keywords")
	account_, found := ctx.Get("account")
	if !found {
		craneerr := cranerror.NewError(CodeRegistryUnauthorized, "invalid user")
		httpresponse.Error(ctx, craneerr)
		return
	}
	account := account_.(auth.Account)

	var images []*Image
	var err error
	if len(keywords) > 0 {
		err = registry.DbClient.Where("namespace = ? AND (namespace like ? OR image like ?)", RegistryNamespaceForAccount(account),
			LikeParam(keywords), LikeParam(keywords)).Order("created_at DESC").Find(&images).Error
	} else {
		err = registry.DbClient.Where("namespace = ?", RegistryNamespaceForAccount(account)).Order("created_at DESC").Find(&images).Error
	}
	if err != nil {
		craneerr := cranerror.NewError(CodeRegistryCatalogListError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	for _, image := range images {
		registry.DbClient.Model(&ImageAccess{}).Where("namespace = ? AND image = ? AND action='pull'", image.Namespace, image.Image).Count(&image.PullCount)
		registry.DbClient.Model(&ImageAccess{}).Where("namespace = ? AND image = ? AND action='push'", image.Namespace, image.Image).Count(&image.PushCount)
	}

	httpresponse.Ok(ctx, images)
}

func (registry *Registry) PublicRepositories(ctx *gin.Context) {
	keywords := ctx.Query("keywords")
	var images []*Image
	var err error
	if len(keywords) > 0 {
		err = registry.DbClient.Where("Publicity = 1 AND (namespace like ? OR image like ?)", LikeParam(keywords), LikeParam(keywords)).Order("created_at DESC").Find(&images).Error
	} else {
		err = registry.DbClient.Where("Publicity = 1").Order("created_at DESC").Find(&images).Error
	}
	if err != nil {
		craneerr := cranerror.NewError(CodeRegistryCatalogListError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	for _, image := range images {
		registry.DbClient.Model(&ImageAccess{}).Where("namespace = ? AND image = ? AND action='pull'", image.Namespace, image.Image).Count(&image.PullCount)
		registry.DbClient.Model(&ImageAccess{}).Where("namespace = ? AND image = ? AND action='push'", image.Namespace, image.Image).Count(&image.PushCount)
	}

	httpresponse.Ok(ctx, images)
}

func (registry *Registry) ImagePublicity(ctx *gin.Context) {
	var param struct {
		Publicity uint8 `json:"Publicity"`
	}

	if err := ctx.BindJSON(&param); err != nil {
		switch jsonErr := err.(type) {
		case *json.SyntaxError:
			log.Errorf("Notification JSON syntax error at byte %v: %s", jsonErr.Offset, jsonErr.Error())
		case *json.UnmarshalTypeError:
			log.Errorf("Unexpected type at by type %v. Expected %s but received %s.",
				jsonErr.Offset, jsonErr.Type, jsonErr.Value)
		}
		craneerr := cranerror.NewError(CodeRegistryImagePublicityParamError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	var image Image
	registry.DbClient.Where("namespace = ? AND image = ? ", ctx.Param("namespace"), ctx.Param("image")).Find(&image)
	err := registry.DbClient.Model(&image).UpdateColumn("Publicity", param.Publicity).Error
	if err != nil {
		craneerr := cranerror.NewError(CodeRegistryImagePublicityUpdateError, err.Error())
		httpresponse.Error(ctx, craneerr)
		return
	}

	httpresponse.Ok(ctx, "success")
}

func (registry *Registry) HandleNotification(n Event) {
	if n.Action == "push" {
		if len(strings.Split(n.Target.Repository, "/")) != 2 {
			return
		}
		namespace := strings.Split(n.Target.Repository, "/")[0]
		image := strings.Split(n.Target.Repository, "/")[1]

		// create or update an image
		var modelImage Image
		err := registry.DbClient.Where("namespace = ? AND image = ?", namespace, image).Find(&modelImage).Error
		if err != nil && strings.Contains(err.Error(), "not found") {
			modelImage.Namespace = namespace
			modelImage.Image = image
			if namespace == "library" {
				modelImage.Publicity = 1
			}
		}

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

		modelImage.LatestTag = respBody.Tags[0]
		if modelImage.ID != 0 {
			registry.DbClient.Model(&modelImage).Updates(Image{LatestTag: modelImage.LatestTag})
		} else {
			registry.DbClient.Save(&modelImage)
		}

		// create or update tag
		for _, t := range respBody.Tags {
			tag := &Tag{}
			err = registry.DbClient.Where("namespace = ? AND image = ? AND tag = ? ", namespace, image, t).Find(tag).Error
			if err != nil && strings.Contains(err.Error(), "not found") {
				tag.Namespace = namespace
				tag.Image = image
				tag.Tag = t
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

func LikeParam(like string) string {
	return "%" + like + "%"
}
