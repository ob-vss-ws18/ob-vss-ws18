package moviestore

import "testing"

func TestMovie(t *testing.T) {
	cases := []struct {
		in   Movie
		want string
	}{
		{Movie{"Am Limit", FSK0, 12}, "  12. Am Limit (Ab 0)"},
		{Movie{"Texas Chainsaw Massacre", FSK18, 8}, "   8. Texas Chainsaw Massacre (Ab 18)"},
		{Movie{"Inglourious Basterds", FSK16, 13}, "  13. Inglourious Basterds (Ab 16)"},
	}
	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("%q.String() == %q, want %q", c.in, got, c.want)
		}
	}
}
