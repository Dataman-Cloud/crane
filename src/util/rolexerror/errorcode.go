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

	CodeNetworkPredefined = 13002
)
