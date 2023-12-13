package utils

import "testing"

func TestHandelWord(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
		err  bool
	}{
		{
			name: "Test with space",
			arg:  " test ",
			want: "test",
			err:  false,
		},
		{
			name: "Test with words",
			arg:  " test test",
			want: "",
			err:  true,
		},
	}
	for _, test := range tests {
		got, err := HandelWord(test.arg)

		if ok := err != nil; ok != test.err {
			t.Errorf("Test: %s | want %v got %v", test.name, test.err, ok)
		}
		if got != test.want {
			t.Errorf("Test: %s | want %s got %s", test.name, test.want, got)
		}
	}
}
