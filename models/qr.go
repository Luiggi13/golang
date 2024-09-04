package models

// BadRequestError creates a CustomErrorQR with status "Bad Request" and status code 400
// from a given BaseError. It also populates the Details field of the CustomErrorQR
// with the Method and URL (if provided) from the BaseError.
//
// Parameters:
// - be: A BaseError struct containing the method and URL information.
//
// Returns:
// - A CustomErrorQR struct with the specified status, status code, message, and details.
func BadRequestError(be BaseError) CustomErrorQR {

	details := Details{
		Method: be.Method,
	}
	if be.Url != "" {
		details.URL = be.Url
	}
	return CustomErrorQR{
		Status:     "Bad Request",
		StatusCode: 400,
		Message:    be.Message,
		Details:    details,
	}
}
