package stringreverse

import "testing"

func TestReverse(t *testing.T) {
	cases := &[]struct {
		name, in, want string
	}{
		{"normal", "Hello, world", "dlrow ,olleH"},
		{"normal and chinese", "Hello, 世界", "界世 ,olleH"},
		{"numbers", "1234567890", "0987654321"},
		{"symbols", "~!@#$%^&*()_+", "+_)(*&^%$#@!~"},
		{"empty", "", ""},
	}
	for _, c := range *cases {
		t.Run(c.name, func(t *testing.T) {
			got := Reverse(c.in)
			if got != c.want {
				t.Fatalf("[Case %v] Reverse(%q) == %q, want %q", c.name, c.in, got, c.want)
			}
		})
	}
}
