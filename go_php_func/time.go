package go_php_func

import "time"

// Time - Return current Unix timestamp
func Time() int64 {

	return time.Now().Unix()
}
