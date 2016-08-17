package rolexerror

const (
	// Default ok code
	CodeOk = 0

	//Default error code
	CodeUndefined = 10001

	//Container error code
	CodeListContainerParamError        = 11001
	CodePatchContainerParamError       = 11002
	CodePatchContainerMethodUndefined  = 11003
	CodeDeleteContainerParamError      = 11004
	CodeDeleteContainerMethodUndefined = 11005
	CodeContainerNotFound              = 11006
	CodeContainerAlreadyRunning        = 11007
	CodeContainerNotRunning            = 11008
	CodeInvalidImageName               = 11009

	//Image error code
	CodeListImageParamError = 11101

	//Network error code
	CodeConnectNetworkParamError   = 11201
	CodeConnectNetworkMethodError  = 11202
	CodeCreateNetworkParamError    = 11203
	CodeInspectNetworkParamError   = 11204
	CodeListNetworkParamError      = 11205
	CodeNetworkPredefined          = 11206
	CodeNetworkNotFound            = 11207
	CodeNetworkOrContainerNotFound = 11208
	CodeInvalidNetworkName         = 11209

	//Node error code
	CodeUpdateNodeParamError  = 11301
	CodeErrorUpdateNodeMethod = 11302
	CodeErrorNodeRole         = 11303
	CodeErrorNodeAvailability = 11304
	CodeGetNodeInfoError      = 11305

	//Service error code
	CodeUpdateServiceParamError    = 11401
	CodeCreateServiceParamError    = 11402
	CodeScaleServiceParamError     = 11403
	CodeInvalidServiceNanoCPUs     = 11404
	CodeInvalidServiceDelay        = 11405
	CodeInvalidServiceWindow       = 11406
	CodeInvalidServiceEndpoint     = 11407
	CodeInvalidServicePlacement    = 11408
	CodeInvalidServiceMemoryBytes  = 11409
	CodeInvalidServiceUpdateConfig = 11410
	CodeInvalidServiceSpec         = 11411
	CodeInvalidServiceName         = 11412

	//Task error code
	CodeListTaskParamError = 11404

	//Stack error code
	CodeCreateStackParamError = 11501
	CodeInvalidStackName      = 11502
	CodeStackNotFound         = 11503

	//Volume error code
	CodeCreateVolumeParamError = 11601
	CodeInvalidVolumeName      = 11602

	//Go docker client error code
	CodeConnToNodeError          = 11701
	CodeGetNodeEndpointError     = 11702
	CodeNodeEndpointIpMatchError = 11703

	//Get config error code
	CodeGetConfigError = 11901

	//Account
	CodeInvalidGroupId                                  = 12001
	CodeAccountCreateParamError                         = 12002
	CodeAccountCreateAuthenticatorError                 = 12003
	CodeAccountGetAccountError                          = 12004
	CodeAccountGetAccountNotFoundError                  = 12005
	CodeAccountLoginParamError                          = 12006
	CodeAccountLoginFailedError                         = 12007
	CodeAccountLogoutError                              = 12008
	CodeAccountGroupAccountsGroupIdNotValidError        = 12009
	CodeAccountGroupAccountsNotFoundError               = 12010
	CodeAccountAccoutGroupsAccountIdNotValidError       = 12011
	CodeAccountAccoutGroupsNotFoundError                = 12012
	CodeAccountGetGroupGroupIdNotValidError             = 12013
	CodeAccountGetGroupGroupIdNotFoundError             = 12014
	CodeAccountListGroupNotFoundError                   = 12015
	CodeAccountAuthenticatorModificationNotAllowedError = 12016
	CodeAccountCreateGroupParamError                    = 12017
	CodeAccountCreateGroupFailedError                   = 12018
	CodeAccountUpdateGroupParamError                    = 12019
	CodeAccountUpdateGroupFailedError                   = 12020
	CodeAccountDeleteGroupGroupIdNotValidError          = 12021
	CodeAccountDeleteGroupFailedError                   = 12022
	CodeAccountJoinGroupGroupIdNotValidError            = 12023
	CodeAccountJoinGroupAccountIdNotValidError          = 12024
	CodeAccountJoinGroupFailedError                     = 12025
	CodeAccountLeaveGroupGroupIdNotValidError           = 12026
	CodeAccountLeaveGroupAccountIdNotValidError         = 12027
	CodeAccountLeaveGroupFailedError                    = 12028
	CodeAccountGrantServicePermissionParamError         = 12029
	CodeAccountGrantServicePermissionFailedError        = 12030
	CodeAccountRevokeServicePermissionParamError        = 12031
	CodeAccountRevokeServicePermissionFailedError       = 12032

	//Search
	CodeInvalidSearchKeywords = 13001

	//Registry
	CodeRegistryGetManifestError          = 14001
	CodeRegistryManifestParseError        = 14002
	CodeRegistryManifestDeleteError       = 14003
	CodeRegistryImagePublicityParamError  = 14004
	CodeRegistryImagePublicityUpdateError = 14005
	CodeRegistryCatalogListError          = 14006
	CodeRegistryTagsListError             = 14007

	//Catalog
	CodeCatalogGetCatalogError  = 15001
	CodeCatalogListCatalogError = 15002

	//license
	CodeLicenseGetLicenseError    = 16001
	CodeLicenseCreateLicenseError = 16002
)
