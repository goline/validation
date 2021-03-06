package validation

const (
	// Error code format: x.xxx.xxx?x
	// First 0 is used by system.
	// Next 001 - 100 is used by lapi.
	// Validation package is using 101 - 110
	// Last 3 digits for error code, it might be extended to
	// whatever number if necessary

	// Validator errors
	ERR_VALIDATOR_INVALID_TYPE     = "0.101.001"
	ERR_VALIDATOR_NOT_NIL          = "0.101.002"
	ERR_VALIDATOR_NOT_STRING       = "0.101.003"
	ERR_VALIDATOR_NOT_INT          = "0.101.004"
	ERR_VALIDATOR_NOT_FLOAT        = "0.101.005"
	ERR_VALIDATOR_NOT_NUMBER       = "0.101.006"
	ERR_VALIDATOR_INVALID_TAG      = "0.101.007"
	ERR_VALIDATOR_INVALID_ARGUMENT = "0.101.008"

	// Checkers errors
	ERR_VALIDATOR_UNKNOWN_ERROR        = "0.102.000"
	ERR_VALIDATOR_NOT_EMAIL            = "0.102.001"
	ERR_VALIDATOR_NOT_MIN              = "0.102.002"
	ERR_VALIDATOR_NOT_MAX              = "0.102.003"
	ERR_VALIDATOR_INVALID_FORMAT       = "0.102.004"
	ERR_VALIDATOR_NOT_IN_RANGE         = "0.102.005"
	ERR_VALIDATOR_NOT_MIN_LENGTH       = "0.102.006"
	ERR_VALIDATOR_NOT_MAX_LENGTH       = "0.102.007"
	ERR_VALIDATOR_REGEXP_WRONG_PATTERN = "0.102.008"
	ERR_VALIDATOR_REGEXP_NOT_MATCH     = "0.102.009"
	ERR_VALIDATOR_IN_EMPTY_LIST        = "0.102.010"
	ERR_VALIDATOR_NOT_IN_LIST          = "0.102.011"
	ERR_VALIDATOR_NOT_UNIQUE           = "0.102.012"
)
