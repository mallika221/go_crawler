package link_parser

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "only relative",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/blog/intro">Intro</a>
	</body>
</html>`,
			expected: []string{"https://blog.boot.dev/blog/intro"},
		},
		{
			name:     "no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a>No link here</a>
	</body>
</html>`,
			expected: []string{},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %d - %s: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %d - %s FAIL:\nExpected: %#v\nGot:      %#v", i, tc.name, tc.expected, actual)
			} else {
				t.Logf("Test %d - %s PASS", i, tc.name) // Optional: logs successful test
			}
		})
	}

}
