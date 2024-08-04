package lowos

import "golang.org/x/sys/windows"

func IsAdmin() bool {
	return windows.GetCurrentProcessToken().IsElevated()
}
