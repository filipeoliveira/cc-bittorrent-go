package main

import (
	"reflect"
	"testing"

	"github.com/codecrafters-io/bittorrent-starter-go/cmd/mybittorrent/decode"
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
		{"d3:foo3:bar5:helloi52ee", map[string]interface{}{"foo": "bar", "hello": 52}},
		{"de", map[string]interface{}{}},
		{"di44e3:bar5:helloi52ee", map[string]interface{}{"44": "bar", "hello": 52}},
	}

	for _, tc := range testCases {
		result, _, err := decode.Debencode(tc.str)
		if err != nil {
			t.Fatalf("decode(%q) returned error: %s", tc.str, err)
		}
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("decode(%q) = %v, want %v", tc.str, result, tc.expected)
		}
	}
}
