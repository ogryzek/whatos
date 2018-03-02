package whatos

import "runtime"

// TellMe tells me which OS I am using (as long as it is Mac OS or linux)
func TellMe() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "macOS"
	case "linux":
		return "linux"
	default:
		return "other"
	}
}
