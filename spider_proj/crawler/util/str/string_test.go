package str

import "testing"

var cases = []struct {
	i   string
	out string
}{{
	i:   "http:\u002F\u002Fzhenai.com\u002Fu",
	out: "http://zhenai.com/u"}}

func TestDelLongSlash(t *testing.T) {
	for _, c := range cases {
		if c.out != DelLongSlash(c.i) {
			t.Errorf("want:%s, given: %s ", c.out, DelLongSlash(c.i))
		}
	}
}
