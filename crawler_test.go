package main

import "testing"

func TestSanitizeFilename(t *testing.T) {
	cases := map[string]string{
		"My File":         "My_File",
		"A/B/C":           "A-B-C",
		"CleanName":       "CleanName",
		"Name with space": "Name_with_space",
	}
	for input, expected := range cases {
		got := sanitizeFilename(input)
		if got != expected {
			t.Errorf("sanitizeFilename(%q) = %q; want %q", input, got, expected)
		}
	}
}

func TestLoadURLs_Error(t *testing.T) {
	_, err := loadURLs("nonexistent.json")
	if err == nil {
		t.Error("Expected error when reading nonexistent file")
	}
}
