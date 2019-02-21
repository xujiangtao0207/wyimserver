package common

func GetBodyPrefix(bodyStr string) string {
	if bodyStr != "" {
		return bodyStr + "&"
	}
	return bodyStr
}
