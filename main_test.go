package main

import (
  "testing"
)

func TestSerializeRESP(t *testing.T) {
  response := serializeRESP("PONG")
  expected := "$4\r\nPONG\r\n"
  if response != expected {
    t.Errorf("Expected %s, got %s", expected, response)
  }
}
