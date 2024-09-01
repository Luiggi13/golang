package utils

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/skip2/go-qrcode"
)

const DEFAULTURL = "https://www.default-url.com"

func ValidateURL(urlString string) error {
	u, err := url.ParseRequestURI(urlString)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("invalid URL format: %s", urlString)
	}
	return nil
}

func GenerateQrCode(url string) string {
	png, errEncode := qrcode.Encode(url, qrcode.Highest, 1024)
	if errEncode != nil {
		defaultQrCode, _ := qrcode.Encode(DEFAULTURL, qrcode.Highest, 1024)
		encoded := base64.StdEncoding.EncodeToString(defaultQrCode)
		return encoded
	} else {
		encoded := base64.StdEncoding.EncodeToString(png)
		return encoded
	}
}
