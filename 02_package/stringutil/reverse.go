package stringutil

//Reverse a string, accesses unexported function within package
func Reverse(s string) string {
	return reverseTwo(s)
}
