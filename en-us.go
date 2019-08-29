package errorscode

// ErrorStruct - Error Default Struct
type ErrorStruct struct {
	Code     string `json:"code"`
	Kind     string `json:"kind"`
	HTTPCode int    `json:"http_code"`
	Message  string `json:"message"`
}

// GetError - Returns a Default Error Struct with code and message
func GetError(code string) ErrorStruct {
	return Errors[code]
}

// GetErrors - Returns errors colletion
func GetErrors() map[string]ErrorStruct {
	return Errors
}

// Errors - Store errors code and message
var Errors = map[string]ErrorStruct{
	// 10xx - General Server or Network issues
	"1000": ErrorStruct{"1000", "system-issue", 502, "Unknown request error."},
	"1001": ErrorStruct{"1001", "system-issue", 500, "Internal error. Please try again."},

	// 11xx - Request issues
	"1100": ErrorStruct{"1100", "request-issue", 422, "An unknown parameter was sent."},
	"1101": ErrorStruct{"1101", "request-issue", 403, "Invalid country."},
	"1102": ErrorStruct{"1102", "request-issue", 401, "Invalid 2FA code."},
	"1103": ErrorStruct{"1103", "request-issue", 401, "Invalid e-mail 2FA code."},
	"1104": ErrorStruct{"1104", "request-issue", 409, "CPF already exists."},
	"1105": ErrorStruct{"1105", "request-issue", 409, "Username already exists."},
	"1106": ErrorStruct{"1106", "request-issue", 401, "Invalid Login."},
	"1109": ErrorStruct{"1109", "request-issue", 429, "Too many attempts"},
}
