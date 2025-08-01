package main

import (
	"testing"

	"go.bug.st/serial"
)

func TestGetConf(t *testing.T) {
	sc, radio := getConf()
	if sc.dev != "/dev/ttyUSB0" {
		t.Errorf("Expected /dev/ttyUSB0, got %s", sc.dev)
	}
	if sc.baudRate != 38400 {
		t.Errorf("Expected 38400, got %d", sc.baudRate)
	}
	if sc.parity != serial.NoParity {
		t.Errorf("Expected none, got %d", sc.parity)
	}
	if sc.dataBits != 8 {
		t.Errorf("Expected 8, got %d", sc.dataBits)
	}
	if sc.stopBits != 1 {
		t.Errorf("Expected 1, got %d", sc.stopBits)
	}
	if radio != "FT450D" {
		t.Errorf("Expected FT450D, got %s", radio)
	}
}
