package dockerclient

import "time"

const (
	// service running state format string
	TaskRunningState = "running"

	defaultNetworkDriver = "overlay"

	defaultHttpRequestTimeout = time.Second * 10

	// docker remote API version
	API_VERSION = "1.24"
)

const (
	labelNamespace    = "dm.reserved.stack.namespace"
	LabelRegistryAuth = "dm.reserved.registry.auth"
	labelNodeEndpoint = "dm.reserved.node.endpoint"
)

const (
	//Service error code
	CodeInvalidServiceNanoCPUs     = "503-11404"
	CodeInvalidServiceDelay        = "503-11405"
	CodeInvalidServiceWindow       = "503-11406"
	CodeInvalidServiceEndpoint     = "503-11407"
	CodeInvalidServicePlacement    = "503-11408"
	CodeInvalidServiceMemoryBytes  = "503-11409"
	CodeInvalidServiceUpdateConfig = "503-11410"
	CodeInvalidServiceSpec         = "503-11411"
	CodeInvalidServiceName         = "503-11412"

	// stack error code
	CodeInvalidStackName = "503-11502"
	CodeStackNotFound    = "404-11503"

	// node error code
	CodeErrorUpdateNodeMethod = "503-11302"
	CodeErrorNodeRole         = "503-11303"
	CodeErrorNodeAvailability = "503-11304"
	CodeGetNodeInfoError      = "503-11305"

	// network error code
	CodeNetworkPredefined          = "403-11206"
	CodeNetworkNotFound            = "404-11207"
	CodeNetworkOrContainerNotFound = "404-11208"
	CodeInvalidNetworkName         = "503-11209"

	//Container error code
	CodePatchContainerParamError      = "400-11002"
	CodePatchContainerMethodUndefined = "400-11003"
	CodeContainerNotFound             = "404-11006"
	CodeContainerAlreadyRunning       = "400-11007"
	CodeContainerNotRunning           = "400-11008"
	CodeInvalidImageName              = "503-11009"

	//Go docker client error code
	CodeConnToNodeError          = "503-11701"
	CodeGetNodeEndpointError     = "503-11702"
	CodeNodeEndpointIpMatchError = "503-11703"
	CodeVerifyNodeEnpointFailed  = "503-11704"

	//Volume error code
	CodeInvalidVolumeName = "503-11602"
)
