package go_php_func

import "os/exec"

// Exec - Execute an external program
func Exec(s string) {

	exec.Command(s).Run()
}
