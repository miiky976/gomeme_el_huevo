package utils

import "encoding/base64"

func Source(head string, image []byte) string {
	base := base64.StdEncoding.EncodeToString(image)
	return "data:" + head + ";base64," + base
}
