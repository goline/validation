package validation

const (
	// Error code format: x.xxx.xxx?x
	// First 0 is used by system.
	// Next 001 - 100 is used by lapi.
	// Validation package is using 101 - 110
	// Last 3 digits for error code, it might be extended to
	// whatever number if necessary

	// Validator errors
	ERR_VALIDATOR_INVALID_TYPE = "0.101.001"
	ERR_VALIDATOR_NOT_NIL      = "0.101.002"
	ERR_VALIDATOR_NOT_STRING   = "0.101.003"
	ERR_VALIDATOR_NOT_INT      = "0.101.004"
	ERR_VALIDATOR_NOT_FLOAT    = "0.101.005"
	ERR_VALIDATOR_NOT_NUMBER   = "0.101.006"

	// Checkers errors
	ERR_VALIDATOR_UNKNOWN_ERROR  = "0.102.000"
	ERR_VALIDATOR_NOT_EMAIL      = "0.102.001"
	ERR_VALIDATOR_NOT_MIN        = "0.102.002"
	ERR_VALIDATOR_NOT_MAX        = "0.102.003"
	ERR_VALIDATOR_INVALID_FORMAT = "0.102.004"
	ERR_VALIDATOR_NOT_IN_RANGE   = "0.102.005"
)
