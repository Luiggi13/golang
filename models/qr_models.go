package models

type QrInput struct {
	URLString string  `json:"url" validate:"required,url"`
	UserId    *string `json:"userId"`
	Premium   *bool   `json:"premium"`
	IdTag     string  `json:"id_tag" validate:"required"`
}

type QrCodeGenerated struct {
	Id         string `json:"id"`
	StatusCode int64  `json:"status_code"`
	QrCode     string `json:"qr_code"`
	Premium    bool   `json:"premium" validate:"required,premium"`
	TagName    string `json:"tag_name" validate:"required"`
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

type QRStruct struct {
	QrCode  string  `json:"qr_code"`
	UserId  *string `json:"userid"`
	UrlText string  `json:"url_text"`
	IdTag   string  `json:"id_tag"`
	Premium bool    `json:"premium"`
}

// qrs.id, qrs.qr_code, qrs.userid, qrs.url_text, qrs.premium
type QRStructJoin struct {
	QrId    string `json:"id"`
	QrCode  string `json:"qr_code"`
	UserId  string `json:"userid"`
	UrlText string `json:"url_text"`
	Premium bool   `json:"premium"`
	TagName string `json:"tag_name"`
}

type SelectTags struct {
	TagId   int32  `json:"id"`
	TagName string `json:"name"`
	Public  bool   `json:"public"`
}

type PutTags struct {
	TagName string `json:"name"`
	Public  *bool  `json:"public"`
}
