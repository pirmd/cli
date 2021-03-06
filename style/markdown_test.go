package style

import (
	"github.com/pirmd/verify"
	"testing"
)

func TestMkdTextEscaping(t *testing.T) {
	st := NewMarkdown()

	testCases := []struct {
		in  string
		out string
	}{
		{"*Test* is _fun_!", "\\*Test\\* is \\_fun\\_\\!"},
		{"**Test** is [really] _fun_!", "\\*\\*Test\\*\\* is \\[really\\] \\_fun\\_\\!"},
		{"Test a backslash \\", "Test a backslash \\\\"},
	}

	for _, tc := range testCases {
		got := st.Escape(tc.in)
		verify.EqualString(t, got, tc.out, "escape markdown: '%s'", tc.in)

		gotgot := st.Escape(tc.in)
		verify.EqualString(t, gotgot, tc.out, "escape markdown (second pass): '%s'", tc.in)
	}
}
