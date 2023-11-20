package support

import (
	
	"testing"
)



func TestGetDefaultEventValueIfNull(t *testing.T) {
    tests := []struct {
        input string
        expected string
    }{
        {"World", "World"},
    }

    for _, test := range tests {
        actual := getDefaultEventValueIfNull(test.input)
        if actual != test.expected {
            t.Errorf("getDefaultEventValueIfNull(%s) = %s; expected %s", test.input, actual, test.expected)
        }
    }
}


func TestGetWhitespaceStr(t *testing.T) {
    tests := []struct {
        size     int
        expected string
    }{
        {0, ""},
        {1, " "},
        {5, "     "},
        {10, "          "},
    }

    for _, test := range tests {
        actual := getWhitespaceStr(test.size)
        if actual != test.expected {
            t.Errorf("getWhitespaceStr(%d) = %s; expected %s", test.size, actual, test.expected)
        }
    }
}
