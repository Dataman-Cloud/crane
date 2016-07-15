package util

const (
	// General status code
	OPERATION_SUCCESS      int = 0
	PARAMETER_ERROR        int = 10001
	ENGINE_OPERATION_ERROR int = 10002

	// Stack status code
	STACK_OPERATION_ERROR int = 11000

	// Service status code
	SERVICE_OPERATION_ERROR int = 12000

	// Network status code
	NETWORK_OPERATION_ERROR int = 13000
	NETWORK_IN_USE          int = 13001
	NETWORK_PRE_DEFINED     int = 13002

	// Volume status code
	VOLUME_OPERATION_ERROR int = 14000

	// Container status code
	CONTAINER_OPERATION_ERROR int = 15000

	// Node status code
	NODE_OPERATION_ERROR int = 16000
)
