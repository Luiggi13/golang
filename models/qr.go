package models

type QrInput struct {
	URLString string  `json:"url" validate:"required,url"`
	UserId    *string `json:"userId"`
	Premium   *bool   `json:"premium"`
}

type QrCodeGenerated struct {
	Id         string `json:"id"`
	StatusCode int64  `json:"status_code"`
	QrCode     string `json:"qr_code"`
	Premium    bool   `json:"premium" validate:"required,premium"`
}

type BaseError struct {
	Method  string `json:"method"`
	Message string `json:"message"`
	Url     string `json:"url,omitempty"`
}

type CustomErrorQR CustomErrorQRElement

type CustomErrorQRElement struct {
	Status     string  `json:"status"`
	StatusCode int64   `json:"status_code"`
	Message    string  `json:"message"`
	Details    Details `json:"details"`
}

type Details struct {
	Method string `json:"method"`
	URL    string `json:"url,omitempty"`
}

func BadRequestError(ue BaseError) CustomErrorQR {

	details := Details{
		Method: ue.Method,
	}
	if ue.Url != "" {
		details.URL = ue.Url
	}
	return CustomErrorQR{
		Status:     "Bad Request",
		StatusCode: 400,
		Message:    ue.Message,
		Details:    details,
	}
}
