package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/codecrafters-io/bittorrent-starter-go/cmd/mybittorrent/decode"
	"github.com/codecrafters-io/bittorrent-starter-go/cmd/mybittorrent/encode"
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
		{"d2:443:bar5:helloi52ee", map[string]interface{}{"44": "bar", "hello": 52}},
		{
			"d8:announce40:http://torrent.example.com:6969/announce4:infod6:lengthi12345e4:name8:file.txtee",
			map[string]interface{}{
				"announce": "http://torrent.example.com:6969/announce",
				"info": map[string]interface{}{
					"length": 12345,
					"name":   "file.txt",
				},
			},
		},
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

func TestPrintTrackerURL(t *testing.T) {
	bencodedString := "d8:announce40:http://torrent.example.com:6969/announce4:infod6:lengthi12345e4:name8:file.txtee"

	// Redirect standard output to a buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function that prints the tracker URL
	decoded, _, err := decode.Debencode(bencodedString)
	if err != nil {
		t.Fatal(err)
	}

	decodedMap, ok := decoded.(map[string]interface{})
	if !ok {
		t.Fatal("Decoded value is not a map")
	}

	announce, exists := decodedMap["announce"]
	if !exists {
		t.Fatal("Property 'announce' does not exist")
	}

	fmt.Println("Tracker URL:", announce)

	info, ok := decodedMap["info"].(map[string]interface{})
	if !ok {
		fmt.Println("Property 'info' is not a map")
		return
	}

	length, exists := info["length"]
	if !exists {
		fmt.Println("Property 'length' does not exist")
		return
	}

	fmt.Println("Length:", length)

	// Stop redirecting standard output
	w.Close()
	os.Stdout = oldStdout

	// Read the buffer's content
	var buf bytes.Buffer
	buf.ReadFrom(r)

	// Check if the output is as expected
	expectedOutput := "Tracker URL: http://torrent.example.com:6969/announce\nLength: 12345\n"
	if buf.String() != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, buf.String())
	}

}

func TestDecodeEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test 1",
			input:    "d8:announce40:http://torrent.example.com:6969/announce4:infod6:lengthi12345e4:name8:file.txtee",
			expected: "d8:announce40:http://torrent.example.com:6969/announce4:infod6:lengthi12345e4:name8:file.txtee",
		},
		{
			name:     "Test 2",
			input:    "llee",
			expected: "le",
		},
		{
			name:     "Test 3",
			input:    "d2:443:bar5:helloi52ee",
			expected: "d2:443:bar5:helloi52ee",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace DecodeEncode with the actual function you want to test

			decodeOutput, _, _ := decode.Debencode(tt.input)
			output, _ := encode.Bencode(decodeOutput)

			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("got %v, want %v", output, tt.expected)
			}
		})
	}
}
