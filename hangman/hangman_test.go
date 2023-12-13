package hangman

import "testing"

func TestTryLetter(t *testing.T) {
	hangman := New("test")

	tests := []struct {
		name         string
		arg          string
		wantF, wantU int
		err          bool
	}{
		{
			name:  "Test good arg",
			arg:   "t",
			wantF: 3,
			wantU: 2,
			err:   false,
		},
		{
			name:  "Test wrong arg",
			arg:   "z",
			wantF: 2,
			wantU: 3,
			err:   false,
		},
		{
			name:  "Test bad arg",
			arg:   "to",
			wantF: 2,
			wantU: 2,
			err:   true,
		},
	}
	for _, test := range tests {
		hangman.letterFound = []string{"s", "e"}
		hangman.lettersUsed = []string{"a", "u"}
		err := hangman.TryLetter(test.arg)

		if ok := err != nil; ok != test.err {
			t.Errorf("Test: %s | err want %v got %v", test.name, test.err, ok)
		}
		got := len(hangman.letterFound)
		if got != test.wantF {
			t.Errorf("Test: %s | letterFound want %v got %v", test.name, test.wantF, got)
		}
		got = len(hangman.lettersUsed)
		if got != test.wantU {
			t.Errorf("Test: %s | lettersUsed want %v got %v", test.name, test.wantU, got)
		}
	}

}
func TestTryNewt(t *testing.T) {

	tests := []struct {
		arg  string
		want string
	}{
		{
			arg:  "test",
			want: "test",
		},
	}
	for _, test := range tests {
		got := New(test.arg).word

		if got != test.want {
			t.Errorf(" letterFound want %s got %s", test.want, got)
		}

	}

}
