package main

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	testCases := []struct {
		str      string
		expected interface{}
	}{
		{"i-52e", -52},
		{"i35e", 35},
		{"4:spam", "spam"},
		{"lli35eee", []interface{}{[]interface{}{35}}},
		{"lli35eei44ee", []interface{}{[]interface{}{35}, 44}},
		{"llleee", []interface{}{[]interface{}{[]interface{}{}}}},
		{"le", []interface{}{}},
		{"l6:bananai335ee", []interface{}{"banana", 335}},
	}

	for _, tc := range testCases {
		result, _, err := decode(tc.str)
		if err != nil {
			t.Fatalf("decode(%q) returned error: %s", tc.str, err)
		}
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("decode(%q) = %v, want %v", tc.str, result, tc.expected)
		}
	}
}
