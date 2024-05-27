package words_test

import (
	"github.com/eduardoraider/go-unit-test/words"
	"testing"
)

func TestCountWords(t *testing.T) {
	testcases := []struct {
		Text   string
		Filter string
		Expect int
	}{
		{Text: "Hello World", Expect: 2},
		{Text: "Technology is changing everything.", Filter: "is", Expect: 1},
		{Text: "Hello  Go  Developer ", Expect: 3},
		{Text: "", Expect: 0},
		{Text: "Technology is changing everything.", Filter: "here", Expect: 0},
	}

	for _, tc := range testcases {
		count := words.CountWords(tc.Text, tc.Filter)
		if count != tc.Expect {
			t.Fatalf("expected: %d, got: %d", tc.Expect, count)
		}
	}
}
