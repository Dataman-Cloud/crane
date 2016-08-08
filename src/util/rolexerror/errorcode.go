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

	//Node error code
	CodeUpdateNodeParamError  = 11301
	CodeErrorUpdateNodeMethod = 11302
	CodeErrorNodeRole         = 11303
	CodeErrorNodeAvailability = 11304

	//Service error code
	CodeUpdateServiceParamError = 11401
	CodeCreateServiceParamError = 11402
	CodeScaleServiceParamError  = 11403

	//Task error code
	CodeListTaskParamError = 11404

	//Stack error code
	CodeCreateStackParamError = 11501

	//Volume error code
	CodeCreateVolumeParamError = 11601

	//Go docker client error code
	CodeGetDockerClientError = 1

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
	CodeAccountLeaveGroupGroupIdNotValidError           = 12023
	CodeAccountLeaveGroupAccountIdNotValidError         = 12024
	CodeAccountLeaveGroupFailedError                    = 12025
	CodeAccountGrantServicePermissionParamError         = 12026
	CodeAccountGrantServicePermissionFailedError        = 12027
	CodeAccountRevokeServicePermissionParamError        = 12028
	CodeAccountRevokeServicePermissionFailedError       = 12029

	//Search
	CodeInvalidSearchKeywords = 13001

	//Registry
	CodeRegistryGetManifestError          = 14001
	CodeRegistryManifestParseError        = 14002
	CodeRegistryManifestDeleteError       = 14003
	CodeRegistryImagePublicityParamError  = 14004
	CodeRegistryImagePublicityUpdateError = 14005
	CodeRegistryCatalogListError          = 14006

	//Catalog
	CodeCatalogGetCatalogError  = 15001
	CodeCatalogListCatalogError = 15002

	//license
	CodeLicenseGetLicenseError    = 16001
	CodeLicenseCreateLicenseError = 16002
)
