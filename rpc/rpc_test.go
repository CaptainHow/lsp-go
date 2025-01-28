package rpc_test

import (
	"golsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecoding(t *testing.T) {
	expectedLength := 15
	expectedMethod := "hi"
	method, length, err := rpc.DecodeMessage([]byte("Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"))
	if err != nil {
		t.Fatal(err)
	}
	if expectedLength != length {
		t.Fatalf("Expected: %d, Actual: %d", expectedLength, length)
	}
	if expectedMethod != method {
		t.Fatalf("Expected: %s, Actual: %s", expectedMethod, method)
	}
}
