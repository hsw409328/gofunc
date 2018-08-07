package go_php_func

import "os"

// SysGetTempDir - Returns directory path used for temporary files
func SysGetTempDir() string {

	return os.TempDir()
}
