package models

// BadRequestError creates a CustomErrorQR with status "Bad Request" and status code 400
// DeleteResponse creates a CustomErrorQR with status "Resource deleted" and status code 204
// InternalRequestError creates a CustomErrorQR with status "Internal server error" and status code 500
// NotFound creates a CustomErrorQR with status "Not found" and status code 404
// from a given BaseError. It also populates the Details field of the CustomErrorQR
// with the Method and URL (if provided) from the BaseError.
//
// Parameters:
// - be: A BaseError struct containing the method and URL information.
//
// Returns:
// - A CustomErrorQR struct with the specified status, status code, message, and details.
func DeleteResponse(be BaseError) CustomErrorQR {

	details := Details{
		Method: be.Method,
	}
	if be.Url != "" {
		details.URL = be.Url
	}
	return CustomErrorQR{
		Status:     "Resource deleted",
		StatusCode: 204,
		Message:    be.Message,
		Details:    details,
	}
}
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
func InternalRequestError(be BaseError) CustomErrorQR {

	details := Details{
		Method: be.Method,
	}
	if be.Url != "" {
		details.URL = be.Url
	}
	return CustomErrorQR{
		Status:     "Internal server error",
		StatusCode: 500,
		Message:    be.Message,
		Details:    details,
	}
}
func NotFound(be BaseError) CustomErrorQR {

	details := Details{
		Method: be.Method,
	}
	if be.Url != "" {
		details.URL = be.Url
	}
	return CustomErrorQR{
		Status:     "Not found",
		StatusCode: 404,
		Message:    be.Message,
		Details:    details,
	}
}

func MigrationInterface(be BaseError) CustomErrorQR {

	details := Details{
		Method: be.Method,
	}
	if be.Url != "" {
		details.URL = be.Url
	}
	return CustomErrorQR{
		Status:     "Migrations",
		StatusCode: 200,
		Message:    be.Message,
		Details:    details,
	}
}
