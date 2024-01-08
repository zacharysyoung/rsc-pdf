package pdf

import "testing"

func TestLastLine(t *testing.T) {
	for _, tc := range []struct {
		bufs, s string
		want    int
	}{
		{"\nFoo\n", "Foo", 1},
		{"Foo\n", "Foo", -1}, // not preceeded by newline
		{"\nFoo", "Foo", -1}, // not proceeded by newline

		{"\r\nFoo\n", "Foo", 2},
		{"\r\nFoo bar\nFoo", "Foo", -1}, // Foo in "Foo bar" not the complete line
		{"\r\nFoo bar\nFoo", "Foo bar", 2},
		{"\r\nFoo\nFoo\n", "Foo", 6},
	} {
		got := findLastLine([]byte(tc.bufs), tc.s)
		if got != tc.want {
			t.Errorf("findLastLine(%q, %q)=%d != %d", tc.bufs, tc.s, got, tc.want)
		}
	}
}
