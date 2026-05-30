package main

import (
	"strings"
	"testing"
)

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputHTML     string
		expected      string
		errorContains string
	}{
		{
			name:      "extract h1",
			inputHTML: "<h1> heading 1 </h1>",
			expected:  " heading 1 ",
		},
		{
			name:      "extract h1 when h2 present",
			inputHTML: "<h1> heading 1 </h1> <h2> heading 2 </h2>",
			expected:  " heading 1 ",
		},
		{
			name:      "extract h2 when no h1",
			inputHTML: "<h2> heading 2 </h2>",
			expected:  " heading 2 ",
		},
		{
			name:          "extract no heading",
			inputHTML:     "<p> heading 1 </p>",
			expected:      "",
			errorContains: "no headings found",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getHeadingFromHTML(tc.inputHTML)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected HTML: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputHTML     string
		expected      string
		errorContains string
	}{
		{
			name:      "extract p",
			inputHTML: "<p> 1 </p>",
			expected:  " 1 ",
		},
		{
			name:      "extract first p ",
			inputHTML: "<p> 1 </p> <p> 2 </p>",
			expected:  " 1 ",
		},
		{
			name:      "extract main p",
			inputHTML: "<main> <p> main p </p> </main>",
			expected:  " main p ",
		},
		{
			name:          "fall back to regular p",
			inputHTML:     "<main> </main> <p> not main p </p> ",
			expected:      " not main p ",
			errorContains: "",
		},
		{
			name:          "extract no p",
			inputHTML:     "<h1> heading 1 </h1>",
			expected:      "",
			errorContains: "no paragraphs found",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getFirstParagraphFromHTML(tc.inputHTML)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected HTML: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
