package go_php_func

// ArrayPush - Push one or more elements onto the end of array
func ArrayPush(s *[]string, args ...string) {

	*s = append(*s, args...)
}
