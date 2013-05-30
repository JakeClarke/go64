package main

import (
	"testing"
)

func TestCPU(t *testing.T) {
	cpu := NewCPU()
	startingPC := cpu.PC
	cpu.Tick()
	if cpu.PC != startingPC+4 {
		t.Errorf("PC did not increment as expected. PC = %v, wanted = %v.", cpu.PC, startingPC+4)
	}
}
