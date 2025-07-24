package util

import "strings"

func GetMajorVersion(version string) string {
	trimmedVersion := strings.TrimPrefix(version, "v")

	v := strings.Split(trimmedVersion, ".")
	if len(v) > 0 {
		return v[0]
	}
	return ""
}
