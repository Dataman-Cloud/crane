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

	//Image error code
	CodeListImageParamError = 11101

	//Network error code
	CodeConnectNetworkParamError  = 11201
	CodeConnectNetworkMethodError = 11202
	CodeCreateNetworkParamError   = 11203
	CodeInspectNetworkParamError  = 11204
	CodeListNetworkParamError     = 11205
	CodeNetworkPredefined         = 11206

	//Node error code
	CodeUpdateNodeParamError = 11301

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

	//Group error code
	CodeInvalidGroupId = 12001
)
