package str

import "regexp"

var repRe = regexp.MustCompile(`\\u002F`)

func DelLongSlash(in string) string {

	return repRe.ReplaceAllString(in, "/")

}
