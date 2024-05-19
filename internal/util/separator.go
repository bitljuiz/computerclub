package util

import "runtime"

// LineSeparator returns the line separator string based on the operating system.
// On Windows, it returns "\r\n". On other operating systems, it returns "\n".
func LineSeparator() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}
