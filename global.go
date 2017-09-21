package validation

const (
	// Error code format: x.xxx.xxx?x
	// First 0 is used by system.
	// Next 001 - 100 is used by lapi.
	// Validation package is using 101 - 200
	// Last 3 digits for error code, it might be extended to
	// whatever number if necessary

	// Validator errors
	ERR_VALIDATOR_INVALID_TYPE = "0.101.001"
)
