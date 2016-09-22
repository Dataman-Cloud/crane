package dockerclient

import "time"

const (
	// service running state format string
	TaskRunningState = "running"

	defaultNetworkDriver = "overlay"

	defaultHttpRequestTimeout = time.Second * 10
)

const (
	labelNamespace    = "com.docker.stack.namespace"
	LabelRegistryAuth = "crane.reserved.registry.auth"
	labelNodeEndpoint = "crane.reserved.node.endpoint"
)

// sse event type
const (
	SseTypeContainerLogs  = "container-logs"
	SseTypeContainerStats = "container-stats"
	SseTypeServiceLogs    = "service-logs"
	SseTypeServiceStats   = "service-stats"
)

const (
	//Service error code
	CodeInvalidServiceNanoCPUs      = "503-11404"
	CodeInvalidServiceDelay         = "503-11405"
	CodeInvalidServiceWindow        = "503-11406"
	CodeInvalidServiceEndpoint      = "503-11407"
	CodeInvalidServicePlacement     = "503-11408"
	CodeInvalidServiceMemoryBytes   = "503-11409"
	CodeInvalidServiceUpdateConfig  = "503-11410"
	CodeInvalidServiceSpec          = "503-11411"
	CodeInvalidServiceName          = "503-11412"
	CodeGetServicePortConflictError = "503-11413"

	// stack error code
	CodeInvalidStackName = "503-11502"
	CodeStackUnavailable = "400-11503"

	// node error code
	CodeErrorUpdateNodeMethod     = "503-11302"
	CodeErrorNodeRole             = "503-11303"
	CodeErrorNodeAvailability     = "503-11304"
	CodeGetNodeInfoError          = "503-11305"
	CodeGetNodeAdvertiseAddrError = "503-11307"
	CodeJoinNodeError             = "503-11308"

	// network error code
	CodeNetworkPredefined         = "403-11206"
	CodeNetworkInvalid            = "400-11207"
	CodeNetworkOrContainerInvalid = "400-11208"
	CodeInvalidNetworkName        = "503-11209"

	//Container error code
	CodePatchContainerParamError      = "400-11002"
	CodePatchContainerMethodUndefined = "400-11003"
	CodeContainerInvalid              = "400-11006"
	CodeContainerAlreadyRunning       = "400-11007"
	CodeContainerNotRunning           = "400-11008"
	CodeInvalidImageName              = "503-11009"

	//Go docker client error code
	CodeConnToNodeError          = "503-11701"
	CodeGetNodeEndpointError     = "503-11702"
	CodeNodeEndpointIpMatchError = "503-11703"
	CodeVerifyNodeEnpointFailed  = "503-11704"
	CodeGetManagerInfoError      = "503-11705"

	//Volume error code
	CodeInvalidVolumeName = "503-11602"
)
