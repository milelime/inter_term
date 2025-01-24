package auth

import "net/http"

const (
	API_URL  = "https://example.com"
	IS_DEBUG = false
)

func CheckPasskey(passkey string) (validPasskey bool, error error) {
	if IS_DEBUG == true {
		return true, nil
	}
	resp, err := http.Get(API_URL + passkey)
	if err != nil {
		return false, err
	}
	return resp.Status == "200", nil
}
