package stringutils

import "bytes"

// Concat combines strings efficiently
func Concat(strings ...string) string {
	var buffer bytes.Buffer
	for _, str := range strings {
		buffer.WriteString(str)
	}
	return buffer.String()
}
