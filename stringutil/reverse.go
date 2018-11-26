package stringutil

//...

/* Reverse a string .....
//
// - Hallo wird zu ...
// äöüöäö */
func Reverse(t string) string {
	r := []rune(t)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

